package xiami

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xiami"
)

/* 
查询是否签约 
alibaba.xiami.api.contract.issign

查询是否签约
*/
func AlibabaXiamiApiContractIssign(clt *core.SDKClient, req *xiami.AlibabaXiamiApiContractIssignRequest, session string) (*xiami.AlibabaXiamiApiContractIssignResponse, error) {
    var resp xiami.AlibabaXiamiApiContractIssignAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
