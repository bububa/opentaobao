package wdk

import (
	"sync"
)

// AxStoreQueryResponse 结构体
type AxStoreQueryResponse struct {
	// 负责人姓名
	PrincipalName string `json:"principal_name,omitempty" xml:"principal_name,omitempty"`
	// 负责人
	Principal string `json:"principal,omitempty" xml:"principal,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 区编码
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 市编码
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 省编码
	Prov string `json:"prov,omitempty" xml:"prov,omitempty"`
	// 门店名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 门店编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 商家code
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 门店经营状态 1营业 0闭店
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolAxStoreQueryResponse = sync.Pool{
	New: func() any {
		return new(AxStoreQueryResponse)
	},
}

// GetAxStoreQueryResponse() 从对象池中获取AxStoreQueryResponse
func GetAxStoreQueryResponse() *AxStoreQueryResponse {
	return poolAxStoreQueryResponse.Get().(*AxStoreQueryResponse)
}

// ReleaseAxStoreQueryResponse 释放AxStoreQueryResponse
func ReleaseAxStoreQueryResponse(v *AxStoreQueryResponse) {
	v.PrincipalName = ""
	v.Principal = ""
	v.Latitude = ""
	v.Longitude = ""
	v.Address = ""
	v.Area = ""
	v.City = ""
	v.Prov = ""
	v.Name = ""
	v.Code = ""
	v.MerchantCode = ""
	v.Status = 0
	poolAxStoreQueryResponse.Put(v)
}
