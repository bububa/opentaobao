package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItempoolQueryitemsAPIRequest 查询商品池活动下面的商品 API请求
// alibaba.hm.marketing.itempool.queryitems
//
// 查询商品池活动下面的商品
type AlibabaHmMarketingItempoolQueryitemsAPIRequest struct {
	model.Params
	// 查询入参
	_param *ActivitySkuQuery
}

// NewAlibabaHmMarketingItempoolQueryitemsRequest 初始化AlibabaHmMarketingItempoolQueryitemsAPIRequest对象
func NewAlibabaHmMarketingItempoolQueryitemsRequest() *AlibabaHmMarketingItempoolQueryitemsAPIRequest {
	return &AlibabaHmMarketingItempoolQueryitemsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingItempoolQueryitemsAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itempool.queryitems"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingItempoolQueryitemsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingItempoolQueryitemsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 查询入参
func (r *AlibabaHmMarketingItempoolQueryitemsAPIRequest) SetParam(_param *ActivitySkuQuery) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaHmMarketingItempoolQueryitemsAPIRequest) GetParam() *ActivitySkuQuery {
	return r._param
}
