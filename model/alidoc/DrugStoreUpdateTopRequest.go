package alidoc

import (
	"sync"

	"github.com/bububa/opentaobao/model"
)

// DrugStoreUpdateTopRequest 结构体
type DrugStoreUpdateTopRequest struct {
	// 药店编码
	DrugStoreCode string `json:"drug_store_code,omitempty" xml:"drug_store_code,omitempty"`
	// 药店名称
	DrugStoreName string `json:"drug_store_name,omitempty" xml:"drug_store_name,omitempty"`
	// 药店地址
	DrugStoreAddress string `json:"drug_store_address,omitempty" xml:"drug_store_address,omitempty"`
	// 电话
	DrugStrorePhone string `json:"drug_strore_phone,omitempty" xml:"drug_strore_phone,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 1-支持医保，0-不支持医保
	MedicareLabel *model.File `json:"medicare_label,omitempty" xml:"medicare_label,omitempty"`
	// 状态，1-启动，0-封存
	Status *model.File `json:"status,omitempty" xml:"status,omitempty"`
}

var poolDrugStoreUpdateTopRequest = sync.Pool{
	New: func() any {
		return new(DrugStoreUpdateTopRequest)
	},
}

// GetDrugStoreUpdateTopRequest() 从对象池中获取DrugStoreUpdateTopRequest
func GetDrugStoreUpdateTopRequest() *DrugStoreUpdateTopRequest {
	return poolDrugStoreUpdateTopRequest.Get().(*DrugStoreUpdateTopRequest)
}

// ReleaseDrugStoreUpdateTopRequest 释放DrugStoreUpdateTopRequest
func ReleaseDrugStoreUpdateTopRequest(v *DrugStoreUpdateTopRequest) {
	v.DrugStoreCode = ""
	v.DrugStoreName = ""
	v.DrugStoreAddress = ""
	v.DrugStrorePhone = ""
	v.Latitude = ""
	v.Longitude = ""
	v.MedicareLabel = nil
	v.Status = nil
	poolDrugStoreUpdateTopRequest.Put(v)
}
