package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItempoolAdditemAPIRequest 增加商品池里面的商品 API请求
// alibaba.hm.marketing.itempool.additem
//
// 增加商品池里面的商品
type AlibabaHmMarketingItempoolAdditemAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *ItemPoolSku
	// 活动基本信息
	_param1 *CommonActivityParam
}

// NewAlibabaHmMarketingItempoolAdditemRequest 初始化AlibabaHmMarketingItempoolAdditemAPIRequest对象
func NewAlibabaHmMarketingItempoolAdditemRequest() *AlibabaHmMarketingItempoolAdditemAPIRequest {
	return &AlibabaHmMarketingItempoolAdditemAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingItempoolAdditemAPIRequest) Reset() {
	r._param0 = nil
	r._param1 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingItempoolAdditemAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itempool.additem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingItempoolAdditemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingItempoolAdditemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 商品对象
func (r *AlibabaHmMarketingItempoolAdditemAPIRequest) SetParam0(_param0 *ItemPoolSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaHmMarketingItempoolAdditemAPIRequest) GetParam0() *ItemPoolSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 活动基本信息
func (r *AlibabaHmMarketingItempoolAdditemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaHmMarketingItempoolAdditemAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}

var poolAlibabaHmMarketingItempoolAdditemAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingItempoolAdditemRequest()
	},
}

// GetAlibabaHmMarketingItempoolAdditemRequest 从 sync.Pool 获取 AlibabaHmMarketingItempoolAdditemAPIRequest
func GetAlibabaHmMarketingItempoolAdditemAPIRequest() *AlibabaHmMarketingItempoolAdditemAPIRequest {
	return poolAlibabaHmMarketingItempoolAdditemAPIRequest.Get().(*AlibabaHmMarketingItempoolAdditemAPIRequest)
}

// ReleaseAlibabaHmMarketingItempoolAdditemAPIRequest 将 AlibabaHmMarketingItempoolAdditemAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingItempoolAdditemAPIRequest(v *AlibabaHmMarketingItempoolAdditemAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingItempoolAdditemAPIRequest.Put(v)
}
