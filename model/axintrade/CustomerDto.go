package axintrade

import (
	"sync"
)

// CustomerDto 结构体
type CustomerDto struct {
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 英文名
	FirstName string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// 英文姓
	LastName string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	// 入住人类型1. 成人 2. 儿童
	CustomerType int64 `json:"customer_type,omitempty" xml:"customer_type,omitempty"`
}

var poolCustomerDto = sync.Pool{
	New: func() any {
		return new(CustomerDto)
	},
}

// GetCustomerDto() 从对象池中获取CustomerDto
func GetCustomerDto() *CustomerDto {
	return poolCustomerDto.Get().(*CustomerDto)
}

// ReleaseCustomerDto 释放CustomerDto
func ReleaseCustomerDto(v *CustomerDto) {
	v.Name = ""
	v.FirstName = ""
	v.LastName = ""
	v.CustomerType = 0
	poolCustomerDto.Put(v)
}
