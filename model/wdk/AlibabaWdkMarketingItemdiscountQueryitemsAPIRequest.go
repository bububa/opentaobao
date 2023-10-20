package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingitemdiscountqueryitemsAPIRequest 查询特价商品 API请求
// alibaba.wdk.marketing.itemdiscount.queryitems
//
// 查询参加特价活动的商品优惠详情
type AlibabawdkmarketingitemdiscountqueryitemsAPIRequest struct {
	model.Params
	// 查询入参
	_param *ActivitySkuQuery
}

// NewAlibabawdkmarketingitemdiscountqueryitemsRequest 初始化AlibabawdkmarketingitemdiscountqueryitemsAPIRequest对象
func NewAlibabawdkmarketingitemdiscountqueryitemsRequest() *AlibabawdkmarketingitemdiscountqueryitemsAPIRequest {
	return &AlibabawdkmarketingitemdiscountqueryitemsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingitemdiscountqueryitemsAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itemdiscount.queryitems"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingitemdiscountqueryitemsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingitemdiscountqueryitemsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 查询入参
func (r *AlibabawdkmarketingitemdiscountqueryitemsAPIRequest) SetParam(_param *ActivitySkuQuery) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabawdkmarketingitemdiscountqueryitemsAPIRequest) GetParam() *ActivitySkuQuery {
	return r._param
}
