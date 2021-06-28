package nazca

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/nazca"
)

/* 
变更认证回调 
alibaba.nazca.auth.changeauthapply.callback

变更认证回调
*/
func AlibabaNazcaAuthChangeauthapplyCallback(clt *core.SDKClient, req *nazca.AlibabaNazcaAuthChangeauthapplyCallbackRequest, session string) (*nazca.AlibabaNazcaAuthChangeauthapplyCallbackAPIResponse, error) {
    var resp nazca.AlibabaNazcaAuthChangeauthapplyCallbackAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
