/*
 * Copyright (C) 2023 Wind River Systems, Inc.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */

package packages

import (
	"crypto"
	"fmt"
	"log"

	"github.com/Wind-River/lds/cmd/lds/cli/options"
	"github.com/Wind-River/lds/internal/config"
	"github.com/Wind-River/lds/lds/formats/spdxjson"
	"github.com/anchore/syft/syft"
	"github.com/anchore/syft/syft/pkg/cataloger"
	"github.com/anchore/syft/syft/sbom"
	"github.com/anchore/syft/syft/source"
)

const version = "1.0"

func Run(app *config.Application, args []string) error {
	userInput := args[0]

	output_format, output_file := app.GetOutputFormat(), app.GetOutputFile()

	writer, err := makeSBOMWriter(output_format, output_file)
	if err != nil {
		return err
	}

	src, err := makeNewSource(userInput)
	if err != nil {
		return err
	}

	defer func() {
		if src != nil {
			if err := src.Close(); err != nil {
				log.Fatal("unable to close source: %v", err)
			}
		}
	}()

	s, err := generateSBOM(src)
	if err != nil {
		return err
	}

	if s == nil {
		return fmt.Errorf("no SBOM produced for %q", userInput)
	}

	if err := writer.Write(*s); err != nil {
		return fmt.Errorf("failed to write SBOM: %w", err)
	}

	return nil
}

func makeSBOMWriter(format, file string) (sbom.Writer, error) {
	return options.MakeSBOMWriterForFormat(spdxjson.Format(), file)
}

func makeNewSource(userInput string) (source.Source, error) {
	detection, err := source.Detect(
		userInput, source.DetectConfig{
			DefaultImageSource: "",
		})
	if err != nil {
		return nil, fmt.Errorf("could not determinate source: %w", err)
	}

	hashers := []crypto.Hash{crypto.SHA256}

	src, err := detection.NewSource(
		source.DetectionSourceConfig{
			Alias: source.Alias{
				Name:    "",
				Version: "",
			},
			DigestAlgorithms: hashers,
			BasePath:         "",
		})
	if err != nil {
		return nil, fmt.Errorf("failed to construct source from user input %q: %w", userInput, err)
	}

	return src, nil
}

func generateSBOM(src source.Source) (*sbom.SBOM, error) {
	s := &sbom.SBOM{
		Source: src.Describe(),
		Descriptor: sbom.Descriptor{
			Name:    "LDS",
			Version: version,
		},
	}

	buildRelationships(s, src)

	return s, nil
}

func buildRelationships(s *sbom.SBOM, src source.Source) {
	cataloger_config := cataloger.DefaultConfig()
	packageCatalog, relationships, determinated_distro, err := syft.CatalogPackages(src, cataloger_config)

	// FIXME: handle err properly
	if err != nil {
		panic(err)
	}

	s.Artifacts.Packages = packageCatalog
	s.Artifacts.LinuxDistribution = determinated_distro
	s.Relationships = relationships
}
