package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItempoolRemoveitemAPIRequest 移除商品池里面的商品 API请求
// alibaba.hm.marketing.itempool.removeitem
//
// 移除商品池里面的商品
type AlibabaHmMarketingItempoolRemoveitemAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *ItemPoolSku
	// 活动基本信息
	_param1 *CommonActivityParam
}

// NewAlibabaHmMarketingItempoolRemoveitemRequest 初始化AlibabaHmMarketingItempoolRemoveitemAPIRequest对象
func NewAlibabaHmMarketingItempoolRemoveitemRequest() *AlibabaHmMarketingItempoolRemoveitemAPIRequest {
	return &AlibabaHmMarketingItempoolRemoveitemAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingItempoolRemoveitemAPIRequest) Reset() {
	r._param0 = nil
	r._param1 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingItempoolRemoveitemAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itempool.removeitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingItempoolRemoveitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingItempoolRemoveitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 商品对象
func (r *AlibabaHmMarketingItempoolRemoveitemAPIRequest) SetParam0(_param0 *ItemPoolSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaHmMarketingItempoolRemoveitemAPIRequest) GetParam0() *ItemPoolSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 活动基本信息
func (r *AlibabaHmMarketingItempoolRemoveitemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaHmMarketingItempoolRemoveitemAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}

var poolAlibabaHmMarketingItempoolRemoveitemAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingItempoolRemoveitemRequest()
	},
}

// GetAlibabaHmMarketingItempoolRemoveitemRequest 从 sync.Pool 获取 AlibabaHmMarketingItempoolRemoveitemAPIRequest
func GetAlibabaHmMarketingItempoolRemoveitemAPIRequest() *AlibabaHmMarketingItempoolRemoveitemAPIRequest {
	return poolAlibabaHmMarketingItempoolRemoveitemAPIRequest.Get().(*AlibabaHmMarketingItempoolRemoveitemAPIRequest)
}

// ReleaseAlibabaHmMarketingItempoolRemoveitemAPIRequest 将 AlibabaHmMarketingItempoolRemoveitemAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingItempoolRemoveitemAPIRequest(v *AlibabaHmMarketingItempoolRemoveitemAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingItempoolRemoveitemAPIRequest.Put(v)
}
