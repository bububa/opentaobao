package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
口碑综合体手机号获取验证码 APIResponse
taobao.koubei.tribe.open.verify.code.apply

口碑综合体通过手机号获取验证码对外开放接口
*/
type TaobaoKoubeiTribeOpenVerifyCodeApplyAPIResponse struct {
    model.CommonResponse
    Response *TaobaoKoubeiTribeOpenVerifyCodeApplyResponse `json:"taobao_koubei_tribe_open_verify_code_apply_response,omitempty"`
}

type TaobaoKoubeiTribeOpenVerifyCodeApplyResponse struct {

    // 接口返回model
    Result   *TaobaoKoubeiTribeOpenVerifyCodeApplyResult `json:"result,omitempty"`

}
