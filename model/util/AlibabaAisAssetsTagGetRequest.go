package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
基础设施资产标签获取 APIRequest
alibaba.ais.assets.tag.get

提供浪潮，英业达等厂商供应阿里巴巴基础设施资产的标签QR code获取
*/
type AlibabaAisAssetsTagGetRequest struct {
    model.Params

    // 二维码生成唯一标识
    uNonce   string 

}

func NewAlibabaAisAssetsTagGetRequest() *AlibabaAisAssetsTagGetRequest{
    return &AlibabaAisAssetsTagGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAisAssetsTagGetRequest) GetApiMethodName() string {
    return "alibaba.ais.assets.tag.get"
}

func (r AlibabaAisAssetsTagGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAisAssetsTagGetRequest) SetUNonce(uNonce string) error {
    r.uNonce = uNonce
    r.Set("u_nonce", uNonce)
    return nil
}

func (r AlibabaAisAssetsTagGetRequest) GetUNonce() string {
    return r.uNonce
}

