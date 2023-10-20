package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeAPIResponse 通过一个码，查询这个码对应的上游企业出库单的单据号 API返回值
// alibaba.alihealth.drugtrace.top.lsyd.query.upbillcode
//
// 一个查询上游出库单号的接口。企业在扫码入库时，接口通过扫到的码判定这个码对应的上游企业所属的出库单据号
type AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeAPIResponseModel is 通过一个码，查询这个码对应的上游企业出库单的单据号 成功返回结果
type AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_lsyd_query_upbillcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeAPIResponse)
	},
}

// GetAlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeAPIResponse
func GetAlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeAPIResponse() *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeAPIResponse {
	return poolAlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeAPIResponse.Get().(*AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeAPIResponse)
}

// ReleaseAlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeAPIResponse 将 AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeAPIResponse(v *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeAPIResponse.Put(v)
}
