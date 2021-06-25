package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
基础设施资产标签废弃 APIRequest
alibaba.ais.assets.tag.abort

提供浪潮，英业达等厂商供应阿里巴巴基础设施资产的标签QR code未使用的废弃
*/
type AlibabaAisAssetsTagAbortRequest struct {
    model.Params

    // 请求资产信息
    requestParam   string 

}

func NewAlibabaAisAssetsTagAbortRequest() *AlibabaAisAssetsTagAbortRequest{
    return &AlibabaAisAssetsTagAbortRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAisAssetsTagAbortRequest) GetApiMethodName() string {
    return "alibaba.ais.assets.tag.abort"
}

func (r AlibabaAisAssetsTagAbortRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAisAssetsTagAbortRequest) SetRequestParam(requestParam string) error {
    r.requestParam = requestParam
    r.Set("request_param", requestParam)
    return nil
}

func (r AlibabaAisAssetsTagAbortRequest) GetRequestParam() string {
    return r.requestParam
}

