package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingitemdiscountqueryactivityAPIRequest 查找特价活动 API请求
// alibaba.hm.marketing.itemdiscount.queryactivity
//
// 查找特价活动
type AlibabahmmarketingitemdiscountqueryactivityAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *CommonActivityParam
}

// NewAlibabahmmarketingitemdiscountqueryactivityRequest 初始化AlibabahmmarketingitemdiscountqueryactivityAPIRequest对象
func NewAlibabahmmarketingitemdiscountqueryactivityRequest() *AlibabahmmarketingitemdiscountqueryactivityAPIRequest {
	return &AlibabahmmarketingitemdiscountqueryactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingitemdiscountqueryactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itemdiscount.queryactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingitemdiscountqueryactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingitemdiscountqueryactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 商品对象
func (r *AlibabahmmarketingitemdiscountqueryactivityAPIRequest) SetParam0(_param0 *CommonActivityParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabahmmarketingitemdiscountqueryactivityAPIRequest) GetParam0() *CommonActivityParam {
	return r._param0
}
