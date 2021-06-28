package game

import (
    "github.com/bububa/opentaobao/model"
)

/* 
老用户激活并兑换通知接口 APIResponse
taobao.apple.olduser.charge.notify

老用户激活并兑换通知接口
*/
type TaobaoAppleOlduserChargeNotifyAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAppleOlduserChargeNotifyResponse `json:"apple_olduser_charge_notify_response,omitempty"` 
    TaobaoAppleOlduserChargeNotifyResponse
}

/* model for simplify = false
type TaobaoAppleOlduserChargeNotifyResponse struct {

    // 处理结果说明
    
    ResultMsg   string `json:"result_msg,omitempty"`
    

    // 处理结果码
    
    ResultCode   string `json:"result_code,omitempty"`
    

}
*/

type TaobaoAppleOlduserChargeNotifyResponse struct {

    // 处理结果说明
    ResultMsg   string `json:"result_msg,omitempty"`

    // 处理结果码
    ResultCode   string `json:"result_code,omitempty"`

}
