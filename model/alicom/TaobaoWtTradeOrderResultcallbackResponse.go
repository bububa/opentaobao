package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家回调接口 APIResponse
taobao.wt.trade.order.resultcallback

阿里通信定制服务，商家发货后进行调用该接口，用于自动发货并确认收货
*/
type TaobaoWtTradeOrderResultcallbackAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wt_trade_order_resultcallback_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *CommonRtnDo `json:"result,omitempty" xml:"