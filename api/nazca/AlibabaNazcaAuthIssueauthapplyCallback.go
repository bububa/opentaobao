package nazca

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/nazca"
)

/* 
出证申请回调 
alibaba.nazca.auth.issueauthapply.callback

出证申请回调
*/
func AlibabaNazcaAuthIssueauthapplyCallback(clt *core.SDKClient, req *nazca.AlibabaNazcaAuthIssueauthapplyCallbackAPIRequest, session string) (*nazca.AlibabaNazcaAuthIssueauthapplyCallbackAPIResponse, error) {
    var resp nazca.AlibabaNazcaAuthIssueauthapplyCallbackAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
