package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
基础设施资产标签获取 API请求
alibaba.ais.assets.tag.get

提供浪潮，英业达等厂商供应阿里巴巴基础设施资产的标签QR code获取
*/
type AlibabaAisAssetsTagGetAPIRequest struct {
    model.Params
    // 二维码生成唯一标识
    _uNonce   string
}

// 初始化AlibabaAisAssetsTagGetAPIRequest对象
func NewAlibabaAisAssetsTagGetRequest() *AlibabaAisAssetsTagGetAPIRequest{
    return &AlibabaAisAssetsTagGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAisAssetsTagGetAPIRequest) GetApiMethodName() string {
    return "alibaba.ais.assets.tag.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAisAssetsTagGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UNonce Setter
// 二维码生成唯一标识
func (r *AlibabaAisAssetsTagGetAPIRequest) SetUNonce(_uNonce string) error {
    r._uNonce = _uNonce
    r.Set("u_nonce", _uNonce)
    return nil
}

// UNonce Getter
func (r AlibabaAisAssetsTagGetAPIRequest) GetUNonce() string {
    return r._uNonce
}
