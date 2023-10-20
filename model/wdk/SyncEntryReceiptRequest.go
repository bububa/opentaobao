package wdk

import (
	"sync"
)

// SyncEntryReceiptRequest 结构体
type SyncEntryReceiptRequest struct {
	// 地址信息
	AddressInfo []AddressInfo `json:"address_info,omitempty" xml:"address_info>address_info,omitempty"`
	// 联系人信息
	ContactInfo []ContactInfo `json:"contact_info,omitempty" xml:"contact_info>contact_info,omitempty"`
	// 教育经历
	EducationExpInfo []EducationExpInfo `json:"education_exp_info,omitempty" xml:"education_exp_info>education_exp_info,omitempty"`
	// 家庭信息
	FamilyInfo []FamilyInfo `json:"family_info,omitempty" xml:"family_info>family_info,omitempty"`
	// 工作履历
	JobExpInfo []JobExpInfo `json:"job_exp_info,omitempty" xml:"job_exp_info>job_exp_info,omitempty"`
	// 证件信息
	OfferLicenseInfo []OfferLicenseInfo `json:"offer_license_info,omitempty" xml:"offer_license_info>offer_license_info,omitempty"`
	// 关联方申报信息
	RelatedPartyInfo []RelatedPartyInfo `json:"related_party_info,omitempty" xml:"related_party_info>related_party_info,omitempty"`
	// 语言信息
	LanguageInfo []LanguageInfo `json:"language_info,omitempty" xml:"language_info>language_info,omitempty"`
	// 员工主要信息
	EmployeeBasic *EmployeeBasic `json:"employee_basic,omitempty" xml:"employee_basic,omitempty"`
}

var poolSyncEntryReceiptRequest = sync.Pool{
	New: func() any {
		return new(SyncEntryReceiptRequest)
	},
}

// GetSyncEntryReceiptRequest() 从对象池中获取SyncEntryReceiptRequest
func GetSyncEntryReceiptRequest() *SyncEntryReceiptRequest {
	return poolSyncEntryReceiptRequest.Get().(*SyncEntryReceiptRequest)
}

// ReleaseSyncEntryReceiptRequest 释放SyncEntryReceiptRequest
func ReleaseSyncEntryReceiptRequest(v *SyncEntryReceiptRequest) {
	v.AddressInfo = v.AddressInfo[:0]
	v.ContactInfo = v.ContactInfo[:0]
	v.EducationExpInfo = v.EducationExpInfo[:0]
	v.FamilyInfo = v.FamilyInfo[:0]
	v.JobExpInfo = v.JobExpInfo[:0]
	v.OfferLicenseInfo = v.OfferLicenseInfo[:0]
	v.RelatedPartyInfo = v.RelatedPartyInfo[:0]
	v.LanguageInfo = v.LanguageInfo[:0]
	v.EmployeeBasic = nil
	poolSyncEntryReceiptRequest.Put(v)
}
