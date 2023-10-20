package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItempoolStairRemoveitemAPIRequest 删除换购活动商品 API请求
// alibaba.hm.marketing.itempool.stair.removeitem
//
// 删除换购商品
type AlibabaHmMarketingItempoolStairRemoveitemAPIRequest struct {
	model.Params
	// 商品sku信息
	_param0 *ItemPoolSku
	// 活动信息
	_param1 *CommonActivityParam
}

// NewAlibabaHmMarketingItempoolStairRemoveitemRequest 初始化AlibabaHmMarketingItempoolStairRemoveitemAPIRequest对象
func NewAlibabaHmMarketingItempoolStairRemoveitemRequest() *AlibabaHmMarketingItempoolStairRemoveitemAPIRequest {
	return &AlibabaHmMarketingItempoolStairRemoveitemAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingItempoolStairRemoveitemAPIRequest) Reset() {
	r._param0 = nil
	r._param1 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingItempoolStairRemoveitemAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itempool.stair.removeitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingItempoolStairRemoveitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingItempoolStairRemoveitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 商品sku信息
func (r *AlibabaHmMarketingItempoolStairRemoveitemAPIRequest) SetParam0(_param0 *ItemPoolSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaHmMarketingItempoolStairRemoveitemAPIRequest) GetParam0() *ItemPoolSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 活动信息
func (r *AlibabaHmMarketingItempoolStairRemoveitemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaHmMarketingItempoolStairRemoveitemAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}

var poolAlibabaHmMarketingItempoolStairRemoveitemAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingItempoolStairRemoveitemRequest()
	},
}

// GetAlibabaHmMarketingItempoolStairRemoveitemRequest 从 sync.Pool 获取 AlibabaHmMarketingItempoolStairRemoveitemAPIRequest
func GetAlibabaHmMarketingItempoolStairRemoveitemAPIRequest() *AlibabaHmMarketingItempoolStairRemoveitemAPIRequest {
	return poolAlibabaHmMarketingItempoolStairRemoveitemAPIRequest.Get().(*AlibabaHmMarketingItempoolStairRemoveitemAPIRequest)
}

// ReleaseAlibabaHmMarketingItempoolStairRemoveitemAPIRequest 将 AlibabaHmMarketingItempoolStairRemoveitemAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingItempoolStairRemoveitemAPIRequest(v *AlibabaHmMarketingItempoolStairRemoveitemAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingItempoolStairRemoveitemAPIRequest.Put(v)
}
