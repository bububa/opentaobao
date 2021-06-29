package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询是否签约 API请求
alibaba.xiami.api.contract.issign

查询是否签约
*/
type AlibabaXiamiApiContractIssignRequest struct {
    model.Params
}

// 初始化AlibabaXiamiApiContractIssignRequest对象
func NewAlibabaXiamiApiContractIssignRequest() *AlibabaXiamiApiContractIssignRequest{
    return &AlibabaXiamiApiContractIssignRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiContractIssignRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.contract.issign"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiContractIssignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
