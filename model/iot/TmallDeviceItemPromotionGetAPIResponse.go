package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallDeviceItemPromotionGetAPIResponse 智能硬件上商品优惠获取 API返回值
// tmall.device.item.promotion.get
//
// 商品优惠详情查询，可查询商品设置的详细优惠。包括限时折扣，满就送等官方优惠以及第三方优惠。
type TmallDeviceItemPromotionGetAPIResponse struct {
	model.CommonResponse
	TmallDeviceItemPromotionGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallDeviceItemPromotionGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallDeviceItemPromotionGetAPIResponseModel).Reset()
}

// TmallDeviceItemPromotionGetAPIResponseModel is 智能硬件上商品优惠获取 成功返回结果
type TmallDeviceItemPromotionGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_device_item_promotion_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 优惠详细信息
	Promotions *PromotionDisplayTop `json:"promotions,omitempty" xml:"promotions,omitempty"`
}

// Reset 清空结构体
func (m *TmallDeviceItemPromotionGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Promotions = nil
}

var poolTmallDeviceItemPromotionGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallDeviceItemPromotionGetAPIResponse)
	},
}

// GetTmallDeviceItemPromotionGetAPIResponse 从 sync.Pool 获取 TmallDeviceItemPromotionGetAPIResponse
func GetTmallDeviceItemPromotionGetAPIResponse() *TmallDeviceItemPromotionGetAPIResponse {
	return poolTmallDeviceItemPromotionGetAPIResponse.Get().(*TmallDeviceItemPromotionGetAPIResponse)
}

// ReleaseTmallDeviceItemPromotionGetAPIResponse 将 TmallDeviceItemPromotionGetAPIResponse 保存到 sync.Pool
func ReleaseTmallDeviceItemPromotionGetAPIResponse(v *TmallDeviceItemPromotionGetAPIResponse) {
	v.Reset()
	poolTmallDeviceItemPromotionGetAPIResponse.Put(v)
}
