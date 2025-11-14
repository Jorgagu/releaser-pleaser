package updater

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCargoUpdater_Files(t *testing.T) {
	assert.Equal(t, []string{"Cargo.toml"}, Cargo().Files())
}

func TestCargoUpdater_CreateNewFiles(t *testing.T) {
	assert.False(t, Cargo().CreateNewFiles())
}

func TestCargoUpdater_Update(t *testing.T) {
	tests := []updaterTestCase{
		{
			name: "simple Cargo.toml",
			content: `[package]
name = "test"
version = "1.0.0"`,
			info: ReleaseInfo{
				Version: "v2.0.5",
			},
			want: `[package]
name = "test"
version = "2.0.5"`,
			wantErr: assert.NoError,
		},
		{
			name: "complex Cargo.toml",
			content: `[package]
name = "test"
version = "1.0.0"
edition = "2021"

[dependencies]
serde = "1.0"`,
			info: ReleaseInfo{
				Version: "v2.0.0",
			},
			want: `[package]
name = "test"
version = "2.0.0"
edition = "2021"

[dependencies]
serde = "1.0"`,
			wantErr: assert.NoError,
		},
		{
			name:    "different spacing",
			content: `version="1.0.0"`,
			info: ReleaseInfo{
				Version: "v3.1.4",
			},
			want:    `version="3.1.4"`,
			wantErr: assert.NoError,
		},
		{
			name:    "extra whitespace",
			content: `version   =   "1.0.0"`,
			info: ReleaseInfo{
				Version: "v1.2.3",
			},
			want:    `version   =   "1.2.3"`,
			wantErr: assert.NoError,
		},
		{
			name:    "invalid toml",
			content: `not toml`,
			info: ReleaseInfo{
				Version: "v2.0.0",
			},
			want:    `not toml`,
			wantErr: assert.NoError,
		},
		{
			name: "toml without version",
			content: `[package]
name = "test"`,
			info: ReleaseInfo{
				Version: "v2.0.0",
			},
			want: `[package]
name = "test"`,
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runUpdaterTest(t, Cargo(), tt)
		})
	}
}
