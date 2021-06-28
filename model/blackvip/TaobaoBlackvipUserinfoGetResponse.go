package blackvip

import (
    "github.com/bububa/opentaobao/model"
)

/* 
88VIP用户信息查询 APIResponse
taobao.blackvip.userinfo.get

查询88VIP用户信息，比如用户是否是88VIP，88VIP的失效时间等
*/
type TaobaoBlackvipUserinfoGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBlackvipUserinfoGetResponse `json:"blackvip_userinfo_get_response,omitempty"` 
    TaobaoBlackvipUserinfoGetResponse
}

/* model for simplify = false
type TaobaoBlackvipUserinfoGetResponse struct {

    // 结果支持对象
    
    Result  *struct {
        ResultSupport  *ResultSupport `json:"result_support,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoBlackvipUserinfoGetResponse struct {

    // 结果支持对象
    Result   *ResultSupport `json:"result,omitempty"`

}
