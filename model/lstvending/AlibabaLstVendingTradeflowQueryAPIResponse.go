package lstvending

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstVendingTradeflowQueryAPIResponse 自动售卖机交易流水查询 API返回值
// alibaba.lst.vending.tradeflow.query
//
// 零售通自动售卖机交易流水查询接口，品牌商通过此接口同步商品交易数据。
type AlibabaLstVendingTradeflowQueryAPIResponse struct {
	model.CommonResponse
	AlibabaLstVendingTradeflowQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstVendingTradeflowQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstVendingTradeflowQueryAPIResponseModel).Reset()
}

// AlibabaLstVendingTradeflowQueryAPIResponseModel is 自动售卖机交易流水查询 成功返回结果
type AlibabaLstVendingTradeflowQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_vending_tradeflow_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果集
	Result *AlibabaLstVendingTradeflowQueryResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstVendingTradeflowQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLstVendingTradeflowQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstVendingTradeflowQueryAPIResponse)
	},
}

// GetAlibabaLstVendingTradeflowQueryAPIResponse 从 sync.Pool 获取 AlibabaLstVendingTradeflowQueryAPIResponse
func GetAlibabaLstVendingTradeflowQueryAPIResponse() *AlibabaLstVendingTradeflowQueryAPIResponse {
	return poolAlibabaLstVendingTradeflowQueryAPIResponse.Get().(*AlibabaLstVendingTradeflowQueryAPIResponse)
}

// ReleaseAlibabaLstVendingTradeflowQueryAPIResponse 将 AlibabaLstVendingTradeflowQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstVendingTradeflowQueryAPIResponse(v *AlibabaLstVendingTradeflowQueryAPIResponse) {
	v.Reset()
	poolAlibabaLstVendingTradeflowQueryAPIResponse.Put(v)
}
