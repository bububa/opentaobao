package user

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/user"
)

/* 
Open Account导入数据 
taobao.open.account.create

Open Account导入数据
*/
func TaobaoOpenAccountCreate(clt *core.SDKClient, req *user.TaobaoOpenAccountCreateAPIRequest, session string) (*user.TaobaoOpenAccountCreateAPIResponse, error) {
    var resp user.TaobaoOpenAccountCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
