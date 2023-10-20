package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingitembuygiftqueryitemsAPIRequest 查询买赠活动下的商品 API请求
// alibaba.hm.marketing.itembuygift.queryitems
//
// 查询买赠活动下的商品
type AlibabahmmarketingitembuygiftqueryitemsAPIRequest struct {
	model.Params
	// 查询入参
	_param *ActivitySkuQuery
}

// NewAlibabahmmarketingitembuygiftqueryitemsRequest 初始化AlibabahmmarketingitembuygiftqueryitemsAPIRequest对象
func NewAlibabahmmarketingitembuygiftqueryitemsRequest() *AlibabahmmarketingitembuygiftqueryitemsAPIRequest {
	return &AlibabahmmarketingitembuygiftqueryitemsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingitembuygiftqueryitemsAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itembuygift.queryitems"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingitembuygiftqueryitemsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingitembuygiftqueryitemsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 查询入参
func (r *AlibabahmmarketingitembuygiftqueryitemsAPIRequest) SetParam(_param *ActivitySkuQuery) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabahmmarketingitembuygiftqueryitemsAPIRequest) GetParam() *ActivitySkuQuery {
	return r._param
}
