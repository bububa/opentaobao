package alidoc

import (
	"sync"
)

// DrugStoreAddTopRequest 结构体
type DrugStoreAddTopRequest struct {
	// 药店地址
	DrugStoreAddress string `json:"drug_store_address,omitempty" xml:"drug_store_address,omitempty"`
	// 维度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 电话
	DrugStrorePhone string `json:"drug_strore_phone,omitempty" xml:"drug_strore_phone,omitempty"`
	// 药店编码
	DrugStoreCode string `json:"drug_store_code,omitempty" xml:"drug_store_code,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 药店名称
	DrugStoreName string `json:"drug_store_name,omitempty" xml:"drug_store_name,omitempty"`
	// 医保标签
	MedicareLabel int64 `json:"medicare_label,omitempty" xml:"medicare_label,omitempty"`
}

var poolDrugStoreAddTopRequest = sync.Pool{
	New: func() any {
		return new(DrugStoreAddTopRequest)
	},
}

// GetDrugStoreAddTopRequest() 从对象池中获取DrugStoreAddTopRequest
func GetDrugStoreAddTopRequest() *DrugStoreAddTopRequest {
	return poolDrugStoreAddTopRequest.Get().(*DrugStoreAddTopRequest)
}

// ReleaseDrugStoreAddTopRequest 释放DrugStoreAddTopRequest
func ReleaseDrugStoreAddTopRequest(v *DrugStoreAddTopRequest) {
	v.DrugStoreAddress = ""
	v.Latitude = ""
	v.DrugStrorePhone = ""
	v.DrugStoreCode = ""
	v.Longitude = ""
	v.DrugStoreName = ""
	v.MedicareLabel = 0
	poolDrugStoreAddTopRequest.Put(v)
}
