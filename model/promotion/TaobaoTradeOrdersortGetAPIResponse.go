package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTradeOrdersortGetAPIResponse
获取前N有礼活动的开奖订单列表 API返回值
taobao.trade.ordersort.get

获取前N有礼活动的开奖订单列表 */
type TaobaoTradeOrdersortGetAPIResponse struct {
	model.CommonResponse
	TaobaoTradeOrdersortGetAPIResponseModel
}

// TaobaoTradeOrdersortGetAPIResponseModel is 获取前N有礼活动的开奖订单列表 成功返回结果
type TaobaoTradeOrdersortGetAPIResponseModel struct {
	XMLName xml.Name `xml:"trade_ordersort_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoTradeOrdersortGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
