package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItemdiscountRemoveitemAPIRequest 移除报名的商品 API请求
// alibaba.hm.marketing.itemdiscount.removeitem
//
// 将报名特价活动的商品从特价活动中移除
type AlibabaHmMarketingItemdiscountRemoveitemAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *ItemDiscountSku
	// 活动基本信息
	_param1 *CommonActivityParam
}

// NewAlibabaHmMarketingItemdiscountRemoveitemRequest 初始化AlibabaHmMarketingItemdiscountRemoveitemAPIRequest对象
func NewAlibabaHmMarketingItemdiscountRemoveitemRequest() *AlibabaHmMarketingItemdiscountRemoveitemAPIRequest {
	return &AlibabaHmMarketingItemdiscountRemoveitemAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingItemdiscountRemoveitemAPIRequest) Reset() {
	r._param0 = nil
	r._param1 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingItemdiscountRemoveitemAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itemdiscount.removeitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingItemdiscountRemoveitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingItemdiscountRemoveitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 商品对象
func (r *AlibabaHmMarketingItemdiscountRemoveitemAPIRequest) SetParam0(_param0 *ItemDiscountSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaHmMarketingItemdiscountRemoveitemAPIRequest) GetParam0() *ItemDiscountSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 活动基本信息
func (r *AlibabaHmMarketingItemdiscountRemoveitemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaHmMarketingItemdiscountRemoveitemAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}

var poolAlibabaHmMarketingItemdiscountRemoveitemAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingItemdiscountRemoveitemRequest()
	},
}

// GetAlibabaHmMarketingItemdiscountRemoveitemRequest 从 sync.Pool 获取 AlibabaHmMarketingItemdiscountRemoveitemAPIRequest
func GetAlibabaHmMarketingItemdiscountRemoveitemAPIRequest() *AlibabaHmMarketingItemdiscountRemoveitemAPIRequest {
	return poolAlibabaHmMarketingItemdiscountRemoveitemAPIRequest.Get().(*AlibabaHmMarketingItemdiscountRemoveitemAPIRequest)
}

// ReleaseAlibabaHmMarketingItemdiscountRemoveitemAPIRequest 将 AlibabaHmMarketingItemdiscountRemoveitemAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingItemdiscountRemoveitemAPIRequest(v *AlibabaHmMarketingItemdiscountRemoveitemAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingItemdiscountRemoveitemAPIRequest.Put(v)
}
