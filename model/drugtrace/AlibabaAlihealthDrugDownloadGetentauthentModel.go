package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugDownloadGetentauthentModel 结构体
type AlibabaAlihealthDrugDownloadGetentauthentModel struct {
	// 企业ID
	AuthRefEntId string `json:"auth_ref_ent_id,omitempty" xml:"auth_ref_ent_id,omitempty"`
	// 企业名称
	EntName string `json:"ent_name,omitempty" xml:"ent_name,omitempty"`
	// 授权企业ID
	RefEntId string `json:"ref_ent_id,omitempty" xml:"ref_ent_id,omitempty"`
	// 授权日期
	AuthDate string `json:"auth_date,omitempty" xml:"auth_date,omitempty"`
}

var poolAlibabaAlihealthDrugDownloadGetentauthentModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugDownloadGetentauthentModel)
	},
}

// GetAlibabaAlihealthDrugDownloadGetentauthentModel() 从对象池中获取AlibabaAlihealthDrugDownloadGetentauthentModel
func GetAlibabaAlihealthDrugDownloadGetentauthentModel() *AlibabaAlihealthDrugDownloadGetentauthentModel {
	return poolAlibabaAlihealthDrugDownloadGetentauthentModel.Get().(*AlibabaAlihealthDrugDownloadGetentauthentModel)
}

// ReleaseAlibabaAlihealthDrugDownloadGetentauthentModel 释放AlibabaAlihealthDrugDownloadGetentauthentModel
func ReleaseAlibabaAlihealthDrugDownloadGetentauthentModel(v *AlibabaAlihealthDrugDownloadGetentauthentModel) {
	v.AuthRefEntId = ""
	v.EntName = ""
	v.RefEntId = ""
	v.AuthDate = ""
	poolAlibabaAlihealthDrugDownloadGetentauthentModel.Put(v)
}
