package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询买家信息API APIResponse
taobao.user.buyer.get

查询买家信息API，只能买家类应用调用。
*/
type TaobaoUserBuyerGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoUserBuyerGetResponse `json:"taobao_user_buyer_get_response,omitempty"`
}

type TaobaoUserBuyerGetResponse struct {

    // 只返回nick,avatar参数
    User   *User `json:"user,omitempty"`

}
