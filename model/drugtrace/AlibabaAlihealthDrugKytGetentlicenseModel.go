package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytGetentlicenseModel 结构体
type AlibabaAlihealthDrugKytGetentlicenseModel struct {
	// 资质名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 资质描述
	DescInfo string `json:"desc_info,omitempty" xml:"desc_info,omitempty"`
	// 主键ID
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 企业ID
	RefEntId string `json:"ref_ent_id,omitempty" xml:"ref_ent_id,omitempty"`
	// 过期时间
	ExpireDate string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
	// 资质图片存储地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 资质编号
	LicenseNo string `json:"license_no,omitempty" xml:"license_no,omitempty"`
	// 资质类型
	LicenseType string `json:"license_type,omitempty" xml:"license_type,omitempty"`
	// 资质图片访问地址
	TruthUrl string `json:"truth_url,omitempty" xml:"truth_url,omitempty"`
	// 是否必须
	Requisite int64 `json:"requisite,omitempty" xml:"requisite,omitempty"`
}

var poolAlibabaAlihealthDrugKytGetentlicenseModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytGetentlicenseModel)
	},
}

// GetAlibabaAlihealthDrugKytGetentlicenseModel() 从对象池中获取AlibabaAlihealthDrugKytGetentlicenseModel
func GetAlibabaAlihealthDrugKytGetentlicenseModel() *AlibabaAlihealthDrugKytGetentlicenseModel {
	return poolAlibabaAlihealthDrugKytGetentlicenseModel.Get().(*AlibabaAlihealthDrugKytGetentlicenseModel)
}

// ReleaseAlibabaAlihealthDrugKytGetentlicenseModel 释放AlibabaAlihealthDrugKytGetentlicenseModel
func ReleaseAlibabaAlihealthDrugKytGetentlicenseModel(v *AlibabaAlihealthDrugKytGetentlicenseModel) {
	v.Name = ""
	v.DescInfo = ""
	v.Id = ""
	v.RefEntId = ""
	v.ExpireDate = ""
	v.Url = ""
	v.GmtModified = ""
	v.LicenseNo = ""
	v.LicenseType = ""
	v.TruthUrl = ""
	v.Requisite = 0
	poolAlibabaAlihealthDrugKytGetentlicenseModel.Put(v)
}
