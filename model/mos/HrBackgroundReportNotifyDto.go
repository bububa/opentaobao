package mos

// HrBackgroundReportNotifyDto 结构体
type HrBackgroundReportNotifyDto struct {
	// 下单时MOS传的唯一编码
	SourceId string `json:"source_id,omitempty" xml:"source_id,omitempty"`
	// 报告结果(1:绿灯，2:红灯， 3:黄灯，10:蓝灯)
	FinalReportResSel string `json:"final_report_res_sel,omitempty" xml:"final_report_res_sel,omitempty"`
	// 外部订单号
	OrderNo string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 获取结果时背调状态(0:未背调，3:待授权，4:背调中，5:已完成，6:系统取消，7:银泰取消，8:授权超期取消，9:无需背调)
	RecordStatus string `json:"record_status,omitempty" xml:"record_status,omitempty"`
	// 报告格式化json数据
	ReportData string `json:"report_data,omitempty" xml:"report_data,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 是否最终版本(1:最终版本 0:非终版)
	FinalVersion string `json:"final_version,omitempty" xml:"final_version,omitempty"`
	// 预计报告时间
	EstimateFinishTime string `json:"estimate_finish_time,omitempty" xml:"estimate_finish_time,omitempty"`
	// 背调公司
	BackgroundCompanyName string `json:"background_company_name,omitempty" xml:"background_company_name,omitempty"`
}
