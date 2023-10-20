package aedropshiper

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressDsTradeOrderGetAPIResponse 交易订单查询 API返回值
// aliexpress.ds.trade.order.get
//
// 交易订单查询
type AliexpressDsTradeOrderGetAPIResponse struct {
	model.CommonResponse
	AliexpressDsTradeOrderGetAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressDsTradeOrderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressDsTradeOrderGetAPIResponseModel).Reset()
}

// AliexpressDsTradeOrderGetAPIResponseModel is 交易订单查询 成功返回结果
type AliexpressDsTradeOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_ds_trade_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Error message
	RspMsg string `json:"rsp_msg,omitempty" xml:"rsp_msg,omitempty"`
	// Error code
	RspCode string `json:"rsp_code,omitempty" xml:"rsp_code,omitempty"`
	// Result object
	Result *AeOrderInfoResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressDsTradeOrderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RspMsg = ""
	m.RspCode = ""
	m.Result = nil
}

var poolAliexpressDsTradeOrderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressDsTradeOrderGetAPIResponse)
	},
}

// GetAliexpressDsTradeOrderGetAPIResponse 从 sync.Pool 获取 AliexpressDsTradeOrderGetAPIResponse
func GetAliexpressDsTradeOrderGetAPIResponse() *AliexpressDsTradeOrderGetAPIResponse {
	return poolAliexpressDsTradeOrderGetAPIResponse.Get().(*AliexpressDsTradeOrderGetAPIResponse)
}

// ReleaseAliexpressDsTradeOrderGetAPIResponse 将 AliexpressDsTradeOrderGetAPIResponse 保存到 sync.Pool
func ReleaseAliexpressDsTradeOrderGetAPIResponse(v *AliexpressDsTradeOrderGetAPIResponse) {
	v.Reset()
	poolAliexpressDsTradeOrderGetAPIResponse.Put(v)
}
