package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItembuygiftQueryitemsAPIRequest 查询买赠活动下的商品 API请求
// alibaba.hm.marketing.itembuygift.queryitems
//
// 查询买赠活动下的商品
type AlibabaHmMarketingItembuygiftQueryitemsAPIRequest struct {
	model.Params
	// 查询入参
	_param *ActivitySkuQuery
}

// NewAlibabaHmMarketingItembuygiftQueryitemsRequest 初始化AlibabaHmMarketingItembuygiftQueryitemsAPIRequest对象
func NewAlibabaHmMarketingItembuygiftQueryitemsRequest() *AlibabaHmMarketingItembuygiftQueryitemsAPIRequest {
	return &AlibabaHmMarketingItembuygiftQueryitemsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingItembuygiftQueryitemsAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itembuygift.queryitems"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingItembuygiftQueryitemsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingItembuygiftQueryitemsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 查询入参
func (r *AlibabaHmMarketingItembuygiftQueryitemsAPIRequest) SetParam(_param *ActivitySkuQuery) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaHmMarketingItembuygiftQueryitemsAPIRequest) GetParam() *ActivitySkuQuery {
	return r._param
}
