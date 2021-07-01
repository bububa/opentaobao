package openmall

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenmallTradeRenderAPIResponse
渲染订单价格 API返回值
taobao.openmall.trade.render

请求渲染订单价格 */
type TaobaoOpenmallTradeRenderAPIResponse struct {
	model.CommonResponse
	TaobaoOpenmallTradeRenderAPIResponseModel
}

// TaobaoOpenmallTradeRenderAPIResponseModel is 渲染订单价格 成功返回结果
type TaobaoOpenmallTradeRenderAPIResponseModel struct {
	XMLName xml.Name `xml:"openmall_trade_render_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TopTradeResultVo `json:"result,omitempty" xml:"result,omitempty"`
}
