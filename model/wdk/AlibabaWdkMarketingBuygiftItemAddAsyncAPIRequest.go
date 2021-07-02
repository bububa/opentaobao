package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest 批量发布买赠商品 API请求
// alibaba.wdk.marketing.buygift.item.add.async
//
// 批量发布买赠商品
type AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest struct {
	model.Params
	// sku信息
	_param0 []ItemBuyGiftSku
	// 系统自动生成
	_param1 *CommonActivityParam
	// alibaba.wdk.marketing.version.generate接口生成
	_version int64
}

// NewAlibabaWdkMarketingBuygiftItemAddAsyncRequest 初始化AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest对象
func NewAlibabaWdkMarketingBuygiftItemAddAsyncRequest() *AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest {
	return &AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.buygift.item.add.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// sku信息
func (r *AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest) SetParam0(_param0 []ItemBuyGiftSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest) GetParam0() []ItemBuyGiftSku {
	return r._param0
}

// Set is Param1 Setter
// 系统自动生成
func (r *AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// Get Param1 Getter
func (r AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}

// Set is Version Setter
// alibaba.wdk.marketing.version.generate接口生成
func (r *AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest) SetVersion(_version int64) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// Get Version Getter
func (r AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest) GetVersion() int64 {
	return r._version
}
