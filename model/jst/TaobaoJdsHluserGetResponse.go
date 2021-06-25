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
    Response *TaobaoJdsHluserGetResponse `json:"taobao_jds_hluser_get_response,omitempty"`
}

type TaobaoJdsHluserGetResponse struct {

    // 回流用户信息
    HlUser   *HlUserDO `json:"hl_user,omitempty"`

}
