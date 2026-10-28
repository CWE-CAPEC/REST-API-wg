package repo

import "embed"

// Embed the generated v1/v2 content and relationship trees so the API can serve
// them at runtime without relying on files outside the compiled binary.
// The following line is a compiler directive that tells the Go compiler to embed the specified files and directories into the binary.
// The `embed.FS` type is used to access the embedded files at runtime.

//go:embed v1 v2 *.json
var Repo embed.FS
