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
    Response *TaobaoTmcUserGetResponse `json:"taobao_tmc_user_get_response,omitempty"`
}

type TaobaoTmcUserGetResponse struct {

    // 开通的用户数据
    TmcUser   *TmcUser `json:"tmc_user,omitempty"`

}
