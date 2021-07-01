package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest
删除买赠活动 API请求
alibaba.wdk.marketing.itembuygift.deleteactivity

删除买赠活动 */
type AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest struct {
	model.Params
	// 要删除的活动信息
	_param *CommonActivityParam
}

// NewAlibabaWdkMarketingItembuygiftDeleteactivityRequest 初始化AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest对象
func NewAlibabaWdkMarketingItembuygiftDeleteactivityRequest() *AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest {
	return &AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itembuygift.deleteactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 要删除的活动信息
func (r *AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest) SetParam(_param *CommonActivityParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest) GetParam() *CommonActivityParam {
	return r._param
}
