package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItempoolRemoveitemAPIRequest 移除商品池里面的商品 API请求
// alibaba.wdk.marketing.itempool.removeitem
//
// 移除商品池里面的商品
type AlibabaWdkMarketingItempoolRemoveitemAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *ItemPoolSku
	// 活动基本信息
	_param1 *CommonActivityParam
}

// NewAlibabaWdkMarketingItempoolRemoveitemRequest 初始化AlibabaWdkMarketingItempoolRemoveitemAPIRequest对象
func NewAlibabaWdkMarketingItempoolRemoveitemRequest() *AlibabaWdkMarketingItempoolRemoveitemAPIRequest {
	return &AlibabaWdkMarketingItempoolRemoveitemAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingItempoolRemoveitemAPIRequest) Reset() {
	r._param0 = nil
	r._param1 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolRemoveitemAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itempool.removeitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolRemoveitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingItempoolRemoveitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 商品对象
func (r *AlibabaWdkMarketingItempoolRemoveitemAPIRequest) SetParam0(_param0 *ItemPoolSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaWdkMarketingItempoolRemoveitemAPIRequest) GetParam0() *ItemPoolSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 活动基本信息
func (r *AlibabaWdkMarketingItempoolRemoveitemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaWdkMarketingItempoolRemoveitemAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}

var poolAlibabaWdkMarketingItempoolRemoveitemAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingItempoolRemoveitemRequest()
	},
}

// GetAlibabaWdkMarketingItempoolRemoveitemRequest 从 sync.Pool 获取 AlibabaWdkMarketingItempoolRemoveitemAPIRequest
func GetAlibabaWdkMarketingItempoolRemoveitemAPIRequest() *AlibabaWdkMarketingItempoolRemoveitemAPIRequest {
	return poolAlibabaWdkMarketingItempoolRemoveitemAPIRequest.Get().(*AlibabaWdkMarketingItempoolRemoveitemAPIRequest)
}

// ReleaseAlibabaWdkMarketingItempoolRemoveitemAPIRequest 将 AlibabaWdkMarketingItempoolRemoveitemAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingItempoolRemoveitemAPIRequest(v *AlibabaWdkMarketingItempoolRemoveitemAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingItempoolRemoveitemAPIRequest.Put(v)
}
