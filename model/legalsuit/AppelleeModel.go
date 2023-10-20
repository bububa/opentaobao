package legalsuit

import (
	"sync"
)

// AppelleeModel 结构体
type AppelleeModel struct {
	// 承办律师姓名
	LawyerName string `json:"lawyer_name,omitempty" xml:"lawyer_name,omitempty"`
	// 承办律所名称
	LawFirmName string `json:"law_firm_name,omitempty" xml:"law_firm_name,omitempty"`
	// 住所地
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 证件编号
	CertifyNumber string `json:"certify_number,omitempty" xml:"certify_number,omitempty"`
	// 证件类型
	CertifyType string `json:"certify_type,omitempty" xml:"certify_type,omitempty"`
	// 联系方式
	Contact string `json:"contact,omitempty" xml:"contact,omitempty"`
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 承办律师联系方式
	LawyerContact string `json:"lawyer_contact,omitempty" xml:"lawyer_contact,omitempty"`
	// 序号
	SerialNumber int64 `json:"serial_number,omitempty" xml:"serial_number,omitempty"`
	// 是否为集团公司
	IsAlibabaCompany bool `json:"is_alibaba_company,omitempty" xml:"is_alibaba_company,omitempty"`
}

var poolAppelleeModel = sync.Pool{
	New: func() any {
		return new(AppelleeModel)
	},
}

// GetAppelleeModel() 从对象池中获取AppelleeModel
func GetAppelleeModel() *AppelleeModel {
	return poolAppelleeModel.Get().(*AppelleeModel)
}

// ReleaseAppelleeModel 释放AppelleeModel
func ReleaseAppelleeModel(v *AppelleeModel) {
	v.LawyerName = ""
	v.LawFirmName = ""
	v.Address = ""
	v.CertifyNumber = ""
	v.CertifyType = ""
	v.Contact = ""
	v.Name = ""
	v.LawyerContact = ""
	v.SerialNumber = 0
	v.IsAlibabaCompany = false
	poolAppelleeModel.Put(v)
}
