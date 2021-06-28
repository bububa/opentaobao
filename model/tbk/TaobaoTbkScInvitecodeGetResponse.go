package tbk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-公用-私域用户邀请码生成 APIResponse
taobao.tbk.sc.invitecode.get

私域用户管理(即渠道管理或会员运营管理)功能中，通过此API可生成淘宝客自身的邀请码。
*/
type TaobaoTbkScInvitecodeGetAPIResponse struct {
    model.CommonResponse
    TaobaoTbkScInvitecodeGetResponse
}

type TaobaoTbkScInvitecodeGetResponse struct {
    XMLName xml.Name `xml:"tbk_sc_invitecode_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // data
    
    Data   *TaobaoTbkScInvitecodeGetData `json:"data,omitempty" xml:"data,omitempty"`

    
}
