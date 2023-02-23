package metadata

import (
	"encoding/json"

	// 使用go 1.16以上版本的embed功能
	_ "embed"

	"github.com/bububa/opentaobao/metadata/util"
)

// PkgConfig SDK分包配置结构体
type PkgConfig struct {
	Name string `json:"name,omitempty"` // api catelog name
	Pkg  string `json:"pkg,omitempty"`  // go package name
	Id   int64  `json:"id,omitempty"`   // api catelog id
}

//go:embed assets/package.json
var pkgBytes []byte

//go:embed assets/conflict_models.json
var conflictModelsBytes []byte

// PkgsConfig SDK各分包配置
var PkgsConfig []PkgConfig

// PkgConfigSlice SDK分包配置按Id排序
type PkgConfigSlice []PkgConfig

func (p PkgConfigSlice) Len() int           { return len(p) }
func (p PkgConfigSlice) Less(i, j int) bool { return p[i].Id < p[j].Id }
func (p PkgConfigSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// ConflictModels conflict_models 命名冲突struct
var ConflictModels []string

func init() {
	buf := util.GetBufferPool()
	defer util.PutBufferPool(buf)
	if err := json.Compact(buf, pkgBytes); err != nil {
		panic(err)
	}
	if err := json.NewDecoder(buf).Decode(&PkgsConfig); err != nil {
		panic(err)
	}

	buf.Reset()
	if err := json.Compact(buf, conflictModelsBytes); err != nil {
		panic(err)
	}
	if err := json.NewDecoder(buf).Decode(&ConflictModels); err != nil {
		panic(err)
	}
}
