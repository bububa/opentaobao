package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaisassetstaggetAPIRequest 基础设施资产标签获取 API请求
// alibaba.ais.assets.tag.get
//
// 提供浪潮，英业达等厂商供应阿里巴巴基础设施资产的标签QR code获取
type AlibabaaisassetstaggetAPIRequest struct {
	model.Params
	// 二维码生成唯一标识
	_uNonce string
}

// NewAlibabaaisassetstaggetRequest 初始化AlibabaaisassetstaggetAPIRequest对象
func NewAlibabaaisassetstaggetRequest() *AlibabaaisassetstaggetAPIRequest {
	return &AlibabaaisassetstaggetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaisassetstaggetAPIRequest) GetApiMethodName() string {
	return "alibaba.ais.assets.tag.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaisassetstaggetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaisassetstaggetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUNonce is UNonce Setter
// 二维码生成唯一标识
func (r *AlibabaaisassetstaggetAPIRequest) SetUNonce(_uNonce string) error {
	r._uNonce = _uNonce
	r.Set("u_nonce", _uNonce)
	return nil
}

// GetUNonce UNonce Getter
func (r AlibabaaisassetstaggetAPIRequest) GetUNonce() string {
	return r._uNonce
}
