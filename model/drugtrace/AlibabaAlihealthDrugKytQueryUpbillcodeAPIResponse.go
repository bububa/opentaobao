package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytQueryUpbillcodeAPIResponse 通过一个码查询上游出库单 API返回值
// alibaba.alihealth.drug.kyt.query.upbillcode
//
// 一个查询上游出库单号的接口。企业在扫码入库时，接口通过扫到的码判定这个码对应的上游企业所属的出库单据号
type AlibabaAlihealthDrugKytQueryUpbillcodeAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytQueryUpbillcodeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytQueryUpbillcodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytQueryUpbillcodeAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytQueryUpbillcodeAPIResponseModel is 通过一个码查询上游出库单 成功返回结果
type AlibabaAlihealthDrugKytQueryUpbillcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_query_upbillcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugKytQueryUpbillcodeResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytQueryUpbillcodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytQueryUpbillcodeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytQueryUpbillcodeAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytQueryUpbillcodeAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytQueryUpbillcodeAPIResponse
func GetAlibabaAlihealthDrugKytQueryUpbillcodeAPIResponse() *AlibabaAlihealthDrugKytQueryUpbillcodeAPIResponse {
	return poolAlibabaAlihealthDrugKytQueryUpbillcodeAPIResponse.Get().(*AlibabaAlihealthDrugKytQueryUpbillcodeAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytQueryUpbillcodeAPIResponse 将 AlibabaAlihealthDrugKytQueryUpbillcodeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytQueryUpbillcodeAPIResponse(v *AlibabaAlihealthDrugKytQueryUpbillcodeAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytQueryUpbillcodeAPIResponse.Put(v)
}
