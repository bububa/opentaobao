package tbk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-公用-私域用户邀请码生成 APIResponse
taobao.tbk.sc.invitecode.get

私域用户管理(即渠道管理或会员运营管理)功能中，通过此API可生成淘宝客自身的邀请码。
*/
type TaobaoTbkScInvitecodeGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTbkScInvitecodeGetResponse `json:"tbk_sc_invitecode_get_response,omitempty"` 
    TaobaoTbkScInvitecodeGetResponse
}

/* model for simplify = false
type TaobaoTbkScInvitecodeGetResponse struct {

    // data
    
    Data  *struct {
        TaobaoTbkScInvitecodeGetData  *TaobaoTbkScInvitecodeGetData `json:"taobao_tbk_sc_invitecode_get_data,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

type TaobaoTbkScInvitecodeGetResponse struct {

    // data
    Data   *TaobaoTbkScInvitecodeGetData `json:"data,omitempty"`

}
