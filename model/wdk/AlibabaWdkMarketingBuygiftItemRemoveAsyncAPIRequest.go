package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIRequest 批量删除买赠商品 API请求
// alibaba.wdk.marketing.buygift.item.remove.async
//
// 批量删除买赠商品
type AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIRequest struct {
	model.Params
	// sku信息
	_param0 []ItemBuyGiftSku
	// 系统自动生成
	_param1 *CommonActivityParam
	// alibaba.wdk.marketing.version.generate接口生成
	_version int64
}

// NewAlibabaWdkMarketingBuygiftItemRemoveAsyncRequest 初始化AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIRequest对象
func NewAlibabaWdkMarketingBuygiftItemRemoveAsyncRequest() *AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIRequest {
	return &AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.buygift.item.remove.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam0 is Param0 Setter
// sku信息
func (r *AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIRequest) SetParam0(_param0 []ItemBuyGiftSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIRequest) GetParam0() []ItemBuyGiftSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 系统自动生成
func (r *AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}

// SetVersion is Version Setter
// alibaba.wdk.marketing.version.generate接口生成
func (r *AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIRequest) SetVersion(_version int64) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIRequest) GetVersion() int64 {
	return r._version
}
