package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytDrVaequipmentListModel 结构体
type AlibabaAlihealthDrugKytDrVaequipmentListModel struct {
	// 查询列表
	List []AlibabaAlihealthDrugKytDrVaequipmentListResult `json:"list,omitempty" xml:"list>alibaba_alihealth_drug_kyt_dr_vaequipment_list_result,omitempty"`
	// 页数
	Pages int64 `json:"pages,omitempty" xml:"pages,omitempty"`
	// 总数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 每页显示数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 页码
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
}

var poolAlibabaAlihealthDrugKytDrVaequipmentListModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrVaequipmentListModel)
	},
}

// GetAlibabaAlihealthDrugKytDrVaequipmentListModel() 从对象池中获取AlibabaAlihealthDrugKytDrVaequipmentListModel
func GetAlibabaAlihealthDrugKytDrVaequipmentListModel() *AlibabaAlihealthDrugKytDrVaequipmentListModel {
	return poolAlibabaAlihealthDrugKytDrVaequipmentListModel.Get().(*AlibabaAlihealthDrugKytDrVaequipmentListModel)
}

// ReleaseAlibabaAlihealthDrugKytDrVaequipmentListModel 释放AlibabaAlihealthDrugKytDrVaequipmentListModel
func ReleaseAlibabaAlihealthDrugKytDrVaequipmentListModel(v *AlibabaAlihealthDrugKytDrVaequipmentListModel) {
	v.List = v.List[:0]
	v.Pages = 0
	v.TotalNum = 0
	v.PageSize = 0
	v.Page = 0
	poolAlibabaAlihealthDrugKytDrVaequipmentListModel.Put(v)
}
