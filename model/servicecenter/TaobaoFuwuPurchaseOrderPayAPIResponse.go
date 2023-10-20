package servicecenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFuwuPurchaseOrderPayAPIResponse 内购服务订单付款页获取接口 API返回值
// taobao.fuwu.purchase.order.pay
//
// 通过接口获取某一订单的付款页面链接
type TaobaoFuwuPurchaseOrderPayAPIResponse struct {
	model.CommonResponse
	TaobaoFuwuPurchaseOrderPayAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFuwuPurchaseOrderPayAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFuwuPurchaseOrderPayAPIResponseModel).Reset()
}

// TaobaoFuwuPurchaseOrderPayAPIResponseModel is 内购服务订单付款页获取接口 成功返回结果
type TaobaoFuwuPurchaseOrderPayAPIResponseModel struct {
	XMLName xml.Name `xml:"fuwu_purchase_order_pay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 该url用于订单付款
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFuwuPurchaseOrderPayAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Url = ""
}

var poolTaobaoFuwuPurchaseOrderPayAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFuwuPurchaseOrderPayAPIResponse)
	},
}

// GetTaobaoFuwuPurchaseOrderPayAPIResponse 从 sync.Pool 获取 TaobaoFuwuPurchaseOrderPayAPIResponse
func GetTaobaoFuwuPurchaseOrderPayAPIResponse() *TaobaoFuwuPurchaseOrderPayAPIResponse {
	return poolTaobaoFuwuPurchaseOrderPayAPIResponse.Get().(*TaobaoFuwuPurchaseOrderPayAPIResponse)
}

// ReleaseTaobaoFuwuPurchaseOrderPayAPIResponse 将 TaobaoFuwuPurchaseOrderPayAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFuwuPurchaseOrderPayAPIResponse(v *TaobaoFuwuPurchaseOrderPayAPIResponse) {
	v.Reset()
	poolTaobaoFuwuPurchaseOrderPayAPIResponse.Put(v)
}
