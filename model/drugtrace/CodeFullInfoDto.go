package drugtrace

// CodeFullInfoDto 结构体
type CodeFullInfoDto struct {
	// 字段校验，1、本位码为空/疑似有误：空、连续7位数（含）为同一数字、循环数字； 2、有效期至疑似有误：不符合平台任意效期计算规则的数据（1.生产日期+有效期后的当天；2.生产日期+有效期后的前一天；3.生产日期+有效期后的上个月的最后一天；4.生产日期+有效期后的本月最后一天）；3、制剂规格疑似有误：被标记有误的数据
	ValidationRuleList []Validationruledtos `json:"validation_rule_list,omitempty" xml:"validation_rule_list>validationruledtos,omitempty"`
	// 单据流向信息
	BillInOutDTOS []BillInOutDto `json:"bill_in_out_d_t_o_s,omitempty" xml:"bill_in_out_d_t_o_s>bill_in_out_dto,omitempty"`
	// 表示成功
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 追溯码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 码生产信息对象
	CodeProduceInfoDTO *CodeProduceInfoDto `json:"code_produce_info_d_t_o,omitempty" xml:"code_produce_info_d_t_o,omitempty"`
	// 药品基本信息对象
	DrugEntBaseDTO *DrugEntBaseDto `json:"drug_ent_base_d_t_o,omitempty" xml:"drug_ent_base_d_t_o,omitempty"`
	// 企业信息对象
	PUserEntDTO *PUserEntDto `json:"p_user_ent_d_t_o,omitempty" xml:"p_user_ent_d_t_o,omitempty"`
	// 追溯码状态对象
	CodeStatusTypeDTO *CodeStatusTypeDto `json:"code_status_type_d_t_o,omitempty" xml:"code_status_type_d_t_o,omitempty"`
	// 码包装层级
	PackageLevel int64 `json:"package_level,omitempty" xml:"package_level,omitempty"`
	// 码激活信息
	CodeActiveInfoListApiDTO *CodeActiveInfoListApiDto `json:"code_active_info_list_api_d_t_o,omitempty" xml:"code_active_info_list_api_d_t_o,omitempty"`
	// 结果
	Model *CodeFullInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 调用结果
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}
