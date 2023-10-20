package drugtrace

import (
	"sync"
)

// PiecesProduceInfoDto 结构体
type PiecesProduceInfoDto struct {
	// 生产管理图片（上传图片）图片建议尺寸：height: 310px;width: 670px;
	ProductionPictures []string `json:"production_pictures,omitempty" xml:"production_pictures>string,omitempty"`
	// 生产开始日期yyyy-MM-dd
	ProductionStartDate string `json:"production_start_date,omitempty" xml:"production_start_date,omitempty"`
	// 包装规格
	PackageSpec string `json:"package_spec,omitempty" xml:"package_spec,omitempty"`
	// 生产工艺
	ProductionProcess string `json:"production_process,omitempty" xml:"production_process,omitempty"`
	// 生产数量
	ProductionSum string `json:"production_sum,omitempty" xml:"production_sum,omitempty"`
	// 生产结束日期yyyy-MM-dd
	ProductionEndDate string `json:"production_end_date,omitempty" xml:"production_end_date,omitempty"`
	// 包装材质
	PackingMaterial string `json:"packing_material,omitempty" xml:"packing_material,omitempty"`
}

var poolPiecesProduceInfoDto = sync.Pool{
	New: func() any {
		return new(PiecesProduceInfoDto)
	},
}

// GetPiecesProduceInfoDto() 从对象池中获取PiecesProduceInfoDto
func GetPiecesProduceInfoDto() *PiecesProduceInfoDto {
	return poolPiecesProduceInfoDto.Get().(*PiecesProduceInfoDto)
}

// ReleasePiecesProduceInfoDto 释放PiecesProduceInfoDto
func ReleasePiecesProduceInfoDto(v *PiecesProduceInfoDto) {
	v.ProductionPictures = v.ProductionPictures[:0]
	v.ProductionStartDate = ""
	v.PackageSpec = ""
	v.ProductionProcess = ""
	v.ProductionSum = ""
	v.ProductionEndDate = ""
	v.PackingMaterial = ""
	poolPiecesProduceInfoDto.Put(v)
}
