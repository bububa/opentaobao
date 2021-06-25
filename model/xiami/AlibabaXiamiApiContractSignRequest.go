package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提供签约链接 APIRequest
alibaba.xiami.api.contract.sign

提供签约链接。in：商家id；out：签约url
*/
type AlibabaXiamiApiContractSignRequest struct {
    model.Params

}

func NewAlibabaXiamiApiContractSignRequest() *AlibabaXiamiApiContractSignRequest{
    return &AlibabaXiamiApiContractSignRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaXiamiApiContractSignRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.contract.sign"
}

func (r AlibabaXiamiApiContractSignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


