package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingFullrangeQueryactivityAPIRequest 全场活动查询活动 API请求
// alibaba.hm.marketing.fullrange.queryactivity
//
// 全场活动查询活动
type AlibabaHmMarketingFullrangeQueryactivityAPIRequest struct {
	model.Params
	// 查询活动入参
	_param0 *CommonActivityParam
}

// NewAlibabaHmMarketingFullrangeQueryactivityRequest 初始化AlibabaHmMarketingFullrangeQueryactivityAPIRequest对象
func NewAlibabaHmMarketingFullrangeQueryactivityRequest() *AlibabaHmMarketingFullrangeQueryactivityAPIRequest {
	return &AlibabaHmMarketingFullrangeQueryactivityAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingFullrangeQueryactivityAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingFullrangeQueryactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.fullrange.queryactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingFullrangeQueryactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingFullrangeQueryactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 查询活动入参
func (r *AlibabaHmMarketingFullrangeQueryactivityAPIRequest) SetParam0(_param0 *CommonActivityParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaHmMarketingFullrangeQueryactivityAPIRequest) GetParam0() *CommonActivityParam {
	return r._param0
}

var poolAlibabaHmMarketingFullrangeQueryactivityAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingFullrangeQueryactivityRequest()
	},
}

// GetAlibabaHmMarketingFullrangeQueryactivityRequest 从 sync.Pool 获取 AlibabaHmMarketingFullrangeQueryactivityAPIRequest
func GetAlibabaHmMarketingFullrangeQueryactivityAPIRequest() *AlibabaHmMarketingFullrangeQueryactivityAPIRequest {
	return poolAlibabaHmMarketingFullrangeQueryactivityAPIRequest.Get().(*AlibabaHmMarketingFullrangeQueryactivityAPIRequest)
}

// ReleaseAlibabaHmMarketingFullrangeQueryactivityAPIRequest 将 AlibabaHmMarketingFullrangeQueryactivityAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingFullrangeQueryactivityAPIRequest(v *AlibabaHmMarketingFullrangeQueryactivityAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingFullrangeQueryactivityAPIRequest.Put(v)
}
