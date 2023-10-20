package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingitemdiscountqueryactivityAPIRequest 查找特价活动 API请求
// alibaba.wdk.marketing.itemdiscount.queryactivity
//
// 查找特价活动
type AlibabawdkmarketingitemdiscountqueryactivityAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *CommonActivityParam
}

// NewAlibabawdkmarketingitemdiscountqueryactivityRequest 初始化AlibabawdkmarketingitemdiscountqueryactivityAPIRequest对象
func NewAlibabawdkmarketingitemdiscountqueryactivityRequest() *AlibabawdkmarketingitemdiscountqueryactivityAPIRequest {
	return &AlibabawdkmarketingitemdiscountqueryactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingitemdiscountqueryactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itemdiscount.queryactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingitemdiscountqueryactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingitemdiscountqueryactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 商品对象
func (r *AlibabawdkmarketingitemdiscountqueryactivityAPIRequest) SetParam0(_param0 *CommonActivityParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabawdkmarketingitemdiscountqueryactivityAPIRequest) GetParam0() *CommonActivityParam {
	return r._param0
}
