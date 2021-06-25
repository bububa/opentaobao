package util

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/util"
)

/* 
获取授权session信息 
taobao.openlink.session.get

帮助第三方isv生成三方session
*/
func TaobaoOpenlinkSessionGet(clt *core.SDKClient, req *util.TaobaoOpenlinkSessionGetRequest, session string) (*util.TaobaoOpenlinkSessionGetResponse, error) {
    var resp util.TaobaoOpenlinkSessionGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
