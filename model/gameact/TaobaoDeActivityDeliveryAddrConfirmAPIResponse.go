package gameact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDeActivityDeliveryAddrConfirmAPIResponse 用户收件地址确认 API返回值
// taobao.de.activity.delivery.addr.confirm
//
// 用户收件地址确认
type TaobaoDeActivityDeliveryAddrConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoDeActivityDeliveryAddrConfirmAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoDeActivityDeliveryAddrConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoDeActivityDeliveryAddrConfirmAPIResponseModel).Reset()
}

// TaobaoDeActivityDeliveryAddrConfirmAPIResponseModel is 用户收件地址确认 成功返回结果
type TaobaoDeActivityDeliveryAddrConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"de_activity_delivery_addr_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 更新或确认收件地址
	UpdateDeliveryAddressVo *UpdateDeliveryAddressVo `json:"update_delivery_address_vo,omitempty" xml:"update_delivery_address_vo,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoDeActivityDeliveryAddrConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.UpdateDeliveryAddressVo = nil
}

var poolTaobaoDeActivityDeliveryAddrConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoDeActivityDeliveryAddrConfirmAPIResponse)
	},
}

// GetTaobaoDeActivityDeliveryAddrConfirmAPIResponse 从 sync.Pool 获取 TaobaoDeActivityDeliveryAddrConfirmAPIResponse
func GetTaobaoDeActivityDeliveryAddrConfirmAPIResponse() *TaobaoDeActivityDeliveryAddrConfirmAPIResponse {
	return poolTaobaoDeActivityDeliveryAddrConfirmAPIResponse.Get().(*TaobaoDeActivityDeliveryAddrConfirmAPIResponse)
}

// ReleaseTaobaoDeActivityDeliveryAddrConfirmAPIResponse 将 TaobaoDeActivityDeliveryAddrConfirmAPIResponse 保存到 sync.Pool
func ReleaseTaobaoDeActivityDeliveryAddrConfirmAPIResponse(v *TaobaoDeActivityDeliveryAddrConfirmAPIResponse) {
	v.Reset()
	poolTaobaoDeActivityDeliveryAddrConfirmAPIResponse.Put(v)
}
