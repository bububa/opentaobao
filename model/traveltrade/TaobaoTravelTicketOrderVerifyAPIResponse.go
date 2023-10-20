package traveltrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTravelTicketOrderVerifyAPIResponse 飞猪门票核销通知 API返回值
// taobao.travel.ticket.order.verify
//
// 系统商通过TOP接口调用通知飞猪门票核销情况
type TaobaoTravelTicketOrderVerifyAPIResponse struct {
	model.CommonResponse
	TaobaoTravelTicketOrderVerifyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTravelTicketOrderVerifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTravelTicketOrderVerifyAPIResponseModel).Reset()
}

// TaobaoTravelTicketOrderVerifyAPIResponseModel is 飞猪门票核销通知 成功返回结果
type TaobaoTravelTicketOrderVerifyAPIResponseModel struct {
	XMLName xml.Name `xml:"travel_ticket_order_verify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功状态true or false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTravelTicketOrderVerifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoTravelTicketOrderVerifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTravelTicketOrderVerifyAPIResponse)
	},
}

// GetTaobaoTravelTicketOrderVerifyAPIResponse 从 sync.Pool 获取 TaobaoTravelTicketOrderVerifyAPIResponse
func GetTaobaoTravelTicketOrderVerifyAPIResponse() *TaobaoTravelTicketOrderVerifyAPIResponse {
	return poolTaobaoTravelTicketOrderVerifyAPIResponse.Get().(*TaobaoTravelTicketOrderVerifyAPIResponse)
}

// ReleaseTaobaoTravelTicketOrderVerifyAPIResponse 将 TaobaoTravelTicketOrderVerifyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTravelTicketOrderVerifyAPIResponse(v *TaobaoTravelTicketOrderVerifyAPIResponse) {
	v.Reset()
	poolTaobaoTravelTicketOrderVerifyAPIResponse.Put(v)
}
