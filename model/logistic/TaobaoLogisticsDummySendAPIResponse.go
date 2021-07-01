package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLogisticsDummySendAPIResponse
无需物流（虚拟）发货处理 API返回值
taobao.logistics.dummy.send

用户调用该接口可实现无需物流（虚拟）发货,使用该接口发货，交易订单状态会直接变成卖家已发货 */
type TaobaoLogisticsDummySendAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsDummySendAPIResponseModel
}

// TaobaoLogisticsDummySendAPIResponseModel is 无需物流（虚拟）发货处理 成功返回结果
type TaobaoLogisticsDummySendAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_dummy_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回发货是否成功is_success
	Shipping *Shipping `json:"shipping,omitempty" xml:"shipping,omitempty"`
}
