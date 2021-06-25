package user

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/user"
)

/* 
查询appstore应用订购关系 
taobao.appstore.subscribe.get

查询appstore应用订购关系(对于新上架的多版本应用，建议使用taobao.vas.subscribe.get)
*/
func TaobaoAppstoreSubscribeGet(clt *core.SDKClient, req *user.TaobaoAppstoreSubscribeGetRequest, session string) (*user.TaobaoAppstoreSubscribeGetResponse, error) {
    var resp user.TaobaoAppstoreSubscribeGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
