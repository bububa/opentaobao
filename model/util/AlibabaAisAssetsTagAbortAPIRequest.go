package util

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAisAssetsTagAbortAPIRequest 基础设施资产标签废弃 API请求
// alibaba.ais.assets.tag.abort
//
// 提供浪潮，英业达等厂商供应阿里巴巴基础设施资产的标签QR code未使用的废弃
type AlibabaAisAssetsTagAbortAPIRequest struct {
	model.Params
	// 请求资产信息
	_requestParam string
}

// NewAlibabaAisAssetsTagAbortRequest 初始化AlibabaAisAssetsTagAbortAPIRequest对象
func NewAlibabaAisAssetsTagAbortRequest() *AlibabaAisAssetsTagAbortAPIRequest {
	return &AlibabaAisAssetsTagAbortAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAisAssetsTagAbortAPIRequest) Reset() {
	r._requestParam = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAisAssetsTagAbortAPIRequest) GetApiMethodName() string {
	return "alibaba.ais.assets.tag.abort"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAisAssetsTagAbortAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAisAssetsTagAbortAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestParam is RequestParam Setter
// 请求资产信息
func (r *AlibabaAisAssetsTagAbortAPIRequest) SetRequestParam(_requestParam string) error {
	r._requestParam = _requestParam
	r.Set("request_param", _requestParam)
	return nil
}

// GetRequestParam RequestParam Getter
func (r AlibabaAisAssetsTagAbortAPIRequest) GetRequestParam() string {
	return r._requestParam
}

var poolAlibabaAisAssetsTagAbortAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAisAssetsTagAbortRequest()
	},
}

// GetAlibabaAisAssetsTagAbortRequest 从 sync.Pool 获取 AlibabaAisAssetsTagAbortAPIRequest
func GetAlibabaAisAssetsTagAbortAPIRequest() *AlibabaAisAssetsTagAbortAPIRequest {
	return poolAlibabaAisAssetsTagAbortAPIRequest.Get().(*AlibabaAisAssetsTagAbortAPIRequest)
}

// ReleaseAlibabaAisAssetsTagAbortAPIRequest 将 AlibabaAisAssetsTagAbortAPIRequest 放入 sync.Pool
func ReleaseAlibabaAisAssetsTagAbortAPIRequest(v *AlibabaAisAssetsTagAbortAPIRequest) {
	v.Reset()
	poolAlibabaAisAssetsTagAbortAPIRequest.Put(v)
}
