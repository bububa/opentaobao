package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingDiscountItemAddAsyncAPIRequest 特价批量新增商品 API请求
// alibaba.hm.marketing.discount.item.add.async
//
// 新分组模型下新增商品
type AlibabaHmMarketingDiscountItemAddAsyncAPIRequest struct {
	model.Params
	// sku信息
	_param0 []ItemDiscountSku
	// 系统自动生成
	_param1 *CommonActivityParam
	// alibaba.wdk.marketing.version.generate接口生成
	_version int64
}

// NewAlibabaHmMarketingDiscountItemAddAsyncRequest 初始化AlibabaHmMarketingDiscountItemAddAsyncAPIRequest对象
func NewAlibabaHmMarketingDiscountItemAddAsyncRequest() *AlibabaHmMarketingDiscountItemAddAsyncAPIRequest {
	return &AlibabaHmMarketingDiscountItemAddAsyncAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingDiscountItemAddAsyncAPIRequest) Reset() {
	r._param0 = r._param0[:0]
	r._param1 = nil
	r._version = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingDiscountItemAddAsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.discount.item.add.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingDiscountItemAddAsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingDiscountItemAddAsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// sku信息
func (r *AlibabaHmMarketingDiscountItemAddAsyncAPIRequest) SetParam0(_param0 []ItemDiscountSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaHmMarketingDiscountItemAddAsyncAPIRequest) GetParam0() []ItemDiscountSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 系统自动生成
func (r *AlibabaHmMarketingDiscountItemAddAsyncAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaHmMarketingDiscountItemAddAsyncAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}

// SetVersion is Version Setter
// alibaba.wdk.marketing.version.generate接口生成
func (r *AlibabaHmMarketingDiscountItemAddAsyncAPIRequest) SetVersion(_version int64) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r AlibabaHmMarketingDiscountItemAddAsyncAPIRequest) GetVersion() int64 {
	return r._version
}

var poolAlibabaHmMarketingDiscountItemAddAsyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingDiscountItemAddAsyncRequest()
	},
}

// GetAlibabaHmMarketingDiscountItemAddAsyncRequest 从 sync.Pool 获取 AlibabaHmMarketingDiscountItemAddAsyncAPIRequest
func GetAlibabaHmMarketingDiscountItemAddAsyncAPIRequest() *AlibabaHmMarketingDiscountItemAddAsyncAPIRequest {
	return poolAlibabaHmMarketingDiscountItemAddAsyncAPIRequest.Get().(*AlibabaHmMarketingDiscountItemAddAsyncAPIRequest)
}

// ReleaseAlibabaHmMarketingDiscountItemAddAsyncAPIRequest 将 AlibabaHmMarketingDiscountItemAddAsyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingDiscountItemAddAsyncAPIRequest(v *AlibabaHmMarketingDiscountItemAddAsyncAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingDiscountItemAddAsyncAPIRequest.Put(v)
}
