package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIResponse 上传单据后处理状态查询 API返回值
// alibaba.alihealth.drugtrace.top.lsyd.query.billstatus
//
// 单据处理状态查询
type AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIResponseModel
}

// AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIResponseModel is 上传单据后处理状态查询 成功返回结果
type AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_lsyd_query_billstatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
