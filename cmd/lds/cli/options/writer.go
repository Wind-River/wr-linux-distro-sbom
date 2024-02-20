/*
 * Copyright (C) 2023 Wind River Systems, Inc.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */

package options

import (
	"fmt"
	"io"
	"os"

	"github.com/anchore/syft/syft/sbom"
)

func MakeSBOMWriterForFormat(format sbom.Format, path string) (sbom.Writer, error) {
	fileOut, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("unable to create report file: %w", err)
	}
	writer := &sbomStreamWriter{
		format: format,
		out:    fileOut,
	}

	return writer, nil
}

type sbomStreamWriter struct {
	format sbom.Format
	out    io.Writer
}

func (w *sbomStreamWriter) Write(s sbom.SBOM) error {
	defer w.Close()
	return w.format.Encode(w.out, s)
}

func (w *sbomStreamWriter) Close() error {
	if closer, ok := w.out.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}
