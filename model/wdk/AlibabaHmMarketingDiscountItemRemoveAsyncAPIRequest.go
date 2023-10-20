package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingDiscountItemRemoveAsyncAPIRequest 特价批量移除商品 API请求
// alibaba.hm.marketing.discount.item.remove.async
//
// 特价批量移除商品
type AlibabaHmMarketingDiscountItemRemoveAsyncAPIRequest struct {
	model.Params
	// sku信息
	_param0 []ItemDiscountSku
	// 系统自动生成
	_param1 *CommonActivityParam
	// alibaba.wdk.marketing.version.generate接口生成
	_version int64
}

// NewAlibabaHmMarketingDiscountItemRemoveAsyncRequest 初始化AlibabaHmMarketingDiscountItemRemoveAsyncAPIRequest对象
func NewAlibabaHmMarketingDiscountItemRemoveAsyncRequest() *AlibabaHmMarketingDiscountItemRemoveAsyncAPIRequest {
	return &AlibabaHmMarketingDiscountItemRemoveAsyncAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingDiscountItemRemoveAsyncAPIRequest) Reset() {
	r._param0 = r._param0[:0]
	r._param1 = nil
	r._version = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingDiscountItemRemoveAsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.discount.item.remove.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingDiscountItemRemoveAsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingDiscountItemRemoveAsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// sku信息
func (r *AlibabaHmMarketingDiscountItemRemoveAsyncAPIRequest) SetParam0(_param0 []ItemDiscountSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaHmMarketingDiscountItemRemoveAsyncAPIRequest) GetParam0() []ItemDiscountSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 系统自动生成
func (r *AlibabaHmMarketingDiscountItemRemoveAsyncAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaHmMarketingDiscountItemRemoveAsyncAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}

// SetVersion is Version Setter
// alibaba.wdk.marketing.version.generate接口生成
func (r *AlibabaHmMarketingDiscountItemRemoveAsyncAPIRequest) SetVersion(_version int64) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r AlibabaHmMarketingDiscountItemRemoveAsyncAPIRequest) GetVersion() int64 {
	return r._version
}

var poolAlibabaHmMarketingDiscountItemRemoveAsyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingDiscountItemRemoveAsyncRequest()
	},
}

// GetAlibabaHmMarketingDiscountItemRemoveAsyncRequest 从 sync.Pool 获取 AlibabaHmMarketingDiscountItemRemoveAsyncAPIRequest
func GetAlibabaHmMarketingDiscountItemRemoveAsyncAPIRequest() *AlibabaHmMarketingDiscountItemRemoveAsyncAPIRequest {
	return poolAlibabaHmMarketingDiscountItemRemoveAsyncAPIRequest.Get().(*AlibabaHmMarketingDiscountItemRemoveAsyncAPIRequest)
}

// ReleaseAlibabaHmMarketingDiscountItemRemoveAsyncAPIRequest 将 AlibabaHmMarketingDiscountItemRemoveAsyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingDiscountItemRemoveAsyncAPIRequest(v *AlibabaHmMarketingDiscountItemRemoveAsyncAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingDiscountItemRemoveAsyncAPIRequest.Put(v)
}
