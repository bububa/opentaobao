package examination

import (
	"sync"
)

// ReportOrderSummaryRequest 结构体
type ReportOrderSummaryRequest struct {
	// 咨询聊天记录
	ReportDiagnoseMessageList []Reportdiagnosemessagelist `json:"report_diagnose_message_list,omitempty" xml:"report_diagnose_message_list>reportdiagnosemessagelist,omitempty"`
	// 异常项指标
	AbnormalItemList []Abnormalitemlist `json:"abnormal_item_list,omitempty" xml:"abnormal_item_list>abnormalitemlist,omitempty"`
	// 建议复查项目
	ReInspectItem string `json:"re_inspect_item,omitempty" xml:"re_inspect_item,omitempty"`
	// 建议复查科室
	ReInspectDepartment string `json:"re_inspect_department,omitempty" xml:"re_inspect_department,omitempty"`
	// 建议复查疾病
	ReInspectDiseases string `json:"re_inspect_diseases,omitempty" xml:"re_inspect_diseases,omitempty"`
	// 外部报告解读ID
	OuterOrderId string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
	// 报告解读ID
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolReportOrderSummaryRequest = sync.Pool{
	New: func() any {
		return new(ReportOrderSummaryRequest)
	},
}

// GetReportOrderSummaryRequest() 从对象池中获取ReportOrderSummaryRequest
func GetReportOrderSummaryRequest() *ReportOrderSummaryRequest {
	return poolReportOrderSummaryRequest.Get().(*ReportOrderSummaryRequest)
}

// ReleaseReportOrderSummaryRequest 释放ReportOrderSummaryRequest
func ReleaseReportOrderSummaryRequest(v *ReportOrderSummaryRequest) {
	v.ReportDiagnoseMessageList = v.ReportDiagnoseMessageList[:0]
	v.AbnormalItemList = v.AbnormalItemList[:0]
	v.ReInspectItem = ""
	v.ReInspectDepartment = ""
	v.ReInspectDiseases = ""
	v.OuterOrderId = ""
	v.OrderId = ""
	poolReportOrderSummaryRequest.Put(v)
}
