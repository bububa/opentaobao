package subuser

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/subuser"
)

/* 
获取指定账户的所有职务信息列表 
taobao.subuser.dutys.get

通过主账号Nick获取该账户下的所有职务信息，职务信息中包括职务ID、职务名称以及职务等级（通过主账号登陆只能获取属于该主账号下的职务信息）
*/
func TaobaoSubuserDutysGet(clt *core.SDKClient, req *subuser.TaobaoSubuserDutysGetRequest, session string) (*subuser.TaobaoSubuserDutysGetResponse, error) {
    var resp subuser.TaobaoSubuserDutysGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
