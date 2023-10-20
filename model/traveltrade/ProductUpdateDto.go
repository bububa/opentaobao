package traveltrade

import (
	"sync"
)

// ProductUpdateDto 结构体
type ProductUpdateDto struct {
	// 场次ID信息
	SessionIds []string `json:"session_ids,omitempty" xml:"session_ids>string,omitempty"`
	// 扩展参数
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
	// 系统商商品码
	ProductId string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 产品变更开始时间 yyyy-MM-dd
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 产品变更结束时间 yyyy-MM-dd
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 产品变更通知类型 1：价格，2：库存，3：价库
	NotifyType int64 `json:"notify_type,omitempty" xml:"notify_type,omitempty"`
}

var poolProductUpdateDto = sync.Pool{
	New: func() any {
		return new(ProductUpdateDto)
	},
}

// GetProductUpdateDto() 从对象池中获取ProductUpdateDto
func GetProductUpdateDto() *ProductUpdateDto {
	return poolProductUpdateDto.Get().(*ProductUpdateDto)
}

// ReleaseProductUpdateDto 释放ProductUpdateDto
func ReleaseProductUpdateDto(v *ProductUpdateDto) {
	v.SessionIds = v.SessionIds[:0]
	v.ExtendParams = ""
	v.ProductId = ""
	v.StartDate = ""
	v.EndDate = ""
	v.NotifyType = 0
	poolProductUpdateDto.Put(v)
}
