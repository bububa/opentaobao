package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLifeReservationItemOrderChangeAPIResponse 生服购后预约单外部发起变更 API返回值
// taobao.life.reservation.item.order.change
//
// 生服购后预约单外部发起变更，例如改期、取消。目前体检场景，用户会直接联系ISV改期/取消，因此开放给ISV这块的能力
type TaobaoLifeReservationItemOrderChangeAPIResponse struct {
	model.CommonResponse
	TaobaoLifeReservationItemOrderChangeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLifeReservationItemOrderChangeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLifeReservationItemOrderChangeAPIResponseModel).Reset()
}

// TaobaoLifeReservationItemOrderChangeAPIResponseModel is 生服购后预约单外部发起变更 成功返回结果
type TaobaoLifeReservationItemOrderChangeAPIResponseModel struct {
	XMLName xml.Name `xml:"life_reservation_item_order_change_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoLifeReservationItemOrderChangeResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLifeReservationItemOrderChangeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoLifeReservationItemOrderChangeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLifeReservationItemOrderChangeAPIResponse)
	},
}

// GetTaobaoLifeReservationItemOrderChangeAPIResponse 从 sync.Pool 获取 TaobaoLifeReservationItemOrderChangeAPIResponse
func GetTaobaoLifeReservationItemOrderChangeAPIResponse() *TaobaoLifeReservationItemOrderChangeAPIResponse {
	return poolTaobaoLifeReservationItemOrderChangeAPIResponse.Get().(*TaobaoLifeReservationItemOrderChangeAPIResponse)
}

// ReleaseTaobaoLifeReservationItemOrderChangeAPIResponse 将 TaobaoLifeReservationItemOrderChangeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLifeReservationItemOrderChangeAPIResponse(v *TaobaoLifeReservationItemOrderChangeAPIResponse) {
	v.Reset()
	poolTaobaoLifeReservationItemOrderChangeAPIResponse.Put(v)
}
