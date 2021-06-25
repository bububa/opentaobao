package nazca

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/nazca"
)

/* 
根据token获取变更认证申请信息 
alibaba.nazca.token.changeauthapply.get

根据token获取变更认证申请信息
*/
func AlibabaNazcaTokenChangeauthapplyGet(clt *core.SDKClient, req *nazca.AlibabaNazcaTokenChangeauthapplyGetRequest, session string) (*nazca.AlibabaNazcaTokenChangeauthapplyGetResponse, error) {
    var resp nazca.AlibabaNazcaTokenChangeauthapplyGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
