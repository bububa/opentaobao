package util

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAisAssetsTagGenerateAPIRequest 基础设施资产标签生成 API请求
// alibaba.ais.assets.tag.generate
//
// 提供浪潮，英业达等厂商供应阿里巴巴基础设施资产的标签QR code生成
type AlibabaAisAssetsTagGenerateAPIRequest struct {
	model.Params
	// 请求资产信息
	_requestParam string
}

// NewAlibabaAisAssetsTagGenerateRequest 初始化AlibabaAisAssetsTagGenerateAPIRequest对象
func NewAlibabaAisAssetsTagGenerateRequest() *AlibabaAisAssetsTagGenerateAPIRequest {
	return &AlibabaAisAssetsTagGenerateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAisAssetsTagGenerateAPIRequest) Reset() {
	r._requestParam = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAisAssetsTagGenerateAPIRequest) GetApiMethodName() string {
	return "alibaba.ais.assets.tag.generate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAisAssetsTagGenerateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAisAssetsTagGenerateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestParam is RequestParam Setter
// 请求资产信息
func (r *AlibabaAisAssetsTagGenerateAPIRequest) SetRequestParam(_requestParam string) error {
	r._requestParam = _requestParam
	r.Set("request_param", _requestParam)
	return nil
}

// GetRequestParam RequestParam Getter
func (r AlibabaAisAssetsTagGenerateAPIRequest) GetRequestParam() string {
	return r._requestParam
}

var poolAlibabaAisAssetsTagGenerateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAisAssetsTagGenerateRequest()
	},
}

// GetAlibabaAisAssetsTagGenerateRequest 从 sync.Pool 获取 AlibabaAisAssetsTagGenerateAPIRequest
func GetAlibabaAisAssetsTagGenerateAPIRequest() *AlibabaAisAssetsTagGenerateAPIRequest {
	return poolAlibabaAisAssetsTagGenerateAPIRequest.Get().(*AlibabaAisAssetsTagGenerateAPIRequest)
}

// ReleaseAlibabaAisAssetsTagGenerateAPIRequest 将 AlibabaAisAssetsTagGenerateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAisAssetsTagGenerateAPIRequest(v *AlibabaAisAssetsTagGenerateAPIRequest) {
	v.Reset()
	poolAlibabaAisAssetsTagGenerateAPIRequest.Put(v)
}
