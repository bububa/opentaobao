package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingitempoolqueryitemsAPIRequest 查询商品池活动下面的商品 API请求
// alibaba.hm.marketing.itempool.queryitems
//
// 查询商品池活动下面的商品
type AlibabahmmarketingitempoolqueryitemsAPIRequest struct {
	model.Params
	// 查询入参
	_param *ActivitySkuQuery
}

// NewAlibabahmmarketingitempoolqueryitemsRequest 初始化AlibabahmmarketingitempoolqueryitemsAPIRequest对象
func NewAlibabahmmarketingitempoolqueryitemsRequest() *AlibabahmmarketingitempoolqueryitemsAPIRequest {
	return &AlibabahmmarketingitempoolqueryitemsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingitempoolqueryitemsAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itempool.queryitems"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingitempoolqueryitemsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingitempoolqueryitemsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 查询入参
func (r *AlibabahmmarketingitempoolqueryitemsAPIRequest) SetParam(_param *ActivitySkuQuery) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabahmmarketingitempoolqueryitemsAPIRequest) GetParam() *ActivitySkuQuery {
	return r._param
}
