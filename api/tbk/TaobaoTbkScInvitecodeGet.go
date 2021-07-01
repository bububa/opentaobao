package tbk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tbk"
)

/* 
淘宝客-公用-私域用户邀请码生成 
taobao.tbk.sc.invitecode.get

私域用户管理(即渠道管理或会员运营管理)功能中，通过此API可生成淘宝客自身的邀请码。
*/
func TaobaoTbkScInvitecodeGet(clt *core.SDKClient, req *tbk.TaobaoTbkScInvitecodeGetAPIRequest, session string) (*tbk.TaobaoTbkScInvitecodeGetAPIResponse, error) {
    var resp tbk.TaobaoTbkScInvitecodeGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
