package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItembuygiftCreateactivityAPIRequest 创建买赠活动 API请求
// alibaba.wdk.marketing.itembuygift.createactivity
//
// 创建买赠活动
type AlibabaWdkMarketingItembuygiftCreateactivityAPIRequest struct {
	model.Params
	// 创建活动请求入参
	_param *ItemBuyGiftActivity
}

// NewAlibabaWdkMarketingItembuygiftCreateactivityRequest 初始化AlibabaWdkMarketingItembuygiftCreateactivityAPIRequest对象
func NewAlibabaWdkMarketingItembuygiftCreateactivityRequest() *AlibabaWdkMarketingItembuygiftCreateactivityAPIRequest {
	return &AlibabaWdkMarketingItembuygiftCreateactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItembuygiftCreateactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itembuygift.createactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItembuygiftCreateactivityAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 创建活动请求入参
func (r *AlibabaWdkMarketingItembuygiftCreateactivityAPIRequest) SetParam(_param *ItemBuyGiftActivity) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r AlibabaWdkMarketingItembuygiftCreateactivityAPIRequest) GetParam() *ItemBuyGiftActivity {
	return r._param
}
