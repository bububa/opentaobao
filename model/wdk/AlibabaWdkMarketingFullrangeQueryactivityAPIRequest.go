package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingFullrangeQueryactivityAPIRequest 全场活动查询活动 API请求
// alibaba.wdk.marketing.fullrange.queryactivity
//
// 全场活动查询活动
type AlibabaWdkMarketingFullrangeQueryactivityAPIRequest struct {
	model.Params
	// 查询活动入参
	_param0 *CommonActivityParam
}

// NewAlibabaWdkMarketingFullrangeQueryactivityRequest 初始化AlibabaWdkMarketingFullrangeQueryactivityAPIRequest对象
func NewAlibabaWdkMarketingFullrangeQueryactivityRequest() *AlibabaWdkMarketingFullrangeQueryactivityAPIRequest {
	return &AlibabaWdkMarketingFullrangeQueryactivityAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingFullrangeQueryactivityAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingFullrangeQueryactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.fullrange.queryactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingFullrangeQueryactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingFullrangeQueryactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 查询活动入参
func (r *AlibabaWdkMarketingFullrangeQueryactivityAPIRequest) SetParam0(_param0 *CommonActivityParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaWdkMarketingFullrangeQueryactivityAPIRequest) GetParam0() *CommonActivityParam {
	return r._param0
}

var poolAlibabaWdkMarketingFullrangeQueryactivityAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingFullrangeQueryactivityRequest()
	},
}

// GetAlibabaWdkMarketingFullrangeQueryactivityRequest 从 sync.Pool 获取 AlibabaWdkMarketingFullrangeQueryactivityAPIRequest
func GetAlibabaWdkMarketingFullrangeQueryactivityAPIRequest() *AlibabaWdkMarketingFullrangeQueryactivityAPIRequest {
	return poolAlibabaWdkMarketingFullrangeQueryactivityAPIRequest.Get().(*AlibabaWdkMarketingFullrangeQueryactivityAPIRequest)
}

// ReleaseAlibabaWdkMarketingFullrangeQueryactivityAPIRequest 将 AlibabaWdkMarketingFullrangeQueryactivityAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingFullrangeQueryactivityAPIRequest(v *AlibabaWdkMarketingFullrangeQueryactivityAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingFullrangeQueryactivityAPIRequest.Put(v)
}
