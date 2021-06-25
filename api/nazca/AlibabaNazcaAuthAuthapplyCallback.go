package nazca

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/nazca"
)

/* 
认证的统一回调接口 
alibaba.nazca.auth.authapply.callback

认证的统一回调接口
*/
func AlibabaNazcaAuthAuthapplyCallback(clt *core.SDKClient, req *nazca.AlibabaNazcaAuthAuthapplyCallbackRequest, session string) (*nazca.AlibabaNazcaAuthAuthapplyCallbackResponse, error) {
    var resp nazca.AlibabaNazcaAuthAuthapplyCallbackAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
