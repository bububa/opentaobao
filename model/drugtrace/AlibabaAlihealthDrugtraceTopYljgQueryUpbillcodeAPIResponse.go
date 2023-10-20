package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIResponse 通过一个码，查询这个码对应的上游企业出库单的单据号 API返回值
// alibaba.alihealth.drugtrace.top.yljg.query.upbillcode
//
// 一个查询上游出库单号的接口。企业在扫码入库时，接口通过扫到的码判定这个码对应的上游企业所属的出库单据号
type AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIResponseModel is 通过一个码，查询这个码对应的上游企业出库单的单据号 成功返回结果
type AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_yljg_query_upbillcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIResponse)
	},
}

// GetAlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIResponse
func GetAlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIResponse() *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIResponse {
	return poolAlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIResponse.Get().(*AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIResponse)
}

// ReleaseAlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIResponse 将 AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIResponse(v *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIResponse.Put(v)
}
