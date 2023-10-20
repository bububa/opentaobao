package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingfullrangequeryitemAPIRequest 全场活动查询换购品 API请求
// alibaba.wdk.marketing.fullrange.queryitem
//
// 全场活动查询换购品
type AlibabawdkmarketingfullrangequeryitemAPIRequest struct {
	model.Params
	// 换购商品查询参数
	_param0 *ActivitySkuQuery
}

// NewAlibabawdkmarketingfullrangequeryitemRequest 初始化AlibabawdkmarketingfullrangequeryitemAPIRequest对象
func NewAlibabawdkmarketingfullrangequeryitemRequest() *AlibabawdkmarketingfullrangequeryitemAPIRequest {
	return &AlibabawdkmarketingfullrangequeryitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingfullrangequeryitemAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.fullrange.queryitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingfullrangequeryitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingfullrangequeryitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 换购商品查询参数
func (r *AlibabawdkmarketingfullrangequeryitemAPIRequest) SetParam0(_param0 *ActivitySkuQuery) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabawdkmarketingfullrangequeryitemAPIRequest) GetParam0() *ActivitySkuQuery {
	return r._param0
}
