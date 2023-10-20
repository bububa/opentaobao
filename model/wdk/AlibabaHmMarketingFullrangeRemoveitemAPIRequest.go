package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingFullrangeRemoveitemAPIRequest 全场活动删除购品 API请求
// alibaba.hm.marketing.fullrange.removeitem
//
// 删除换购商品
type AlibabaHmMarketingFullrangeRemoveitemAPIRequest struct {
	model.Params
	// 商品sku信息
	_param0 *ItemStairSku
	// 活动信息
	_param1 *CommonActivityParam
}

// NewAlibabaHmMarketingFullrangeRemoveitemRequest 初始化AlibabaHmMarketingFullrangeRemoveitemAPIRequest对象
func NewAlibabaHmMarketingFullrangeRemoveitemRequest() *AlibabaHmMarketingFullrangeRemoveitemAPIRequest {
	return &AlibabaHmMarketingFullrangeRemoveitemAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingFullrangeRemoveitemAPIRequest) Reset() {
	r._param0 = nil
	r._param1 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingFullrangeRemoveitemAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.fullrange.removeitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingFullrangeRemoveitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingFullrangeRemoveitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 商品sku信息
func (r *AlibabaHmMarketingFullrangeRemoveitemAPIRequest) SetParam0(_param0 *ItemStairSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaHmMarketingFullrangeRemoveitemAPIRequest) GetParam0() *ItemStairSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 活动信息
func (r *AlibabaHmMarketingFullrangeRemoveitemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaHmMarketingFullrangeRemoveitemAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}

var poolAlibabaHmMarketingFullrangeRemoveitemAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingFullrangeRemoveitemRequest()
	},
}

// GetAlibabaHmMarketingFullrangeRemoveitemRequest 从 sync.Pool 获取 AlibabaHmMarketingFullrangeRemoveitemAPIRequest
func GetAlibabaHmMarketingFullrangeRemoveitemAPIRequest() *AlibabaHmMarketingFullrangeRemoveitemAPIRequest {
	return poolAlibabaHmMarketingFullrangeRemoveitemAPIRequest.Get().(*AlibabaHmMarketingFullrangeRemoveitemAPIRequest)
}

// ReleaseAlibabaHmMarketingFullrangeRemoveitemAPIRequest 将 AlibabaHmMarketingFullrangeRemoveitemAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingFullrangeRemoveitemAPIRequest(v *AlibabaHmMarketingFullrangeRemoveitemAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingFullrangeRemoveitemAPIRequest.Put(v)
}
