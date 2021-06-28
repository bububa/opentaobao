package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取前N有礼活动的开奖订单列表 APIResponse
taobao.trade.ordersort.get

获取前N有礼活动的开奖订单列表
*/
type TaobaoTradeOrdersortGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"trade_ordersort_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *TaobaoTradeOrdersortGetResult `json:"result,omitempty" xml:"