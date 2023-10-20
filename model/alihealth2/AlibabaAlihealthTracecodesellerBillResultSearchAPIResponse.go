package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTracecodesellerBillResultSearchAPIResponse 查询出入库单处理结果api API返回值
// alibaba.alihealth.tracecodeseller.bill.result.search
//
// 查询出入库单处理结果api
type AlibabaAlihealthTracecodesellerBillResultSearchAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthTracecodesellerBillResultSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthTracecodesellerBillResultSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthTracecodesellerBillResultSearchAPIResponseModel).Reset()
}

// AlibabaAlihealthTracecodesellerBillResultSearchAPIResponseModel is 查询出入库单处理结果api 成功返回结果
type AlibabaAlihealthTracecodesellerBillResultSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_tracecodeseller_bill_result_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthTracecodesellerBillResultSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthTracecodesellerBillResultSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthTracecodesellerBillResultSearchAPIResponse)
	},
}

// GetAlibabaAlihealthTracecodesellerBillResultSearchAPIResponse 从 sync.Pool 获取 AlibabaAlihealthTracecodesellerBillResultSearchAPIResponse
func GetAlibabaAlihealthTracecodesellerBillResultSearchAPIResponse() *AlibabaAlihealthTracecodesellerBillResultSearchAPIResponse {
	return poolAlibabaAlihealthTracecodesellerBillResultSearchAPIResponse.Get().(*AlibabaAlihealthTracecodesellerBillResultSearchAPIResponse)
}

// ReleaseAlibabaAlihealthTracecodesellerBillResultSearchAPIResponse 将 AlibabaAlihealthTracecodesellerBillResultSearchAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthTracecodesellerBillResultSearchAPIResponse(v *AlibabaAlihealthTracecodesellerBillResultSearchAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthTracecodesellerBillResultSearchAPIResponse.Put(v)
}
