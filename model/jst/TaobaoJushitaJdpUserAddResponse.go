package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
添加数据推送用户 APIResponse
taobao.jushita.jdp.user.add

提供给接入数据推送的应用添加数据推送服务的用户
*/
type TaobaoJushitaJdpUserAddAPIResponse struct {
    model.CommonResponse
    Response *TaobaoJushitaJdpUserAddResponse `json:"taobao_jushita_jdp_user_add_response,omitempty"`
}

type TaobaoJushitaJdpUserAddResponse struct {

    // 是否添加成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
