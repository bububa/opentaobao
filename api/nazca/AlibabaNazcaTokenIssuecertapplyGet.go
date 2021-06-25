package nazca

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/nazca"
)

/* 
根据token获取出证申请信息 
alibaba.nazca.token.issuecertapply.get

根据token获取出证申请信息
*/
func AlibabaNazcaTokenIssuecertapplyGet(clt *core.SDKClient, req *nazca.AlibabaNazcaTokenIssuecertapplyGetRequest, session string) (*nazca.AlibabaNazcaTokenIssuecertapplyGetResponse, error) {
    var resp nazca.AlibabaNazcaTokenIssuecertapplyGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
