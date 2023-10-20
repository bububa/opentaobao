package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingitemdiscountqueryitemsAPIRequest 查询特价商品 API请求
// alibaba.hm.marketing.itemdiscount.queryitems
//
// 查询参加特价活动的商品优惠详情
type AlibabahmmarketingitemdiscountqueryitemsAPIRequest struct {
	model.Params
	// 查询入参
	_param *ActivitySkuQuery
}

// NewAlibabahmmarketingitemdiscountqueryitemsRequest 初始化AlibabahmmarketingitemdiscountqueryitemsAPIRequest对象
func NewAlibabahmmarketingitemdiscountqueryitemsRequest() *AlibabahmmarketingitemdiscountqueryitemsAPIRequest {
	return &AlibabahmmarketingitemdiscountqueryitemsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingitemdiscountqueryitemsAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itemdiscount.queryitems"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingitemdiscountqueryitemsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingitemdiscountqueryitemsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 查询入参
func (r *AlibabahmmarketingitemdiscountqueryitemsAPIRequest) SetParam(_param *ActivitySkuQuery) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabahmmarketingitemdiscountqueryitemsAPIRequest) GetParam() *ActivitySkuQuery {
	return r._param
}
