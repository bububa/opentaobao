package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取用户openId APIResponse
taobao.koubei.tribe.open.user.query

口碑综合体通过手机号码获取加密后的用户openId
*/
type TaobaoKoubeiTribeOpenUserQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoKoubeiTribeOpenUserQueryResponse `json:"taobao_koubei_tribe_open_user_query_response,omitempty"`
}

type TaobaoKoubeiTribeOpenUserQueryResponse struct {

    // 接口返回model
    Result   *TaobaoKoubeiTribeOpenUserQueryResult `json:"result,omitempty"`

}
