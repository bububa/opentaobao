package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
口碑综合体手机号获取验证码 APIResponse
taobao.koubei.tribe.open.verify.code.apply

口碑综合体通过手机号获取验证码对外开放接口
*/
type TaobaoKoubeiTribeOpenVerifyCodeApplyAPIResponse struct {
    model.CommonResponse
    TaobaoKoubeiTribeOpenVerifyCodeApplyResponse
}

type TaobaoKoubeiTribeOpenVerifyCodeApplyResponse struct {
    XMLName xml.Name `xml:"koubei_tribe_open_verify_code_apply_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *TaobaoKoubeiTribeOpenVerifyCodeApplyResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
