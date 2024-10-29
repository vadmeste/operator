//go:build !dev

// Copyright 2024 - MinIO, Inc. All rights reserved.

package main

import (
	"time"

	"github.com/miniohq/license/go/license"
)

func verifyLicenseLoop() string {
	for {
		_, err := license.LookupAndVerify(func() (string, error) {
			return license.Lookup()
		}, nil)
		if err != nil {
			panic("unable to validate the license, terminating process..")
		}

		time.Sleep(5 * time.Minute)
	}
}
