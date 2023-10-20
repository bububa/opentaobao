package wlbimports

import (
	"sync"
)

// TransferStoreCenterResponse 结构体
type TransferStoreCenterResponse struct {
	// 邮编
	ZipCode string `json:"zip_code,omitempty" xml:"zip_code,omitempty"`
	// 国家
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 公司
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 省份
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 联系人姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 详细地址
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	// 仓名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 邮件
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 仓code
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 国家id
	CountryId int64 `json:"country_id,omitempty" xml:"country_id,omitempty"`
	// 地区id
	AreaId int64 `json:"area_id,omitempty" xml:"area_id,omitempty"`
}

var poolTransferStoreCenterResponse = sync.Pool{
	New: func() any {
		return new(TransferStoreCenterResponse)
	},
}

// GetTransferStoreCenterResponse() 从对象池中获取TransferStoreCenterResponse
func GetTransferStoreCenterResponse() *TransferStoreCenterResponse {
	return poolTransferStoreCenterResponse.Get().(*TransferStoreCenterResponse)
}

// ReleaseTransferStoreCenterResponse 释放TransferStoreCenterResponse
func ReleaseTransferStoreCenterResponse(v *TransferStoreCenterResponse) {
	v.ZipCode = ""
	v.Country = ""
	v.City = ""
	v.CompanyName = ""
	v.Province = ""
	v.Phone = ""
	v.Name = ""
	v.DetailAddress = ""
	v.StoreName = ""
	v.Email = ""
	v.StoreCode = ""
	v.CountryId = 0
	v.AreaId = 0
	poolTransferStoreCenterResponse.Put(v)
}
