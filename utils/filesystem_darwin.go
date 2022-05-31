// Copyright 2022 Netapp Inc. All Rights Reserved.

// NOTE: This file should only contain functions for handling the filesystem for Darwin flavor

package utils

import (
	"errors"

	"golang.org/x/net/context"

	. "github.com/netapp/trident/logger"
)

func GetFilesystemStats(ctx context.Context, _ string) (int64, int64, int64, int64, int64, int64, error) {
	Logc(ctx).Debug(">>>> ofilesystem_darwin.GetFilesystemStats")
	defer Logc(ctx).Debug("<<<< ofilesystem_darwin.GetFilesystemStats")
	return 0, 0, 0, 0, 0, 0, errors.New("GetFilesystemStats is not supported for darwin")
}

func getFilesystemSize(ctx context.Context, _ string) (int64, error) {
	Logc(ctx).Debug(">>>> ofilesystem_darwin.getFilesystemSize")
	defer Logc(ctx).Debug("<<<< ofilesystem_darwin.getFilesystemSize")
	return 0, errors.New("getFilesystemSize is not supported for darwin")
}