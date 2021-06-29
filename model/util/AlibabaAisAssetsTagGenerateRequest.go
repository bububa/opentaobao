package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
基础设施资产标签生成 API请求
alibaba.ais.assets.tag.generate

提供浪潮，英业达等厂商供应阿里巴巴基础设施资产的标签QR code生成
*/
type AlibabaAisAssetsTagGenerateRequest struct {
    model.Params
    // 请求资产信息
    _requestParam   string
}

// 初始化AlibabaAisAssetsTagGenerateRequest对象
func NewAlibabaAisAssetsTagGenerateRequest() *AlibabaAisAssetsTagGenerateRequest{
    return &AlibabaAisAssetsTagGenerateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAisAssetsTagGenerateRequest) GetApiMethodName() string {
    return "alibaba.ais.assets.tag.generate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAisAssetsTagGenerateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RequestParam Setter
// 请求资产信息
func (r *AlibabaAisAssetsTagGenerateRequest) SetRequestParam(_requestParam string) error {
    r._requestParam = _requestParam
    r.Set("request_param", _requestParam)
    return nil
}

// RequestParam Getter
func (r AlibabaAisAssetsTagGenerateRequest) GetRequestParam() string {
    return r._requestParam
}
