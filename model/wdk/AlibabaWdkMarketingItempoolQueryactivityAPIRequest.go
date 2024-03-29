package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItempoolQueryactivityAPIRequest 查找商品池活动 API请求
// alibaba.wdk.marketing.itempool.queryactivity
//
// 查找商品池活动
type AlibabaWdkMarketingItempoolQueryactivityAPIRequest struct {
	model.Params
	// 查询商品池活动入参
	_param0 *CommonActivityParam
}

// NewAlibabaWdkMarketingItempoolQueryactivityRequest 初始化AlibabaWdkMarketingItempoolQueryactivityAPIRequest对象
func NewAlibabaWdkMarketingItempoolQueryactivityRequest() *AlibabaWdkMarketingItempoolQueryactivityAPIRequest {
	return &AlibabaWdkMarketingItempoolQueryactivityAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingItempoolQueryactivityAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolQueryactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itempool.queryactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolQueryactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingItempoolQueryactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 查询商品池活动入参
func (r *AlibabaWdkMarketingItempoolQueryactivityAPIRequest) SetParam0(_param0 *CommonActivityParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaWdkMarketingItempoolQueryactivityAPIRequest) GetParam0() *CommonActivityParam {
	return r._param0
}

var poolAlibabaWdkMarketingItempoolQueryactivityAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingItempoolQueryactivityRequest()
	},
}

// GetAlibabaWdkMarketingItempoolQueryactivityRequest 从 sync.Pool 获取 AlibabaWdkMarketingItempoolQueryactivityAPIRequest
func GetAlibabaWdkMarketingItempoolQueryactivityAPIRequest() *AlibabaWdkMarketingItempoolQueryactivityAPIRequest {
	return poolAlibabaWdkMarketingItempoolQueryactivityAPIRequest.Get().(*AlibabaWdkMarketingItempoolQueryactivityAPIRequest)
}

// ReleaseAlibabaWdkMarketingItempoolQueryactivityAPIRequest 将 AlibabaWdkMarketingItempoolQueryactivityAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingItempoolQueryactivityAPIRequest(v *AlibabaWdkMarketingItempoolQueryactivityAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingItempoolQueryactivityAPIRequest.Put(v)
}
