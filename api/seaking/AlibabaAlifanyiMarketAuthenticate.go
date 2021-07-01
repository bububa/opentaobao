package seaking

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/seaking"
)

/* 
第三方授权 
alibaba.alifanyi.market.authenticate

第三方授权，获取授权码
*/
func AlibabaAlifanyiMarketAuthenticate(clt *core.SDKClient, req *seaking.AlibabaAlifanyiMarketAuthenticateAPIRequest, session string) (*seaking.AlibabaAlifanyiMarketAuthenticateAPIResponse, error) {
    var resp seaking.AlibabaAlifanyiMarketAuthenticateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
