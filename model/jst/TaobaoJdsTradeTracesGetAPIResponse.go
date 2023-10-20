package jst

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJdsTradeTracesGetAPIResponse 获取单条订单跟踪详情 API返回值
// taobao.jds.trade.traces.get
//
// 获取聚石塔数据共享的交易全链路信息
type TaobaoJdsTradeTracesGetAPIResponse struct {
	model.CommonResponse
	TaobaoJdsTradeTracesGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJdsTradeTracesGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJdsTradeTracesGetAPIResponseModel).Reset()
}

// TaobaoJdsTradeTracesGetAPIResponseModel is 获取单条订单跟踪详情 成功返回结果
type TaobaoJdsTradeTracesGetAPIResponseModel struct {
	XMLName xml.Name `xml:"jds_trade_traces_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 跟踪信息列表
	Traces []TradeTrace `json:"traces,omitempty" xml:"traces>trade_trace,omitempty"`
	// 在订单全链路系统中的状态(比如是否存在)
	UserStatus string `json:"user_status,omitempty" xml:"user_status,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJdsTradeTracesGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Traces = m.Traces[:0]
	m.UserStatus = ""
}

var poolTaobaoJdsTradeTracesGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJdsTradeTracesGetAPIResponse)
	},
}

// GetTaobaoJdsTradeTracesGetAPIResponse 从 sync.Pool 获取 TaobaoJdsTradeTracesGetAPIResponse
func GetTaobaoJdsTradeTracesGetAPIResponse() *TaobaoJdsTradeTracesGetAPIResponse {
	return poolTaobaoJdsTradeTracesGetAPIResponse.Get().(*TaobaoJdsTradeTracesGetAPIResponse)
}

// ReleaseTaobaoJdsTradeTracesGetAPIResponse 将 TaobaoJdsTradeTracesGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJdsTradeTracesGetAPIResponse(v *TaobaoJdsTradeTracesGetAPIResponse) {
	v.Reset()
	poolTaobaoJdsTradeTracesGetAPIResponse.Put(v)
}
