package nrpos

import (
	"sync"
)

// QueryMerchandiseDto 结构体
type QueryMerchandiseDto struct {
	// 门店号
	Storeno string `json:"storeno,omitempty" xml:"storeno,omitempty"`
	// 商品编码
	A0501 string `json:"a0501,omitempty" xml:"a0501,omitempty"`
}

var poolQueryMerchandiseDto = sync.Pool{
	New: func() any {
		return new(QueryMerchandiseDto)
	},
}

// GetQueryMerchandiseDto() 从对象池中获取QueryMerchandiseDto
func GetQueryMerchandiseDto() *QueryMerchandiseDto {
	return poolQueryMerchandiseDto.Get().(*QueryMerchandiseDto)
}

// ReleaseQueryMerchandiseDto 释放QueryMerchandiseDto
func ReleaseQueryMerchandiseDto(v *QueryMerchandiseDto) {
	v.Storeno = ""
	v.A0501 = ""
	poolQueryMerchandiseDto.Put(v)
}
