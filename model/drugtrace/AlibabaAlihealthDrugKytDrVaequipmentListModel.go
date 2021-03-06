package drugtrace

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
