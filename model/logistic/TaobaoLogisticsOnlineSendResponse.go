package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
在线订单发货处理（支持货到付款） APIResponse
taobao.logistics.online.send

用户调用该接口可实现在线订单发货（支持货到付款）<br/>调用该接口实现在线下单发货，有两种情况：<br><br/><font color='red'>如果不输入运单号的情况：交易状态不会改变，需要调用taobao.logistics.online.confirm确认发货后交易状态才会变成卖家已发货。<br><br/>如果输入运单号的情况发货：交易订单状态会直接变成卖家已发货 。</font>
*/
type TaobaoLogisticsOnlineSendAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoLogisticsOnlineSendResponse `json:"logistics_online_send_response,omitempty"` 
    TaobaoLogisticsOnlineSendResponse
}

/* model for simplify = false
type TaobaoLogisticsOnlineSendResponse struct {

    // de
    
    Shipping  *struct {
        Shipping  *Shipping `json:"shipping,omitempty"`
    } `json:"shipping,omitempty"`
    

}
*/

type TaobaoLogisticsOnlineSendResponse struct {

    // de
    Shipping   *Shipping `json:"shipping,omitempty"`

}
