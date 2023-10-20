package servicecenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFuwuPurchaseOrderConfirmAPIResponse 服务市场内购服务下单接口 API返回值
// taobao.fuwu.purchase.order.confirm
//
// 通过传入服务市场商品的itemcode等信息，返回给服务商内购服务的下单链接
type TaobaoFuwuPurchaseOrderConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoFuwuPurchaseOrderConfirmAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFuwuPurchaseOrderConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFuwuPurchaseOrderConfirmAPIResponseModel).Reset()
}

// TaobaoFuwuPurchaseOrderConfirmAPIResponseModel is 服务市场内购服务下单接口 成功返回结果
type TaobaoFuwuPurchaseOrderConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"fuwu_purchase_order_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 下单页面url
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFuwuPurchaseOrderConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Url = ""
}

var poolTaobaoFuwuPurchaseOrderConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFuwuPurchaseOrderConfirmAPIResponse)
	},
}

// GetTaobaoFuwuPurchaseOrderConfirmAPIResponse 从 sync.Pool 获取 TaobaoFuwuPurchaseOrderConfirmAPIResponse
func GetTaobaoFuwuPurchaseOrderConfirmAPIResponse() *TaobaoFuwuPurchaseOrderConfirmAPIResponse {
	return poolTaobaoFuwuPurchaseOrderConfirmAPIResponse.Get().(*TaobaoFuwuPurchaseOrderConfirmAPIResponse)
}

// ReleaseTaobaoFuwuPurchaseOrderConfirmAPIResponse 将 TaobaoFuwuPurchaseOrderConfirmAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFuwuPurchaseOrderConfirmAPIResponse(v *TaobaoFuwuPurchaseOrderConfirmAPIResponse) {
	v.Reset()
	poolTaobaoFuwuPurchaseOrderConfirmAPIResponse.Put(v)
}
