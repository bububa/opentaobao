package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItempoolDeleteactivityAPIRequest 删除商品池活动 API请求
// alibaba.wdk.marketing.itempool.deleteactivity
//
// 删除商品池活动
type AlibabaWdkMarketingItempoolDeleteactivityAPIRequest struct {
	model.Params
	// 需要删除的活动的信息
	_param *CommonActivityParam
}

// NewAlibabaWdkMarketingItempoolDeleteactivityRequest 初始化AlibabaWdkMarketingItempoolDeleteactivityAPIRequest对象
func NewAlibabaWdkMarketingItempoolDeleteactivityRequest() *AlibabaWdkMarketingItempoolDeleteactivityAPIRequest {
	return &AlibabaWdkMarketingItempoolDeleteactivityAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingItempoolDeleteactivityAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolDeleteactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itempool.deleteactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolDeleteactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingItempoolDeleteactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 需要删除的活动的信息
func (r *AlibabaWdkMarketingItempoolDeleteactivityAPIRequest) SetParam(_param *CommonActivityParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkMarketingItempoolDeleteactivityAPIRequest) GetParam() *CommonActivityParam {
	return r._param
}

var poolAlibabaWdkMarketingItempoolDeleteactivityAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingItempoolDeleteactivityRequest()
	},
}

// GetAlibabaWdkMarketingItempoolDeleteactivityRequest 从 sync.Pool 获取 AlibabaWdkMarketingItempoolDeleteactivityAPIRequest
func GetAlibabaWdkMarketingItempoolDeleteactivityAPIRequest() *AlibabaWdkMarketingItempoolDeleteactivityAPIRequest {
	return poolAlibabaWdkMarketingItempoolDeleteactivityAPIRequest.Get().(*AlibabaWdkMarketingItempoolDeleteactivityAPIRequest)
}

// ReleaseAlibabaWdkMarketingItempoolDeleteactivityAPIRequest 将 AlibabaWdkMarketingItempoolDeleteactivityAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingItempoolDeleteactivityAPIRequest(v *AlibabaWdkMarketingItempoolDeleteactivityAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingItempoolDeleteactivityAPIRequest.Put(v)
}
