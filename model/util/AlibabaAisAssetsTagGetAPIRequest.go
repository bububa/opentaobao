package util

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAisAssetsTagGetAPIRequest 基础设施资产标签获取 API请求
// alibaba.ais.assets.tag.get
//
// 提供浪潮，英业达等厂商供应阿里巴巴基础设施资产的标签QR code获取
type AlibabaAisAssetsTagGetAPIRequest struct {
	model.Params
	// 二维码生成唯一标识
	_uNonce string
}

// NewAlibabaAisAssetsTagGetRequest 初始化AlibabaAisAssetsTagGetAPIRequest对象
func NewAlibabaAisAssetsTagGetRequest() *AlibabaAisAssetsTagGetAPIRequest {
	return &AlibabaAisAssetsTagGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAisAssetsTagGetAPIRequest) Reset() {
	r._uNonce = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAisAssetsTagGetAPIRequest) GetApiMethodName() string {
	return "alibaba.ais.assets.tag.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAisAssetsTagGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAisAssetsTagGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUNonce is UNonce Setter
// 二维码生成唯一标识
func (r *AlibabaAisAssetsTagGetAPIRequest) SetUNonce(_uNonce string) error {
	r._uNonce = _uNonce
	r.Set("u_nonce", _uNonce)
	return nil
}

// GetUNonce UNonce Getter
func (r AlibabaAisAssetsTagGetAPIRequest) GetUNonce() string {
	return r._uNonce
}

var poolAlibabaAisAssetsTagGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAisAssetsTagGetRequest()
	},
}

// GetAlibabaAisAssetsTagGetRequest 从 sync.Pool 获取 AlibabaAisAssetsTagGetAPIRequest
func GetAlibabaAisAssetsTagGetAPIRequest() *AlibabaAisAssetsTagGetAPIRequest {
	return poolAlibabaAisAssetsTagGetAPIRequest.Get().(*AlibabaAisAssetsTagGetAPIRequest)
}

// ReleaseAlibabaAisAssetsTagGetAPIRequest 将 AlibabaAisAssetsTagGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAisAssetsTagGetAPIRequest(v *AlibabaAisAssetsTagGetAPIRequest) {
	v.Reset()
	poolAlibabaAisAssetsTagGetAPIRequest.Put(v)
}
