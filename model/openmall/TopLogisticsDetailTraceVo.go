package openmall

// TopLogisticsDetailTraceVo 结构体
type TopLogisticsDetailTraceVo struct {
	// 物流公司名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 运单号
	OutSid string `json:"out_sid,omitempty" xml:"out_sid,omitempty"`
	// 订单的物流状态（仅支持线上发货online订单，线下发货offline发出后直接变为已签收，OpenMall场景如无法判断请直接忽略，直接获取最后一个trace节点描述） * 等候发送给物流公司 *已提交给物流公司,等待物流公司接单 *已经确认消息接收，等待物流公司接单 *物流公司已接单 *物流公司不接单 *物流公司揽收失败 *物流公司揽收成功 *签收失败 *对方已签收 *对方拒绝签收
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 淘宝交易单ID
	Tid string `json:"tid,omitempty" xml:"tid,omitempty"`
	// 流转信息列表
	TraceList []TransitStepInfoVo `json:"trace_list,omitempty" xml:"trace_list>transit_step_info_vo,omitempty"`
}
