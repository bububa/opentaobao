package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItempoolDeleteactivityAPIRequest 删除商品池活动 API请求
// alibaba.wdk.marketing.itempool.deleteactivity
//
// 删除商品池活动
type AlibabaWdkMarketingItempoolDeleteactivityAPIRequest struct {
	model.Params
	// 需要删除的活动的信息
	_param *CommonActivityParam
}

// NewAlibabaWdkMarketingItempoolDeleteactivityRequest 初始化AlibabaWdkMarketingItempoolDeleteactivityAPIRequest对象
func NewAlibabaWdkMarketingItempoolDeleteactivityRequest() *AlibabaWdkMarketingItempoolDeleteactivityAPIRequest {
	return &AlibabaWdkMarketingItempoolDeleteactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolDeleteactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itempool.deleteactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolDeleteactivityAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 需要删除的活动的信息
func (r *AlibabaWdkMarketingItempoolDeleteactivityAPIRequest) SetParam(_param *CommonActivityParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkMarketingItempoolDeleteactivityAPIRequest) GetParam() *CommonActivityParam {
	return r._param
}
