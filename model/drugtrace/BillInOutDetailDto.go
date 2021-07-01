package drugtrace

// BillInOutDetailDto 结构体
type BillInOutDetailDto struct {
	// 修改时间
	ModDate string `json:"mod_date,omitempty" xml:"mod_date,omitempty"`
	// 处理时间
	ProcessDate string `json:"process_date,omitempty" xml:"process_date,omitempty"`
	// 单据日期
	BillTime string `json:"bill_time,omitempty" xml:"bill_time,omitempty"`
	// 收货企业id
	ToUserId string `json:"to_user_id,omitempty" xml:"to_user_id,omitempty"`
	// 收货企业名称
	ToEntName string `json:"to_ent_name,omitempty" xml:"to_ent_name,omitempty"`
	// 发货企业id
	FromUserId string `json:"from_user_id,omitempty" xml:"from_user_id,omitempty"`
	// 发货企业名称
	FromEntName string `json:"from_ent_name,omitempty" xml:"from_ent_name,omitempty"`
	// 单据类型名称
	BillTypeName string `json:"bill_type_name,omitempty" xml:"bill_type_name,omitempty"`
	// 单据类型
	BillType string `json:"bill_type,omitempty" xml:"bill_type,omitempty"`
	// 单据号码
	BillCode string `json:"bill_code,omitempty" xml:"bill_code,omitempty"`
	// 单据详情
	BillChkInOutDetailListDTOList []Billchkinoutdetaillistdtolist `json:"bill_chk_in_out_detail_list_d_t_o_list,omitempty" xml:"bill_chk_in_out_detail_list_d_t_o_list>billchkinoutdetaillistdtolist,omitempty"`
	// 上传文件名称
	UploadFileName string `json:"upload_file_name,omitempty" xml:"upload_file_name,omitempty"`
}
