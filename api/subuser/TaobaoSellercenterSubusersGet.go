package subuser

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/subuser"
)

/* 
查询指定账户的子账号列表 
taobao.sellercenter.subusers.get

根据主账号nick查询该账号下所有的子账号列表，只能查询属于自己的账号信息 (主账号以及所属子账号)
*/
func TaobaoSellercenterSubusersGet(clt *core.SDKClient, req *subuser.TaobaoSellercenterSubusersGetRequest, session string) (*subuser.TaobaoSellercenterSubusersGetAPIResponse, error) {
    var resp subuser.TaobaoSellercenterSubusersGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
