package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItempoolDeleteactivityAPIRequest 删除商品池活动 API请求
// alibaba.hm.marketing.itempool.deleteactivity
//
// 删除商品池活动
type AlibabaHmMarketingItempoolDeleteactivityAPIRequest struct {
	model.Params
	// 需要删除的活动的信息
	_param *CommonActivityParam
}

// NewAlibabaHmMarketingItempoolDeleteactivityRequest 初始化AlibabaHmMarketingItempoolDeleteactivityAPIRequest对象
func NewAlibabaHmMarketingItempoolDeleteactivityRequest() *AlibabaHmMarketingItempoolDeleteactivityAPIRequest {
	return &AlibabaHmMarketingItempoolDeleteactivityAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingItempoolDeleteactivityAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingItempoolDeleteactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itempool.deleteactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingItempoolDeleteactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingItempoolDeleteactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 需要删除的活动的信息
func (r *AlibabaHmMarketingItempoolDeleteactivityAPIRequest) SetParam(_param *CommonActivityParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaHmMarketingItempoolDeleteactivityAPIRequest) GetParam() *CommonActivityParam {
	return r._param
}

var poolAlibabaHmMarketingItempoolDeleteactivityAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingItempoolDeleteactivityRequest()
	},
}

// GetAlibabaHmMarketingItempoolDeleteactivityRequest 从 sync.Pool 获取 AlibabaHmMarketingItempoolDeleteactivityAPIRequest
func GetAlibabaHmMarketingItempoolDeleteactivityAPIRequest() *AlibabaHmMarketingItempoolDeleteactivityAPIRequest {
	return poolAlibabaHmMarketingItempoolDeleteactivityAPIRequest.Get().(*AlibabaHmMarketingItempoolDeleteactivityAPIRequest)
}

// ReleaseAlibabaHmMarketingItempoolDeleteactivityAPIRequest 将 AlibabaHmMarketingItempoolDeleteactivityAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingItempoolDeleteactivityAPIRequest(v *AlibabaHmMarketingItempoolDeleteactivityAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingItempoolDeleteactivityAPIRequest.Put(v)
}
