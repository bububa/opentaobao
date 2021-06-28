package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
订单全链路用户信息获取 APIResponse
taobao.jds.hluser.get

订单全链路用户信息获取
*/
type TaobaoJdsHluserGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoJdsHluserGetResponse `json:"jds_hluser_get_response,omitempty"` 
    TaobaoJdsHluserGetResponse
}

/* model for simplify = false
type TaobaoJdsHluserGetResponse struct {

    // 回流用户信息
    
    HlUser  *struct {
        HlUserDO  *HlUserDO `json:"hl_user_do,omitempty"`
    } `json:"hl_user,omitempty"`
    

}
*/

type TaobaoJdsHluserGetResponse struct {

    // 回流用户信息
    HlUser   *HlUserDO `json:"hl_user,omitempty"`

}
