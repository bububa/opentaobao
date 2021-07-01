package jst

// O2oAppointInfoDto 结构体
type O2oAppointInfoDto struct {
	// 预约信息唯一编码，如果使用此参数，customer_nick和mall_code参数不会使用
	AppointCode int64 `json:"appoint_code,omitempty" xml:"appoint_code,omitempty"`
	// 预约属性信息列表
	AppointInfoParams []O2oAppointInfoParam `json:"appoint_info_params,omitempty" xml:"appoint_info_params>o2o_appoint_info_param,omitempty"`
}
