package eleenterprisecartnew

import (
	"sync"
)

// CartExtraDto 结构体
type CartExtraDto struct {
	// 费用
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 费用名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 费用id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 订单项目分类
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
}

var poolCartExtraDto = sync.Pool{
	New: func() any {
		return new(CartExtraDto)
	},
}

// GetCartExtraDto() 从对象池中获取CartExtraDto
func GetCartExtraDto() *CartExtraDto {
	return poolCartExtraDto.Get().(*CartExtraDto)
}

// ReleaseCartExtraDto 释放CartExtraDto
func ReleaseCartExtraDto(v *CartExtraDto) {
	v.Price = ""
	v.Name = ""
	v.Description = ""
	v.Quantity = 0
	v.Id = 0
	v.CategoryId = 0
	poolCartExtraDto.Put(v)
}
