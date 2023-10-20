package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItemdiscountAdditemAPIRequest 报名特价商品 API请求
// alibaba.hm.marketing.itemdiscount.additem
//
// 在商品特价活动中报名特价商品
type AlibabaHmMarketingItemdiscountAdditemAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *ItemDiscountSku
	// 活动基本信息
	_param1 *CommonActivityParam
}

// NewAlibabaHmMarketingItemdiscountAdditemRequest 初始化AlibabaHmMarketingItemdiscountAdditemAPIRequest对象
func NewAlibabaHmMarketingItemdiscountAdditemRequest() *AlibabaHmMarketingItemdiscountAdditemAPIRequest {
	return &AlibabaHmMarketingItemdiscountAdditemAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingItemdiscountAdditemAPIRequest) Reset() {
	r._param0 = nil
	r._param1 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingItemdiscountAdditemAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itemdiscount.additem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingItemdiscountAdditemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingItemdiscountAdditemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 商品对象
func (r *AlibabaHmMarketingItemdiscountAdditemAPIRequest) SetParam0(_param0 *ItemDiscountSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaHmMarketingItemdiscountAdditemAPIRequest) GetParam0() *ItemDiscountSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 活动基本信息
func (r *AlibabaHmMarketingItemdiscountAdditemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaHmMarketingItemdiscountAdditemAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}

var poolAlibabaHmMarketingItemdiscountAdditemAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingItemdiscountAdditemRequest()
	},
}

// GetAlibabaHmMarketingItemdiscountAdditemRequest 从 sync.Pool 获取 AlibabaHmMarketingItemdiscountAdditemAPIRequest
func GetAlibabaHmMarketingItemdiscountAdditemAPIRequest() *AlibabaHmMarketingItemdiscountAdditemAPIRequest {
	return poolAlibabaHmMarketingItemdiscountAdditemAPIRequest.Get().(*AlibabaHmMarketingItemdiscountAdditemAPIRequest)
}

// ReleaseAlibabaHmMarketingItemdiscountAdditemAPIRequest 将 AlibabaHmMarketingItemdiscountAdditemAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingItemdiscountAdditemAPIRequest(v *AlibabaHmMarketingItemdiscountAdditemAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingItemdiscountAdditemAPIRequest.Put(v)
}
