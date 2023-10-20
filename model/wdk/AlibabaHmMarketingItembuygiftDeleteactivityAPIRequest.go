package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItembuygiftDeleteactivityAPIRequest 删除买赠活动 API请求
// alibaba.hm.marketing.itembuygift.deleteactivity
//
// 删除买赠活动
type AlibabaHmMarketingItembuygiftDeleteactivityAPIRequest struct {
	model.Params
	// 要删除的活动信息
	_param *CommonActivityParam
}

// NewAlibabaHmMarketingItembuygiftDeleteactivityRequest 初始化AlibabaHmMarketingItembuygiftDeleteactivityAPIRequest对象
func NewAlibabaHmMarketingItembuygiftDeleteactivityRequest() *AlibabaHmMarketingItembuygiftDeleteactivityAPIRequest {
	return &AlibabaHmMarketingItembuygiftDeleteactivityAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingItembuygiftDeleteactivityAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingItembuygiftDeleteactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itembuygift.deleteactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingItembuygiftDeleteactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingItembuygiftDeleteactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 要删除的活动信息
func (r *AlibabaHmMarketingItembuygiftDeleteactivityAPIRequest) SetParam(_param *CommonActivityParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaHmMarketingItembuygiftDeleteactivityAPIRequest) GetParam() *CommonActivityParam {
	return r._param
}

var poolAlibabaHmMarketingItembuygiftDeleteactivityAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingItembuygiftDeleteactivityRequest()
	},
}

// GetAlibabaHmMarketingItembuygiftDeleteactivityRequest 从 sync.Pool 获取 AlibabaHmMarketingItembuygiftDeleteactivityAPIRequest
func GetAlibabaHmMarketingItembuygiftDeleteactivityAPIRequest() *AlibabaHmMarketingItembuygiftDeleteactivityAPIRequest {
	return poolAlibabaHmMarketingItembuygiftDeleteactivityAPIRequest.Get().(*AlibabaHmMarketingItembuygiftDeleteactivityAPIRequest)
}

// ReleaseAlibabaHmMarketingItembuygiftDeleteactivityAPIRequest 将 AlibabaHmMarketingItembuygiftDeleteactivityAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingItembuygiftDeleteactivityAPIRequest(v *AlibabaHmMarketingItembuygiftDeleteactivityAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingItembuygiftDeleteactivityAPIRequest.Put(v)
}
