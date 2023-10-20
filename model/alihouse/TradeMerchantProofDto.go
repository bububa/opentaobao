package alihouse

import (
	"sync"
)

// TradeMerchantProofDto 结构体
type TradeMerchantProofDto struct {
	// 文件URL
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 文件类型
	FileType int64 `json:"file_type,omitempty" xml:"file_type,omitempty"`
}

var poolTradeMerchantProofDto = sync.Pool{
	New: func() any {
		return new(TradeMerchantProofDto)
	},
}

// GetTradeMerchantProofDto() 从对象池中获取TradeMerchantProofDto
func GetTradeMerchantProofDto() *TradeMerchantProofDto {
	return poolTradeMerchantProofDto.Get().(*TradeMerchantProofDto)
}

// ReleaseTradeMerchantProofDto 释放TradeMerchantProofDto
func ReleaseTradeMerchantProofDto(v *TradeMerchantProofDto) {
	v.Url = ""
	v.Type = 0
	v.FileType = 0
	poolTradeMerchantProofDto.Put(v)
}
