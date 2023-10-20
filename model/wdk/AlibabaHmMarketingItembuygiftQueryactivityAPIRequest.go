package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItembuygiftQueryactivityAPIRequest 查询买赠活动 API请求
// alibaba.hm.marketing.itembuygift.queryactivity
//
// 查询买赠活动
type AlibabaHmMarketingItembuygiftQueryactivityAPIRequest struct {
	model.Params
	// 查询入参
	_param *CommonActivityParam
}

// NewAlibabaHmMarketingItembuygiftQueryactivityRequest 初始化AlibabaHmMarketingItembuygiftQueryactivityAPIRequest对象
func NewAlibabaHmMarketingItembuygiftQueryactivityRequest() *AlibabaHmMarketingItembuygiftQueryactivityAPIRequest {
	return &AlibabaHmMarketingItembuygiftQueryactivityAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingItembuygiftQueryactivityAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingItembuygiftQueryactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itembuygift.queryactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingItembuygiftQueryactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingItembuygiftQueryactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 查询入参
func (r *AlibabaHmMarketingItembuygiftQueryactivityAPIRequest) SetParam(_param *CommonActivityParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaHmMarketingItembuygiftQueryactivityAPIRequest) GetParam() *CommonActivityParam {
	return r._param
}

var poolAlibabaHmMarketingItembuygiftQueryactivityAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingItembuygiftQueryactivityRequest()
	},
}

// GetAlibabaHmMarketingItembuygiftQueryactivityRequest 从 sync.Pool 获取 AlibabaHmMarketingItembuygiftQueryactivityAPIRequest
func GetAlibabaHmMarketingItembuygiftQueryactivityAPIRequest() *AlibabaHmMarketingItembuygiftQueryactivityAPIRequest {
	return poolAlibabaHmMarketingItembuygiftQueryactivityAPIRequest.Get().(*AlibabaHmMarketingItembuygiftQueryactivityAPIRequest)
}

// ReleaseAlibabaHmMarketingItembuygiftQueryactivityAPIRequest 将 AlibabaHmMarketingItembuygiftQueryactivityAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingItembuygiftQueryactivityAPIRequest(v *AlibabaHmMarketingItembuygiftQueryactivityAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingItembuygiftQueryactivityAPIRequest.Put(v)
}
