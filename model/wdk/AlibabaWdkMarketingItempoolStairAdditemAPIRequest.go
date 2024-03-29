package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItempoolStairAdditemAPIRequest 商品池阶梯商品添加 API请求
// alibaba.wdk.marketing.itempool.stair.additem
//
// 添加商品池阶梯商品
type AlibabaWdkMarketingItempoolStairAdditemAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *ItemPoolSku
	// 系统自动生成
	_param1 *CommonActivityParam
}

// NewAlibabaWdkMarketingItempoolStairAdditemRequest 初始化AlibabaWdkMarketingItempoolStairAdditemAPIRequest对象
func NewAlibabaWdkMarketingItempoolStairAdditemRequest() *AlibabaWdkMarketingItempoolStairAdditemAPIRequest {
	return &AlibabaWdkMarketingItempoolStairAdditemAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingItempoolStairAdditemAPIRequest) Reset() {
	r._param0 = nil
	r._param1 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolStairAdditemAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itempool.stair.additem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolStairAdditemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingItempoolStairAdditemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 系统自动生成
func (r *AlibabaWdkMarketingItempoolStairAdditemAPIRequest) SetParam0(_param0 *ItemPoolSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaWdkMarketingItempoolStairAdditemAPIRequest) GetParam0() *ItemPoolSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 系统自动生成
func (r *AlibabaWdkMarketingItempoolStairAdditemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaWdkMarketingItempoolStairAdditemAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}

var poolAlibabaWdkMarketingItempoolStairAdditemAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingItempoolStairAdditemRequest()
	},
}

// GetAlibabaWdkMarketingItempoolStairAdditemRequest 从 sync.Pool 获取 AlibabaWdkMarketingItempoolStairAdditemAPIRequest
func GetAlibabaWdkMarketingItempoolStairAdditemAPIRequest() *AlibabaWdkMarketingItempoolStairAdditemAPIRequest {
	return poolAlibabaWdkMarketingItempoolStairAdditemAPIRequest.Get().(*AlibabaWdkMarketingItempoolStairAdditemAPIRequest)
}

// ReleaseAlibabaWdkMarketingItempoolStairAdditemAPIRequest 将 AlibabaWdkMarketingItempoolStairAdditemAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingItempoolStairAdditemAPIRequest(v *AlibabaWdkMarketingItempoolStairAdditemAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingItempoolStairAdditemAPIRequest.Put(v)
}
