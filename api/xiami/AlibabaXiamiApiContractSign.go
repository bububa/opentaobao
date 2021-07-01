package xiami

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xiami"
)

/* 
提供签约链接 
alibaba.xiami.api.contract.sign

提供签约链接。in：商家id；out：签约url
*/
func AlibabaXiamiApiContractSign(clt *core.SDKClient, req *xiami.AlibabaXiamiApiContractSignAPIRequest, session string) (*xiami.AlibabaXiamiApiContractSignAPIResponse, error) {
    var resp xiami.AlibabaXiamiApiContractSignAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
