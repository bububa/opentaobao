package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
预售预包尾款推单发货 API请求
alibaba.ascp.presalespackage.consign

预售预包尾款发货后推单处理
*/
type AlibabaAscpPresalespackageConsignRequest struct {
    model.Params
    // 入参
    requestParams   *Requestparams
}

// 初始化AlibabaAscpPresalespackageConsignRequest对象
func NewAlibabaAscpPresalespackageConsignRequest() *AlibabaAscpPresalespackageConsignRequest{
    return &AlibabaAscpPresalespackageConsignRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpPresalespackageConsignRequest) GetApiMethodName() string {
    return "alibaba.ascp.presalespackage.consign"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpPresalespackageConsignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RequestParams Setter
// 入参
func (r *AlibabaAscpPresalespackageConsignRequest) SetRequestParams(requestParams *Requestparams) error {
    r.requestParams = requestParams
    r.Set("request_params", requestParams)
    return nil
}

// RequestParams Getter
func (r AlibabaAscpPresalespackageConsignRequest) GetRequestParams() *Requestparams {
    return r.requestParams
}
