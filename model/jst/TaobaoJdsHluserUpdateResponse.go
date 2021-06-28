package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
订单全链路用户信息修改 APIResponse
taobao.jds.hluser.update

订单全链路用户信息修改，比如是否开放买家端展示
*/
type TaobaoJdsHluserUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoJdsHluserUpdateResponse `json:"jds_hluser_update_response,omitempty"` 
    TaobaoJdsHluserUpdateResponse
}

/* model for simplify = false
type TaobaoJdsHluserUpdateResponse struct {

    // 是否成功
    
    Result   bool `json:"result,omitempty"`
    

}
*/

type TaobaoJdsHluserUpdateResponse struct {

    // 是否成功
    Result   bool `json:"result,omitempty"`

}
