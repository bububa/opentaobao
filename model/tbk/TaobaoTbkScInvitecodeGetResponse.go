package tbk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-公用-私域用户邀请码生成 API返回值 
taobao.tbk.sc.invitecode.get

私域用户管理(即渠道管理或会员运营管理)功能中，通过此API可生成淘宝客自身的邀请码。
*/
type TaobaoTbkScInvitecodeGetAPIResponse struct {
    model.CommonResponse
    TaobaoTbkScInvitecodeGetResponse
}

// 淘宝客-公用-私域用户邀请码生成 成功返回结果
type TaobaoTbkScInvitecodeGetResponse struct {
    XMLName xml.Name `xml:"tbk_sc_invitecode_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // data
    Data   *TaobaoTbkScInvitecodeGetData `json:"data,omitempty" xml:"data,omitempty"`
}
