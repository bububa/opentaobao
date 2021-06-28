package xhotel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
线下信用住取消结账专用接口 APIResponse
taobao.xhotel.order.offline.settle.cancel

线下信用住取消结账专用接口
*/
type TaobaoXhotelOrderOfflineSettleCancelAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoXhotelOrderOfflineSettleCancelResponse `json:"xhotel_order_offline_settle_cancel_response,omitempty"` 
    TaobaoXhotelOrderOfflineSettleCancelResponse
}

/* model for simplify = false
type TaobaoXhotelOrderOfflineSettleCancelResponse struct {

    // 返回信息
    
    Result   string `json:"result,omitempty"`
    

}
*/

type TaobaoXhotelOrderOfflineSettleCancelResponse struct {

    // 返回信息
    Result   string `json:"result,omitempty"`

}
