package normalvisa

// NormalVisaDetailInfo 结构体
type NormalVisaDetailInfo struct {
	// 用户信息数组
	NVisaDetailPersonResultVOList []NormalVisaPersonDetailVo `json:"n_visa_detail_person_result_v_o_list,omitempty" xml:"n_visa_detail_person_result_v_o_list>normal_visa_person_detail_vo,omitempty"`
	// 结束状态描述
	EndStatusDesc string `json:"end_status_desc,omitempty" xml:"end_status_desc,omitempty"`
	// 结束状态
	EndStatus int64 `json:"end_status,omitempty" xml:"end_status,omitempty"`
	// 1:贴纸签 2:电子签 3:面试
	VisaType int64 `json:"visa_type,omitempty" xml:"visa_type,omitempty"`
}
