package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaisassetstaggenerateAPIRequest 基础设施资产标签生成 API请求
// alibaba.ais.assets.tag.generate
//
// 提供浪潮，英业达等厂商供应阿里巴巴基础设施资产的标签QR code生成
type AlibabaaisassetstaggenerateAPIRequest struct {
	model.Params
	// 请求资产信息
	_requestParam string
}

// NewAlibabaaisassetstaggenerateRequest 初始化AlibabaaisassetstaggenerateAPIRequest对象
func NewAlibabaaisassetstaggenerateRequest() *AlibabaaisassetstaggenerateAPIRequest {
	return &AlibabaaisassetstaggenerateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaisassetstaggenerateAPIRequest) GetApiMethodName() string {
	return "alibaba.ais.assets.tag.generate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaisassetstaggenerateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaisassetstaggenerateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestParam is RequestParam Setter
// 请求资产信息
func (r *AlibabaaisassetstaggenerateAPIRequest) SetRequestParam(_requestParam string) error {
	r._requestParam = _requestParam
	r.Set("request_param", _requestParam)
	return nil
}

// GetRequestParam RequestParam Getter
func (r AlibabaaisassetstaggenerateAPIRequest) GetRequestParam() string {
	return r._requestParam
}
