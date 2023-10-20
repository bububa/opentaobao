package tmallgeniescp

import (
	"sync"
)

// IbpMaterielDto 结构体
type IbpMaterielDto struct {
	// 层级的code
	Level6 string `json:"level6,omitempty" xml:"level6,omitempty"`
	// 层级的code
	Level5 string `json:"level5,omitempty" xml:"level5,omitempty"`
	// 层级的code
	Level4 string `json:"level4,omitempty" xml:"level4,omitempty"`
	// 层级的code
	Level3 string `json:"level3,omitempty" xml:"level3,omitempty"`
	// 层级的code
	Level2 string `json:"level2,omitempty" xml:"level2,omitempty"`
	// 层级的code
	Level1 string `json:"level1,omitempty" xml:"level1,omitempty"`
	// PLM状态：1. EOP（停止生产） 2. NPI（新品） 3. RUN（正在使用产品）4. EOL（退市品）
	PlmStatus string `json:"plm_status,omitempty" xml:"plm_status,omitempty"`
	// 单位描述：写死个
	UnitDesc string `json:"unit_desc,omitempty" xml:"unit_desc,omitempty"`
	// 单位：写死PC
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 1. FIN（成品） 2.RAW（原料）
	MaterielType string `json:"materiel_type,omitempty" xml:"materiel_type,omitempty"`
	// 物料名称
	MaterielName string `json:"materiel_name,omitempty" xml:"materiel_name,omitempty"`
	// 物料编码
	MaterielCode string `json:"materiel_code,omitempty" xml:"materiel_code,omitempty"`
	// 扩展字段
	ExtendJson string `json:"extend_json,omitempty" xml:"extend_json,omitempty"`
	// 租户
	Tenant string `json:"tenant,omitempty" xml:"tenant,omitempty"`
}

var poolIbpMaterielDto = sync.Pool{
	New: func() any {
		return new(IbpMaterielDto)
	},
}

// GetIbpMaterielDto() 从对象池中获取IbpMaterielDto
func GetIbpMaterielDto() *IbpMaterielDto {
	return poolIbpMaterielDto.Get().(*IbpMaterielDto)
}

// ReleaseIbpMaterielDto 释放IbpMaterielDto
func ReleaseIbpMaterielDto(v *IbpMaterielDto) {
	v.Level6 = ""
	v.Level5 = ""
	v.Level4 = ""
	v.Level3 = ""
	v.Level2 = ""
	v.Level1 = ""
	v.PlmStatus = ""
	v.UnitDesc = ""
	v.Unit = ""
	v.MaterielType = ""
	v.MaterielName = ""
	v.MaterielCode = ""
	v.ExtendJson = ""
	v.Tenant = ""
	poolIbpMaterielDto.Put(v)
}
