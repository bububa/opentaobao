package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIResponse 摩卡确认入职后往入职单据表写数据接口 API返回值
// alibaba.wdk.hrworkbench.moka.entry.receipt.write
//
// 摩卡确认入职后往入职单据表写数据接口
type AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIResponse struct {
	model.CommonResponse
	AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIResponseModel).Reset()
}

// AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIResponseModel is 摩卡确认入职后往入职单据表写数据接口 成功返回结果
type AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_hrworkbench_moka_entry_receipt_write_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// trace_id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// data
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 失败
	Fail bool `json:"fail,omitempty" xml:"fail,omitempty"`
	// 成功并且结果非空
	SuccAndNotNull bool `json:"succ_and_not_null,omitempty" xml:"succ_and_not_null,omitempty"`
	// 成功结果为空
	SuccAndNull bool `json:"succ_and_null,omitempty" xml:"succ_and_null,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.TraceId = ""
	m.Data = false
	m.Fail = false
	m.SuccAndNotNull = false
	m.SuccAndNull = false
}

var poolAlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIResponse)
	},
}

// GetAlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIResponse 从 sync.Pool 获取 AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIResponse
func GetAlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIResponse() *AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIResponse {
	return poolAlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIResponse.Get().(*AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIResponse)
}

// ReleaseAlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIResponse 将 AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIResponse(v *AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIResponse) {
	v.Reset()
	poolAlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIResponse.Put(v)
}
