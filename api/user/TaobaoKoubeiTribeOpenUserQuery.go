package user

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/user"
)

/* 
获取用户openId 
taobao.koubei.tribe.open.user.query

口碑综合体通过手机号码获取加密后的用户openId
*/
func TaobaoKoubeiTribeOpenUserQuery(clt *core.SDKClient, req *user.TaobaoKoubeiTribeOpenUserQueryRequest, session string) (*user.TaobaoKoubeiTribeOpenUserQueryAPIResponse, error) {
    var resp user.TaobaoKoubeiTribeOpenUserQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
