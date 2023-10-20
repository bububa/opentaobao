package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItemdiscountDeleteactivityAPIRequest 删除商品特价活动 API请求
// alibaba.hm.marketing.itemdiscount.deleteactivity
//
// 删除商品特价活动
type AlibabaHmMarketingItemdiscountDeleteactivityAPIRequest struct {
	model.Params
	// 需要删除的活动的信息
	_param *CommonActivityRequest
}

// NewAlibabaHmMarketingItemdiscountDeleteactivityRequest 初始化AlibabaHmMarketingItemdiscountDeleteactivityAPIRequest对象
func NewAlibabaHmMarketingItemdiscountDeleteactivityRequest() *AlibabaHmMarketingItemdiscountDeleteactivityAPIRequest {
	return &AlibabaHmMarketingItemdiscountDeleteactivityAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingItemdiscountDeleteactivityAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingItemdiscountDeleteactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itemdiscount.deleteactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingItemdiscountDeleteactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingItemdiscountDeleteactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 需要删除的活动的信息
func (r *AlibabaHmMarketingItemdiscountDeleteactivityAPIRequest) SetParam(_param *CommonActivityRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaHmMarketingItemdiscountDeleteactivityAPIRequest) GetParam() *CommonActivityRequest {
	return r._param
}

var poolAlibabaHmMarketingItemdiscountDeleteactivityAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingItemdiscountDeleteactivityRequest()
	},
}

// GetAlibabaHmMarketingItemdiscountDeleteactivityRequest 从 sync.Pool 获取 AlibabaHmMarketingItemdiscountDeleteactivityAPIRequest
func GetAlibabaHmMarketingItemdiscountDeleteactivityAPIRequest() *AlibabaHmMarketingItemdiscountDeleteactivityAPIRequest {
	return poolAlibabaHmMarketingItemdiscountDeleteactivityAPIRequest.Get().(*AlibabaHmMarketingItemdiscountDeleteactivityAPIRequest)
}

// ReleaseAlibabaHmMarketingItemdiscountDeleteactivityAPIRequest 将 AlibabaHmMarketingItemdiscountDeleteactivityAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingItemdiscountDeleteactivityAPIRequest(v *AlibabaHmMarketingItemdiscountDeleteactivityAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingItemdiscountDeleteactivityAPIRequest.Put(v)
}
