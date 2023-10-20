package servicecenter

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofuwupurchaseorderconfirmAPIResponse 服务市场内购服务下单接口 API返回值
// taobao.fuwu.purchase.order.confirm
//
// 通过传入服务市场商品的itemcode等信息，返回给服务商内购服务的下单链接
type TaobaofuwupurchaseorderconfirmAPIResponse struct {
	model.CommonResponse
	TaobaofuwupurchaseorderconfirmAPIResponseModel
}

// TaobaofuwupurchaseorderconfirmAPIResponseModel is 服务市场内购服务下单接口 成功返回结果
type TaobaofuwupurchaseorderconfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"fuwu_purchase_order_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 下单页面url
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}
