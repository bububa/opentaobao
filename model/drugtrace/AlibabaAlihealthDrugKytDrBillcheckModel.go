package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytDrBillcheckModel 结构体
type AlibabaAlihealthDrugKytDrBillcheckModel struct {
	// 码验证信息
	DetailInfoList []DetailInfoList `json:"detail_info_list,omitempty" xml:"detail_info_list>detail_info_list,omitempty"`
	// 单据日期
	BillTime string `json:"bill_time,omitempty" xml:"bill_time,omitempty"`
	// 收发货企业名称
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// 验证率
	MatchedRateShow string `json:"matched_rate_show,omitempty" xml:"matched_rate_show,omitempty"`
	// 单据编号
	BillCode string `json:"bill_code,omitempty" xml:"bill_code,omitempty"`
	// 单据类型
	BillType int64 `json:"bill_type,omitempty" xml:"bill_type,omitempty"`
	// 追溯验证数量
	MatchedCount int64 `json:"matched_count,omitempty" xml:"matched_count,omitempty"`
}

var poolAlibabaAlihealthDrugKytDrBillcheckModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrBillcheckModel)
	},
}

// GetAlibabaAlihealthDrugKytDrBillcheckModel() 从对象池中获取AlibabaAlihealthDrugKytDrBillcheckModel
func GetAlibabaAlihealthDrugKytDrBillcheckModel() *AlibabaAlihealthDrugKytDrBillcheckModel {
	return poolAlibabaAlihealthDrugKytDrBillcheckModel.Get().(*AlibabaAlihealthDrugKytDrBillcheckModel)
}

// ReleaseAlibabaAlihealthDrugKytDrBillcheckModel 释放AlibabaAlihealthDrugKytDrBillcheckModel
func ReleaseAlibabaAlihealthDrugKytDrBillcheckModel(v *AlibabaAlihealthDrugKytDrBillcheckModel) {
	v.DetailInfoList = v.DetailInfoList[:0]
	v.BillTime = ""
	v.UserName = ""
	v.MatchedRateShow = ""
	v.BillCode = ""
	v.BillType = 0
	v.MatchedCount = 0
	poolAlibabaAlihealthDrugKytDrBillcheckModel.Put(v)
}
