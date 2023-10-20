package jst

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJdsRefundTracesGetAPIResponse 获取单条退款跟踪详情 API返回值
// taobao.jds.refund.traces.get
//
// 获取聚石塔数据共享的交易全链路的退款信息
type TaobaoJdsRefundTracesGetAPIResponse struct {
	model.CommonResponse
	TaobaoJdsRefundTracesGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJdsRefundTracesGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJdsRefundTracesGetAPIResponseModel).Reset()
}

// TaobaoJdsRefundTracesGetAPIResponseModel is 获取单条退款跟踪详情 成功返回结果
type TaobaoJdsRefundTracesGetAPIResponseModel struct {
	XMLName xml.Name `xml:"jds_refund_traces_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 退款跟踪列表
	Traces []RefundTrace `json:"traces,omitempty" xml:"traces>refund_trace,omitempty"`
	// 用户在全链路系统中的状态(比如是否存在)
	UserStatus string `json:"user_status,omitempty" xml:"user_status,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJdsRefundTracesGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Traces = m.Traces[:0]
	m.UserStatus = ""
}

var poolTaobaoJdsRefundTracesGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJdsRefundTracesGetAPIResponse)
	},
}

// GetTaobaoJdsRefundTracesGetAPIResponse 从 sync.Pool 获取 TaobaoJdsRefundTracesGetAPIResponse
func GetTaobaoJdsRefundTracesGetAPIResponse() *TaobaoJdsRefundTracesGetAPIResponse {
	return poolTaobaoJdsRefundTracesGetAPIResponse.Get().(*TaobaoJdsRefundTracesGetAPIResponse)
}

// ReleaseTaobaoJdsRefundTracesGetAPIResponse 将 TaobaoJdsRefundTracesGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJdsRefundTracesGetAPIResponse(v *TaobaoJdsRefundTracesGetAPIResponse) {
	v.Reset()
	poolTaobaoJdsRefundTracesGetAPIResponse.Put(v)
}
