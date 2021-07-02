package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest 创建商品特价活动 API请求
// alibaba.wdk.marketing.itemdiscount.createactivity
//
// 创建商品特价活动
type AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest struct {
	model.Params
	// 创建活动请求入参
	_param *ItemDiscountActivityRequest
}

// NewAlibabaWdkMarketingItemdiscountCreateactivityRequest 初始化AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest对象
func NewAlibabaWdkMarketingItemdiscountCreateactivityRequest() *AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest {
	return &AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itemdiscount.createactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 创建活动请求入参
func (r *AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest) SetParam(_param *ItemDiscountActivityRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest) GetParam() *ItemDiscountActivityRequest {
	return r._param
}
