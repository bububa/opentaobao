package qimen

import (
	"sync"
)

// TaobaoQimenSupplierSynchronizeRequest 结构体
type TaobaoQimenSupplierSynchronizeRequest struct {
	// add|update, 必填
	ActionType string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// 供应商编码 string (50), 必填
	SupplierCode string `json:"supplierCode,omitempty" xml:"supplierCode,omitempty"`
	// 供应商名称 string (200) , 必填
	SupplierName string `json:"supplierName,omitempty" xml:"supplierName,omitempty"`
	// 联系人姓名, string (50)
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 联系电话, string (50)
	Tel string `json:"tel,omitempty" xml:"tel,omitempty"`
	// 电子邮箱, string (50)
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 国家二字码，string（50）
	CountryCode string `json:"countryCode,omitempty" xml:"countryCode,omitempty"`
	// 省份, string (50)
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 城市, string (50)
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 区域, string (50)
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 村镇, string (50)
	Town string `json:"town,omitempty" xml:"town,omitempty"`
	// 详细地址, string (200)
	DetailAddress string `json:"detailAddress,omitempty" xml:"detailAddress,omitempty"`
	// 是否有效, Y/N (默认为Y)
	IsValid string `json:"isValid,omitempty" xml:"isValid,omitempty"`
	// 备注, string (500)
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
}

var poolTaobaoQimenSupplierSynchronizeRequest = sync.Pool{
	New: func() any {
		return new(TaobaoQimenSupplierSynchronizeRequest)
	},
}

// GetTaobaoQimenSupplierSynchronizeRequest() 从对象池中获取TaobaoQimenSupplierSynchronizeRequest
func GetTaobaoQimenSupplierSynchronizeRequest() *TaobaoQimenSupplierSynchronizeRequest {
	return poolTaobaoQimenSupplierSynchronizeRequest.Get().(*TaobaoQimenSupplierSynchronizeRequest)
}

// ReleaseTaobaoQimenSupplierSynchronizeRequest 释放TaobaoQimenSupplierSynchronizeRequest
func ReleaseTaobaoQimenSupplierSynchronizeRequest(v *TaobaoQimenSupplierSynchronizeRequest) {
	v.ActionType = ""
	v.SupplierCode = ""
	v.SupplierName = ""
	v.Name = ""
	v.Tel = ""
	v.Email = ""
	v.CountryCode = ""
	v.Province = ""
	v.City = ""
	v.Area = ""
	v.Town = ""
	v.DetailAddress = ""
	v.IsValid = ""
	v.Remark = ""
	poolTaobaoQimenSupplierSynchronizeRequest.Put(v)
}
