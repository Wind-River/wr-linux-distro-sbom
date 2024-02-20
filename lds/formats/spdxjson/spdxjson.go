/*
 * Copyright (C) 2023 Wind River Systems, Inc.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */

package spdxjson

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/anchore/syft/syft/formats/common/spdxhelpers"
	"github.com/anchore/syft/syft/sbom"
	"github.com/spdx/tools-golang/spdx"
)

func Format() sbom.Format {
	return sbom.NewFormat(
		"2.3",
		encoder,
		nil,
		nil,
		"spdx-json",
	)
}

func encoder(output io.Writer, s sbom.SBOM) error {
	doc := spdxhelpers.ToFormatModel(s)

	doc.CreationInfo.Creators[0] = spdx.Creator{
		Creator:     "Wind River, Inc",
		CreatorType: "Organization",
	}
	doc.DocumentComment = "DISTRO: " + s.Artifacts.LinuxDistribution.ID + "-" + s.Artifacts.LinuxDistribution.VersionID
	if s.Artifacts.LinuxDistribution.ID == "rhel" {
		doc.DocumentComment = "DISTRO: " + "red hat" + "-" + s.Artifacts.LinuxDistribution.VersionID
	}

	//The dependent go modules produce no CVE. Remove the dependency packages.
	for i := len(doc.Packages) - 1; i >= 0; i-- {
		if strings.Contains(doc.Packages[i].PackageSourceInfo, "go module") {
			if strings.Contains(doc.Packages[i].PackageVersion, "(devel)") {
				continue
			}

			mod_name := strings.Split(doc.Packages[i].PackageName, "/")
			source_path := strings.Fields(doc.Packages[i].PackageSourceInfo)
			if strings.Contains(source_path[len(source_path)-1], mod_name[len(mod_name)-1]) {
				continue
			}
			doc.Packages = append(doc.Packages[:i], doc.Packages[i+1:]...)
		}
	}
	return encodeJSON(output, doc)
}

func encodeJSON(output io.Writer, doc interface{}) error {
	enc := json.NewEncoder(output)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", " ")

	return enc.Encode(doc)
}
