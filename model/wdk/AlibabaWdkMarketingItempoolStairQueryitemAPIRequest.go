package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItempoolStairQueryitemAPIRequest 换购商品查询 API请求
// alibaba.wdk.marketing.itempool.stair.queryitem
//
// 换购商品查询
type AlibabaWdkMarketingItempoolStairQueryitemAPIRequest struct {
	model.Params
	// 换购商品查询参数
	_param0 *ActivitySkuQuery
}

// NewAlibabaWdkMarketingItempoolStairQueryitemRequest 初始化AlibabaWdkMarketingItempoolStairQueryitemAPIRequest对象
func NewAlibabaWdkMarketingItempoolStairQueryitemRequest() *AlibabaWdkMarketingItempoolStairQueryitemAPIRequest {
	return &AlibabaWdkMarketingItempoolStairQueryitemAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingItempoolStairQueryitemAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolStairQueryitemAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itempool.stair.queryitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolStairQueryitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingItempoolStairQueryitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 换购商品查询参数
func (r *AlibabaWdkMarketingItempoolStairQueryitemAPIRequest) SetParam0(_param0 *ActivitySkuQuery) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaWdkMarketingItempoolStairQueryitemAPIRequest) GetParam0() *ActivitySkuQuery {
	return r._param0
}

var poolAlibabaWdkMarketingItempoolStairQueryitemAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingItempoolStairQueryitemRequest()
	},
}

// GetAlibabaWdkMarketingItempoolStairQueryitemRequest 从 sync.Pool 获取 AlibabaWdkMarketingItempoolStairQueryitemAPIRequest
func GetAlibabaWdkMarketingItempoolStairQueryitemAPIRequest() *AlibabaWdkMarketingItempoolStairQueryitemAPIRequest {
	return poolAlibabaWdkMarketingItempoolStairQueryitemAPIRequest.Get().(*AlibabaWdkMarketingItempoolStairQueryitemAPIRequest)
}

// ReleaseAlibabaWdkMarketingItempoolStairQueryitemAPIRequest 将 AlibabaWdkMarketingItempoolStairQueryitemAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingItempoolStairQueryitemAPIRequest(v *AlibabaWdkMarketingItempoolStairQueryitemAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingItempoolStairQueryitemAPIRequest.Put(v)
}
