package openmall

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenmallTradeGetAPIResponse
查询订单详情 API返回值
taobao.openmall.trade.get

查询订单详情 */
type TaobaoOpenmallTradeGetAPIResponse struct {
	model.CommonResponse
	TaobaoOpenmallTradeGetAPIResponseModel
}

// TaobaoOpenmallTradeGetAPIResponseModel is 查询订单详情 成功返回结果
type TaobaoOpenmallTradeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"openmall_trade_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TopTradeDetailVo `json:"result,omitempty" xml:"result,omitempty"`
}
