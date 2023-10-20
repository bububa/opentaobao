package scbp

import (
	"sync"
)

// AlibabaScbpAdCustomerFindCustomerInfoResultDto 结构体
type AlibabaScbpAdCustomerFindCustomerInfoResultDto struct {
	// 返回实体
	Result *TopCustomerDto `json:"result,omitempty" xml:"result,omitempty"`
}

var poolAlibabaScbpAdCustomerFindCustomerInfoResultDto = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdCustomerFindCustomerInfoResultDto)
	},
}

// GetAlibabaScbpAdCustomerFindCustomerInfoResultDto() 从对象池中获取AlibabaScbpAdCustomerFindCustomerInfoResultDto
func GetAlibabaScbpAdCustomerFindCustomerInfoResultDto() *AlibabaScbpAdCustomerFindCustomerInfoResultDto {
	return poolAlibabaScbpAdCustomerFindCustomerInfoResultDto.Get().(*AlibabaScbpAdCustomerFindCustomerInfoResultDto)
}

// ReleaseAlibabaScbpAdCustomerFindCustomerInfoResultDto 释放AlibabaScbpAdCustomerFindCustomerInfoResultDto
func ReleaseAlibabaScbpAdCustomerFindCustomerInfoResultDto(v *AlibabaScbpAdCustomerFindCustomerInfoResultDto) {
	v.Result = nil
	poolAlibabaScbpAdCustomerFindCustomerInfoResultDto.Put(v)
}
