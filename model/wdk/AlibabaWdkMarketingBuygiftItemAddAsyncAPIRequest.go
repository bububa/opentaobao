package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest) Reset() {
	r._param0 = r._param0[:0]
	r._param1 = nil
	r._version = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.buygift.item.add.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// sku信息
func (r *AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest) SetParam0(_param0 []ItemBuyGiftSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest) GetParam0() []ItemBuyGiftSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 系统自动生成
func (r *AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}

// SetVersion is Version Setter
// alibaba.wdk.marketing.version.generate接口生成
func (r *AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest) SetVersion(_version int64) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest) GetVersion() int64 {
	return r._version
}

var poolAlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingBuygiftItemAddAsyncRequest()
	},
}

// GetAlibabaWdkMarketingBuygiftItemAddAsyncRequest 从 sync.Pool 获取 AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest
func GetAlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest() *AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest {
	return poolAlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest.Get().(*AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest)
}

// ReleaseAlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest 将 AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest(v *AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest.Put(v)
}
