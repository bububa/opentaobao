package mtopopen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mtopopen"
)

/* 
淘宝OauthCode颁发 
taobao.oauth.code.create

手淘无线开放的oauthCode颁发接口
*/
func TaobaoOauthCodeCreate(clt *core.SDKClient, req *mtopopen.TaobaoOauthCodeCreateRequest, session string) (*mtopopen.TaobaoOauthCodeCreateResponse, error) {
    var resp mtopopen.TaobaoOauthCodeCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
