package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytGetdruglicenseModel 结构体
type AlibabaAlihealthDrugKytGetdruglicenseModel struct {
	// 资质名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 资质描述
	DescInfo string `json:"desc_info,omitempty" xml:"desc_info,omitempty"`
	// 企业ID
	RefEntId string `json:"ref_ent_id,omitempty" xml:"ref_ent_id,omitempty"`
	// 过期时间
	ExpireDate string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
	// 图片存储地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 药品ID
	DrugId string `json:"drug_id,omitempty" xml:"drug_id,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 资质编号
	LicenseNo string `json:"license_no,omitempty" xml:"license_no,omitempty"`
	// 资质图片地址
	TruthUrl string `json:"truth_url,omitempty" xml:"truth_url,omitempty"`
	// 是否必须1：是0否
	Requisite int64 `json:"requisite,omitempty" xml:"requisite,omitempty"`
	// 主键ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 资质类型
	LicenseType int64 `json:"license_type,omitempty" xml:"license_type,omitempty"`
}

var poolAlibabaAlihealthDrugKytGetdruglicenseModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytGetdruglicenseModel)
	},
}

// GetAlibabaAlihealthDrugKytGetdruglicenseModel() 从对象池中获取AlibabaAlihealthDrugKytGetdruglicenseModel
func GetAlibabaAlihealthDrugKytGetdruglicenseModel() *AlibabaAlihealthDrugKytGetdruglicenseModel {
	return poolAlibabaAlihealthDrugKytGetdruglicenseModel.Get().(*AlibabaAlihealthDrugKytGetdruglicenseModel)
}

// ReleaseAlibabaAlihealthDrugKytGetdruglicenseModel 释放AlibabaAlihealthDrugKytGetdruglicenseModel
func ReleaseAlibabaAlihealthDrugKytGetdruglicenseModel(v *AlibabaAlihealthDrugKytGetdruglicenseModel) {
	v.Name = ""
	v.DescInfo = ""
	v.RefEntId = ""
	v.ExpireDate = ""
	v.Url = ""
	v.DrugId = ""
	v.GmtModified = ""
	v.LicenseNo = ""
	v.TruthUrl = ""
	v.Requisite = 0
	v.Id = 0
	v.LicenseType = 0
	poolAlibabaAlihealthDrugKytGetdruglicenseModel.Put(v)
}
