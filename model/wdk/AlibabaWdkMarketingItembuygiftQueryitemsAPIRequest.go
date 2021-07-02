package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItembuygiftQueryitemsAPIRequest 查询买赠活动下的商品 API请求
// alibaba.wdk.marketing.itembuygift.queryitems
//
// 查询买赠活动下的商品
type AlibabaWdkMarketingItembuygiftQueryitemsAPIRequest struct {
	model.Params
	// 查询入参
	_param *ActivitySkuQuery
}

// NewAlibabaWdkMarketingItembuygiftQueryitemsRequest 初始化AlibabaWdkMarketingItembuygiftQueryitemsAPIRequest对象
func NewAlibabaWdkMarketingItembuygiftQueryitemsRequest() *AlibabaWdkMarketingItembuygiftQueryitemsAPIRequest {
	return &AlibabaWdkMarketingItembuygiftQueryitemsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItembuygiftQueryitemsAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itembuygift.queryitems"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItembuygiftQueryitemsAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 查询入参
func (r *AlibabaWdkMarketingItembuygiftQueryitemsAPIRequest) SetParam(_param *ActivitySkuQuery) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r AlibabaWdkMarketingItembuygiftQueryitemsAPIRequest) GetParam() *ActivitySkuQuery {
	return r._param
}
