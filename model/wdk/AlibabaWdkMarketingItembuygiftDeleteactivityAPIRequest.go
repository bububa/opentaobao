package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest 删除买赠活动 API请求
// alibaba.wdk.marketing.itembuygift.deleteactivity
//
// 删除买赠活动
type AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest struct {
	model.Params
	// 要删除的活动信息
	_param *CommonActivityParam
}

// NewAlibabaWdkMarketingItembuygiftDeleteactivityRequest 初始化AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest对象
func NewAlibabaWdkMarketingItembuygiftDeleteactivityRequest() *AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest {
	return &AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itembuygift.deleteactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 要删除的活动信息
func (r *AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest) SetParam(_param *CommonActivityParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest) GetParam() *CommonActivityParam {
	return r._param
}

var poolAlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingItembuygiftDeleteactivityRequest()
	},
}

// GetAlibabaWdkMarketingItembuygiftDeleteactivityRequest 从 sync.Pool 获取 AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest
func GetAlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest() *AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest {
	return poolAlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest.Get().(*AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest)
}

// ReleaseAlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest 将 AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest(v *AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest.Put(v)
}
