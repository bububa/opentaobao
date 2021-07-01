package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAisAssetsTagGenerateAPIRequest
基础设施资产标签生成 API请求
alibaba.ais.assets.tag.generate

提供浪潮，英业达等厂商供应阿里巴巴基础设施资产的标签QR code生成 */
type AlibabaAisAssetsTagGenerateAPIRequest struct {
	model.Params
	// 请求资产信息
	_requestParam string
}

// NewAlibabaAisAssetsTagGenerateRequest 初始化AlibabaAisAssetsTagGenerateAPIRequest对象
func NewAlibabaAisAssetsTagGenerateRequest() *AlibabaAisAssetsTagGenerateAPIRequest {
	return &AlibabaAisAssetsTagGenerateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAisAssetsTagGenerateAPIRequest) GetApiMethodName() string {
	return "alibaba.ais.assets.tag.generate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAisAssetsTagGenerateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RequestParam Setter
// 请求资产信息
func (r *AlibabaAisAssetsTagGenerateAPIRequest) SetRequestParam(_requestParam string) error {
	r._requestParam = _requestParam
	r.Set("request_param", _requestParam)
	return nil
}

// Get RequestParam Getter
func (r AlibabaAisAssetsTagGenerateAPIRequest) GetRequestParam() string {
	return r._requestParam
}
