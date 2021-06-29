package metadata

import (
	"bytes"
	_ "embed"
	"encoding/json"
)

type PkgConfig struct {
	Id   int64  `json:"id,omitempty"`   // api catelog id
	Name string `json:"name,omitempty"` // api catelog name
	Pkg  string `json:"pkg,omitempty"`  // go package name
}

//go:embed assets/package.json
var pkgBytes []byte

//go:embed assets/conflict_models.json
var conflictModelsBytes []byte

var PkgsConfig []PkgConfig

type PkgConfigSlice []PkgConfig

func (p PkgConfigSlice) Len() int           { return len(p) }
func (p PkgConfigSlice) Less(i, j int) bool { return p[i].Id < p[j].Id }
func (p PkgConfigSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

var ConflictModels []string

func init() {
	buf := new(bytes.Buffer)
	if err := json.Compact(buf, pkgBytes); err != nil {
		panic(err)
	}
	if err := json.NewDecoder(buf).Decode(&PkgsConfig); err != nil {
		panic(err)
	}

	mbuf := new(bytes.Buffer)
	if err := json.Compact(mbuf, conflictModelsBytes); err != nil {
		panic(err)
	}
	if err := json.NewDecoder(mbuf).Decode(&ConflictModels); err != nil {
		panic(err)
	}
}
