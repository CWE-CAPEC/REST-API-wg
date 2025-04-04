package repo

import "embed"

//go:embed W C V *.json
var Repo embed.FS
