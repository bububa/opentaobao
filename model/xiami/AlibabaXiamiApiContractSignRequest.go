package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提供签约链接 API请求
alibaba.xiami.api.contract.sign

提供签约链接。in：商家id；out：签约url
*/
type AlibabaXiamiApiContractSignRequest struct {
    model.Params
}

// 初始化AlibabaXiamiApiContractSignRequest对象
func NewAlibabaXiamiApiContractSignRequest() *AlibabaXiamiApiContractSignRequest{
    return &AlibabaXiamiApiContractSignRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiContractSignRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.contract.sign"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiContractSignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
