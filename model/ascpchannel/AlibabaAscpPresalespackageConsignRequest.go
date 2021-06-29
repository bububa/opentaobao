package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
预售预包尾款推单发货 APIRequest
alibaba.ascp.presalespackage.consign

预售预包尾款发货后推单处理
*/
type AlibabaAscpPresalespackageConsignRequest struct {
    model.Params

    // 入参
    requestParams   *Requestparams 

}

func NewAlibabaAscpPresalespackageConsignRequest() *AlibabaAscpPresalespackageConsignRequest{
    return &AlibabaAscpPresalespackageConsignRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpPresalespackageConsignRequest) GetApiMethodName() string {
    return "alibaba.ascp.presalespackage.consign"
}

func (r AlibabaAscpPresalespackageConsignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpPresalespackageConsignRequest) SetRequestParams(requestParams *Requestparams) error {
    r.requestParams = requestParams
    r.Set("request_params", requestParams)
    return nil
}

func (r AlibabaAscpPresalespackageConsignRequest) GetRequestParams() *Requestparams {
    return r.requestParams
}

