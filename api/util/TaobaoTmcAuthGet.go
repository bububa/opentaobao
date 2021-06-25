package util

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/util"
)

/* 
TMC授权token 
taobao.tmc.auth.get

TMC连接授权Token
*/
func TaobaoTmcAuthGet(clt *core.SDKClient, req *util.TaobaoTmcAuthGetRequest, session string) (*util.TaobaoTmcAuthGetResponse, error) {
    var resp util.TaobaoTmcAuthGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
