package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItempoolQueryactivityAPIRequest 查找商品池活动 API请求
// alibaba.wdk.marketing.itempool.queryactivity
//
// 查找商品池活动
type AlibabaWdkMarketingItempoolQueryactivityAPIRequest struct {
	model.Params
	// 查询商品池活动入参
	_param0 *CommonActivityParam
}

// NewAlibabaWdkMarketingItempoolQueryactivityRequest 初始化AlibabaWdkMarketingItempoolQueryactivityAPIRequest对象
func NewAlibabaWdkMarketingItempoolQueryactivityRequest() *AlibabaWdkMarketingItempoolQueryactivityAPIRequest {
	return &AlibabaWdkMarketingItempoolQueryactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolQueryactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itempool.queryactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolQueryactivityAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam0 is Param0 Setter
// 查询商品池活动入参
func (r *AlibabaWdkMarketingItempoolQueryactivityAPIRequest) SetParam0(_param0 *CommonActivityParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaWdkMarketingItempoolQueryactivityAPIRequest) GetParam0() *CommonActivityParam {
	return r._param0
}
