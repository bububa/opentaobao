package traveltrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTravelTicketOrderRefundAPIResponse 飞猪门票退票结果通知 API返回值
// taobao.travel.ticket.order.refund
//
// 门票系统商通过TOP接口通知飞猪门票是否退票成功，以及退票数量。
type TaobaoTravelTicketOrderRefundAPIResponse struct {
	model.CommonResponse
	TaobaoTravelTicketOrderRefundAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTravelTicketOrderRefundAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTravelTicketOrderRefundAPIResponseModel).Reset()
}

// TaobaoTravelTicketOrderRefundAPIResponseModel is 飞猪门票退票结果通知 成功返回结果
type TaobaoTravelTicketOrderRefundAPIResponseModel struct {
	XMLName xml.Name `xml:"travel_ticket_order_refund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTravelTicketOrderRefundAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoTravelTicketOrderRefundAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTravelTicketOrderRefundAPIResponse)
	},
}

// GetTaobaoTravelTicketOrderRefundAPIResponse 从 sync.Pool 获取 TaobaoTravelTicketOrderRefundAPIResponse
func GetTaobaoTravelTicketOrderRefundAPIResponse() *TaobaoTravelTicketOrderRefundAPIResponse {
	return poolTaobaoTravelTicketOrderRefundAPIResponse.Get().(*TaobaoTravelTicketOrderRefundAPIResponse)
}

// ReleaseTaobaoTravelTicketOrderRefundAPIResponse 将 TaobaoTravelTicketOrderRefundAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTravelTicketOrderRefundAPIResponse(v *TaobaoTravelTicketOrderRefundAPIResponse) {
	v.Reset()
	poolTaobaoTravelTicketOrderRefundAPIResponse.Put(v)
}
