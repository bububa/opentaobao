package tmallnr

// NrFetchCodeQueryRespDto 结构体
type NrFetchCodeQueryRespDto struct {
	// 实体信息
	NrFetchCodeDTOList []NrFetchCodeDto `json:"nr_fetch_code_d_t_o_list,omitempty" xml:"nr_fetch_code_d_t_o_list>nr_fetch_code_dto,omitempty"`
	// 返回的订单错误信息映射
	ErrMsg *Errmsg `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
}
