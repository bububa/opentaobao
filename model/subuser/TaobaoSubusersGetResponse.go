package subuser

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取指定账户的子账号简易信息列表 APIResponse
taobao.subusers.get

获取主账号下的子账号简易账号信息集合。（只能通过主账号登陆并且查询该属于主账号的子账号信息）
*/
type TaobaoSubusersGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSubusersGetResponse `json:"taobao_subusers_get_response,omitempty"`
}

type TaobaoSubusersGetResponse struct {

    // 子账号基本信息
    Subaccounts   []SubAccountInfo `json:"subaccounts,omitempty"`

}
