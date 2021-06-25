package tmc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
取消用户的消息服务 APIResponse
taobao.tmc.user.cancel

取消用户的消息服务
*/
type TaobaoTmcUserCancelAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTmcUserCancelResponse `json:"taobao_tmc_user_cancel_response,omitempty"`
}

type TaobaoTmcUserCancelResponse struct {

    // 是否成功,如果为false并且没有错误码，表示删除的用户不存在。
    IsSuccess   bool `json:"is_success,omitempty"`

}
