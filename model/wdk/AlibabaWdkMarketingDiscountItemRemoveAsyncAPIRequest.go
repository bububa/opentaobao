package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingDiscountItemRemoveAsyncAPIRequest 特价批量移除商品 API请求
// alibaba.wdk.marketing.discount.item.remove.async
//
// 特价批量移除商品
type AlibabaWdkMarketingDiscountItemRemoveAsyncAPIRequest struct {
	model.Params
	// sku信息
	_param0 []ItemDiscountSku
	// 系统自动生成
	_param1 *CommonActivityParam
	// alibaba.wdk.marketing.version.generate接口生成
	_version int64
}

// NewAlibabaWdkMarketingDiscountItemRemoveAsyncRequest 初始化AlibabaWdkMarketingDiscountItemRemoveAsyncAPIRequest对象
func NewAlibabaWdkMarketingDiscountItemRemoveAsyncRequest() *AlibabaWdkMarketingDiscountItemRemoveAsyncAPIRequest {
	return &AlibabaWdkMarketingDiscountItemRemoveAsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingDiscountItemRemoveAsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.discount.item.remove.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingDiscountItemRemoveAsyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// sku信息
func (r *AlibabaWdkMarketingDiscountItemRemoveAsyncAPIRequest) SetParam0(_param0 []ItemDiscountSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r AlibabaWdkMarketingDiscountItemRemoveAsyncAPIRequest) GetParam0() []ItemDiscountSku {
	return r._param0
}

// Set is Param1 Setter
// 系统自动生成
func (r *AlibabaWdkMarketingDiscountItemRemoveAsyncAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// Get Param1 Getter
func (r AlibabaWdkMarketingDiscountItemRemoveAsyncAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}

// Set is Version Setter
// alibaba.wdk.marketing.version.generate接口生成
func (r *AlibabaWdkMarketingDiscountItemRemoveAsyncAPIRequest) SetVersion(_version int64) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// Get Version Getter
func (r AlibabaWdkMarketingDiscountItemRemoveAsyncAPIRequest) GetVersion() int64 {
	return r._version
}
