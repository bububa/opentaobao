package caipiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
检查用户是否签署支付宝代购协议 APIResponse
taobao.caipiao.signstatus.check

检查用户是否签署了支付宝代扣协议。如果签署了，返回true; 如果没签署，返回false, 同时返回签署代扣协议的Url。
*/
type TaobaoCaipiaoSignstatusCheckAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoCaipiaoSignstatusCheckResponse `json:"caipiao_signstatus_check_response,omitempty"` 
    TaobaoCaipiaoSignstatusCheckResponse
}

/* model for simplify = false
type TaobaoCaipiaoSignstatusCheckResponse struct {

    // 是否签署了支付宝代扣协议
    
    Sign   bool `json:"sign,omitempty"`
    

    // 签署支付宝代扣协议的Url
    
    SignUrl   string `json:"sign_url,omitempty"`
    

}
*/

type TaobaoCaipiaoSignstatusCheckResponse struct {

    // 是否签署了支付宝代扣协议
    Sign   bool `json:"sign,omitempty"`

    // 签署支付宝代扣协议的Url
    SignUrl   string `json:"sign_url,omitempty"`

}
