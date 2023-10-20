package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniOrderDetailAPIResponse 全渠道订单详情 API返回值
// taobao.omni.order.detail
//
// 全渠道订单详情
type TaobaoOmniOrderDetailAPIResponse struct {
	model.CommonResponse
	TaobaoOmniOrderDetailAPIResponseModel
}

// TaobaoOmniOrderDetailAPIResponseModel is 全渠道订单详情 成功返回结果
type TaobaoOmniOrderDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"omni_order_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *ResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
