package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIResponse 根据单据号查询单据的详情信息【注意：查询的是本企业的单据】 API返回值
// alibaba.alihealth.drugtrace.top.yljg.query.upbilldetail
//
// 根据单据号查询单据的详情信息【注意：这个接口查询的是本企业的单据，如果是查询上游的单据明细信息，使用xxxxxxx.listupout.detail接口】。
type AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIResponseModel is 根据单据号查询单据的详情信息【注意：查询的是本企业的单据】 成功返回结果
type AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_yljg_query_upbilldetail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIResponse)
	},
}

// GetAlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIResponse
func GetAlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIResponse() *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIResponse {
	return poolAlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIResponse.Get().(*AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIResponse)
}

// ReleaseAlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIResponse 将 AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIResponse(v *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIResponse.Put(v)
}
