package tmallgeniescp

import (
	"sync"
)

// IbpPurchaseOrderDto 结构体
type IbpPurchaseOrderDto struct {
	// 周
	WeekNO string `json:"week_n_o,omitempty" xml:"week_n_o,omitempty"`
	// 物料编码
	MaterielCode string `json:"materiel_code,omitempty" xml:"materiel_code,omitempty"`
	// 地点编码
	LocationCode string `json:"location_code,omitempty" xml:"location_code,omitempty"`
	// 扩展参数
	ExtendJson string `json:"extend_json,omitempty" xml:"extend_json,omitempty"`
	// 租户
	Tenant string `json:"tenant,omitempty" xml:"tenant,omitempty"`
	// 关键日期值
	KeyFigureDate string `json:"key_figure_date,omitempty" xml:"key_figure_date,omitempty"`
	// 二级物料-分仓目的仓code
	LocationCodeTo string `json:"location_code_to,omitempty" xml:"location_code_to,omitempty"`
	// 生产实绩数量
	ProductionQuantity int64 `json:"production_quantity,omitempty" xml:"production_quantity,omitempty"`
	// 紧急的计划数量
	UrgentPlanQuantity int64 `json:"urgent_plan_quantity,omitempty" xml:"urgent_plan_quantity,omitempty"`
	// 普通的计划数量
	NormalPlanQuantity int64 `json:"normal_plan_quantity,omitempty" xml:"normal_plan_quantity,omitempty"`
	// 计划数量
	PlanQuantity int64 `json:"plan_quantity,omitempty" xml:"plan_quantity,omitempty"`
}

var poolIbpPurchaseOrderDto = sync.Pool{
	New: func() any {
		return new(IbpPurchaseOrderDto)
	},
}

// GetIbpPurchaseOrderDto() 从对象池中获取IbpPurchaseOrderDto
func GetIbpPurchaseOrderDto() *IbpPurchaseOrderDto {
	return poolIbpPurchaseOrderDto.Get().(*IbpPurchaseOrderDto)
}

// ReleaseIbpPurchaseOrderDto 释放IbpPurchaseOrderDto
func ReleaseIbpPurchaseOrderDto(v *IbpPurchaseOrderDto) {
	v.WeekNO = ""
	v.MaterielCode = ""
	v.LocationCode = ""
	v.ExtendJson = ""
	v.Tenant = ""
	v.KeyFigureDate = ""
	v.LocationCodeTo = ""
	v.ProductionQuantity = 0
	v.UrgentPlanQuantity = 0
	v.NormalPlanQuantity = 0
	v.PlanQuantity = 0
	poolIbpPurchaseOrderDto.Put(v)
}
