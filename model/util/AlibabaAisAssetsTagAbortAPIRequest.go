package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
基础设施资产标签废弃 API请求
alibaba.ais.assets.tag.abort

提供浪潮，英业达等厂商供应阿里巴巴基础设施资产的标签QR code未使用的废弃
*/
type AlibabaAisAssetsTagAbortAPIRequest struct {
    model.Params
    // 请求资产信息
    _requestParam   string
}

// 初始化AlibabaAisAssetsTagAbortAPIRequest对象
func NewAlibabaAisAssetsTagAbortRequest() *AlibabaAisAssetsTagAbortAPIRequest{
    return &AlibabaAisAssetsTagAbortAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAisAssetsTagAbortAPIRequest) GetApiMethodName() string {
    return "alibaba.ais.assets.tag.abort"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAisAssetsTagAbortAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RequestParam Setter
// 请求资产信息
func (r *AlibabaAisAssetsTagAbortAPIRequest) SetRequestParam(_requestParam string) error {
    r._requestParam = _requestParam
    r.Set("request_param", _requestParam)
    return nil
}

// RequestParam Getter
func (r AlibabaAisAssetsTagAbortAPIRequest) GetRequestParam() string {
    return r._requestParam
}
