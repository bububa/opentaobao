package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItempoolCreateactivityAPIRequest 添加商品池活动 API请求
// alibaba.hm.marketing.itempool.createactivity
//
// 添加商品池活动
type AlibabaHmMarketingItempoolCreateactivityAPIRequest struct {
	model.Params
	// 创建活动请求入参
	_param *ItemPoolActivity
}

// NewAlibabaHmMarketingItempoolCreateactivityRequest 初始化AlibabaHmMarketingItempoolCreateactivityAPIRequest对象
func NewAlibabaHmMarketingItempoolCreateactivityRequest() *AlibabaHmMarketingItempoolCreateactivityAPIRequest {
	return &AlibabaHmMarketingItempoolCreateactivityAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingItempoolCreateactivityAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingItempoolCreateactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itempool.createactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingItempoolCreateactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingItempoolCreateactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建活动请求入参
func (r *AlibabaHmMarketingItempoolCreateactivityAPIRequest) SetParam(_param *ItemPoolActivity) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaHmMarketingItempoolCreateactivityAPIRequest) GetParam() *ItemPoolActivity {
	return r._param
}

var poolAlibabaHmMarketingItempoolCreateactivityAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingItempoolCreateactivityRequest()
	},
}

// GetAlibabaHmMarketingItempoolCreateactivityRequest 从 sync.Pool 获取 AlibabaHmMarketingItempoolCreateactivityAPIRequest
func GetAlibabaHmMarketingItempoolCreateactivityAPIRequest() *AlibabaHmMarketingItempoolCreateactivityAPIRequest {
	return poolAlibabaHmMarketingItempoolCreateactivityAPIRequest.Get().(*AlibabaHmMarketingItempoolCreateactivityAPIRequest)
}

// ReleaseAlibabaHmMarketingItempoolCreateactivityAPIRequest 将 AlibabaHmMarketingItempoolCreateactivityAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingItempoolCreateactivityAPIRequest(v *AlibabaHmMarketingItempoolCreateactivityAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingItempoolCreateactivityAPIRequest.Put(v)
}
