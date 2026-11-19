// Copyright (C) 2019 Nicola Murino
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package vfs

import (
	"testing"

	"github.com/sftpgo/sdk"

	"github.com/drakkan/sftpgo/v2/internal/kms"
)

// The memory provider stores and returns users through GetACopy. A copy that
// drops the session token leaves temporary (STS) credentials unusable: S3
// rejects the access key with InvalidAccessKeyId.
func TestS3SessionTokenSurvivesCopy(t *testing.T) {
	fs := Filesystem{
		Provider: sdk.S3FilesystemProvider,
		S3Config: S3FsConfig{
			BaseS3FsConfig: sdk.BaseS3FsConfig{
				Bucket:       "bucket",
				Region:       "ca-central-1",
				AccessKey:    "ASIAEXAMPLE",
				SessionToken: "session-token",
			},
			AccessSecret: kms.NewPlainSecret("secret"),
		},
	}
	fsCopy := fs.GetACopy()
	if fsCopy.S3Config.SessionToken != fs.S3Config.SessionToken {
		t.Fatalf("session token not copied: got %q, want %q", fsCopy.S3Config.SessionToken, fs.S3Config.SessionToken)
	}
	if fsCopy.S3Config.AccessKey != fs.S3Config.AccessKey {
		t.Fatalf("access key not copied: got %q, want %q", fsCopy.S3Config.AccessKey, fs.S3Config.AccessKey)
	}
}
