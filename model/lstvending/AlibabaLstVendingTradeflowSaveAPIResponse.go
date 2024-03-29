package lstvending

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstVendingTradeflowSaveAPIResponse 自动售卖机交易信息回流 API返回值
// alibaba.lst.vending.tradeflow.save
//
// 自动售货机交易信息同步接口，ISV通过此接口上传售货机交易信息。
type AlibabaLstVendingTradeflowSaveAPIResponse struct {
	model.CommonResponse
	AlibabaLstVendingTradeflowSaveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstVendingTradeflowSaveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstVendingTradeflowSaveAPIResponseModel).Reset()
}

// AlibabaLstVendingTradeflowSaveAPIResponseModel is 自动售卖机交易信息回流 成功返回结果
type AlibabaLstVendingTradeflowSaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_vending_tradeflow_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果集
	Result *MultiResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstVendingTradeflowSaveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLstVendingTradeflowSaveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstVendingTradeflowSaveAPIResponse)
	},
}

// GetAlibabaLstVendingTradeflowSaveAPIResponse 从 sync.Pool 获取 AlibabaLstVendingTradeflowSaveAPIResponse
func GetAlibabaLstVendingTradeflowSaveAPIResponse() *AlibabaLstVendingTradeflowSaveAPIResponse {
	return poolAlibabaLstVendingTradeflowSaveAPIResponse.Get().(*AlibabaLstVendingTradeflowSaveAPIResponse)
}

// ReleaseAlibabaLstVendingTradeflowSaveAPIResponse 将 AlibabaLstVendingTradeflowSaveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstVendingTradeflowSaveAPIResponse(v *AlibabaLstVendingTradeflowSaveAPIResponse) {
	v.Reset()
	poolAlibabaLstVendingTradeflowSaveAPIResponse.Put(v)
}
