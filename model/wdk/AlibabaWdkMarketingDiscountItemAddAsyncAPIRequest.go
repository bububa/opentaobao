package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingDiscountItemAddAsyncAPIRequest 特价批量新增商品 API请求
// alibaba.wdk.marketing.discount.item.add.async
//
// 新分组模型下新增商品
type AlibabaWdkMarketingDiscountItemAddAsyncAPIRequest struct {
	model.Params
	// sku信息
	_param0 []ItemDiscountSku
	// 系统自动生成
	_param1 *CommonActivityParam
	// alibaba.wdk.marketing.version.generate接口生成
	_version int64
}

// NewAlibabaWdkMarketingDiscountItemAddAsyncRequest 初始化AlibabaWdkMarketingDiscountItemAddAsyncAPIRequest对象
func NewAlibabaWdkMarketingDiscountItemAddAsyncRequest() *AlibabaWdkMarketingDiscountItemAddAsyncAPIRequest {
	return &AlibabaWdkMarketingDiscountItemAddAsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingDiscountItemAddAsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.discount.item.add.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingDiscountItemAddAsyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// sku信息
func (r *AlibabaWdkMarketingDiscountItemAddAsyncAPIRequest) SetParam0(_param0 []ItemDiscountSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r AlibabaWdkMarketingDiscountItemAddAsyncAPIRequest) GetParam0() []ItemDiscountSku {
	return r._param0
}

// Set is Param1 Setter
// 系统自动生成
func (r *AlibabaWdkMarketingDiscountItemAddAsyncAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// Get Param1 Getter
func (r AlibabaWdkMarketingDiscountItemAddAsyncAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}

// Set is Version Setter
// alibaba.wdk.marketing.version.generate接口生成
func (r *AlibabaWdkMarketingDiscountItemAddAsyncAPIRequest) SetVersion(_version int64) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// Get Version Getter
func (r AlibabaWdkMarketingDiscountItemAddAsyncAPIRequest) GetVersion() int64 {
	return r._version
}
