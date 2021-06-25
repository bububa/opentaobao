package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
基础设施资产标签生成 APIRequest
alibaba.ais.assets.tag.generate

提供浪潮，英业达等厂商供应阿里巴巴基础设施资产的标签QR code生成
*/
type AlibabaAisAssetsTagGenerateRequest struct {
    model.Params

    // 请求资产信息
    requestParam   string 

}

func NewAlibabaAisAssetsTagGenerateRequest() *AlibabaAisAssetsTagGenerateRequest{
    return &AlibabaAisAssetsTagGenerateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAisAssetsTagGenerateRequest) GetApiMethodName() string {
    return "alibaba.ais.assets.tag.generate"
}

func (r AlibabaAisAssetsTagGenerateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAisAssetsTagGenerateRequest) SetRequestParam(requestParam string) error {
    r.requestParam = requestParam
    r.Set("request_param", requestParam)
    return nil
}

func (r AlibabaAisAssetsTagGenerateRequest) GetRequestParam() string {
    return r.requestParam
}

