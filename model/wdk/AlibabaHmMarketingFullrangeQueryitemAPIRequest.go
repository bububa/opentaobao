package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingfullrangequeryitemAPIRequest 全场活动查询换购品 API请求
// alibaba.hm.marketing.fullrange.queryitem
//
// 全场活动查询换购品
type AlibabahmmarketingfullrangequeryitemAPIRequest struct {
	model.Params
	// 换购商品查询参数
	_param0 *ActivitySkuQuery
}

// NewAlibabahmmarketingfullrangequeryitemRequest 初始化AlibabahmmarketingfullrangequeryitemAPIRequest对象
func NewAlibabahmmarketingfullrangequeryitemRequest() *AlibabahmmarketingfullrangequeryitemAPIRequest {
	return &AlibabahmmarketingfullrangequeryitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingfullrangequeryitemAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.fullrange.queryitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingfullrangequeryitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingfullrangequeryitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 换购商品查询参数
func (r *AlibabahmmarketingfullrangequeryitemAPIRequest) SetParam0(_param0 *ActivitySkuQuery) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabahmmarketingfullrangequeryitemAPIRequest) GetParam0() *ActivitySkuQuery {
	return r._param0
}
