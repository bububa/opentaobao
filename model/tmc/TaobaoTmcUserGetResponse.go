package tmc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取用户已开通消息 APIResponse
taobao.tmc.user.get

查询指定用户开通的消息通道和组
*/
type TaobaoTmcUserGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTmcUserGetResponse `json:"tmc_user_get_response,omitempty"` 
    TaobaoTmcUserGetResponse
}

/* model for simplify = false
type TaobaoTmcUserGetResponse struct {

    // 开通的用户数据
    
    TmcUser  *struct {
        TmcUser  *TmcUser `json:"tmc_user,omitempty"`
    } `json:"tmc_user,omitempty"`
    

}
*/

type TaobaoTmcUserGetResponse struct {

    // 开通的用户数据
    TmcUser   *TmcUser `json:"tmc_user,omitempty"`

}
