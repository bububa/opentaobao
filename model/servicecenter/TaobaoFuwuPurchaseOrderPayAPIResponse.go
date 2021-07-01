package servicecenter

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFuwuPurchaseOrderPayAPIResponse
内购服务订单付款页获取接口 API返回值
taobao.fuwu.purchase.order.pay

通过接口获取某一订单的付款页面链接 */
type TaobaoFuwuPurchaseOrderPayAPIResponse struct {
	model.CommonResponse
	TaobaoFuwuPurchaseOrderPayAPIResponseModel
}

// TaobaoFuwuPurchaseOrderPayAPIResponseModel is 内购服务订单付款页获取接口 成功返回结果
type TaobaoFuwuPurchaseOrderPayAPIResponseModel struct {
	XMLName xml.Name `xml:"fuwu_purchase_order_pay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 该url用于订单付款
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}
