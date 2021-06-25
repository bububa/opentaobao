package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询是否签约 APIRequest
alibaba.xiami.api.contract.issign

查询是否签约
*/
type AlibabaXiamiApiContractIssignRequest struct {
    model.Params

}

func NewAlibabaXiamiApiContractIssignRequest() *AlibabaXiamiApiContractIssignRequest{
    return &AlibabaXiamiApiContractIssignRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaXiamiApiContractIssignRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.contract.issign"
}

func (r AlibabaXiamiApiContractIssignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


