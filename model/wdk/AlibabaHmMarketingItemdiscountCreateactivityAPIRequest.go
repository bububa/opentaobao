package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItemdiscountCreateactivityAPIRequest 创建商品特价活动 API请求
// alibaba.hm.marketing.itemdiscount.createactivity
//
// 创建商品特价活动
type AlibabaHmMarketingItemdiscountCreateactivityAPIRequest struct {
	model.Params
	// 创建活动请求入参
	_param *ItemDiscountActivityRequest
}

// NewAlibabaHmMarketingItemdiscountCreateactivityRequest 初始化AlibabaHmMarketingItemdiscountCreateactivityAPIRequest对象
func NewAlibabaHmMarketingItemdiscountCreateactivityRequest() *AlibabaHmMarketingItemdiscountCreateactivityAPIRequest {
	return &AlibabaHmMarketingItemdiscountCreateactivityAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingItemdiscountCreateactivityAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingItemdiscountCreateactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itemdiscount.createactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingItemdiscountCreateactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingItemdiscountCreateactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建活动请求入参
func (r *AlibabaHmMarketingItemdiscountCreateactivityAPIRequest) SetParam(_param *ItemDiscountActivityRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaHmMarketingItemdiscountCreateactivityAPIRequest) GetParam() *ItemDiscountActivityRequest {
	return r._param
}

var poolAlibabaHmMarketingItemdiscountCreateactivityAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingItemdiscountCreateactivityRequest()
	},
}

// GetAlibabaHmMarketingItemdiscountCreateactivityRequest 从 sync.Pool 获取 AlibabaHmMarketingItemdiscountCreateactivityAPIRequest
func GetAlibabaHmMarketingItemdiscountCreateactivityAPIRequest() *AlibabaHmMarketingItemdiscountCreateactivityAPIRequest {
	return poolAlibabaHmMarketingItemdiscountCreateactivityAPIRequest.Get().(*AlibabaHmMarketingItemdiscountCreateactivityAPIRequest)
}

// ReleaseAlibabaHmMarketingItemdiscountCreateactivityAPIRequest 将 AlibabaHmMarketingItemdiscountCreateactivityAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingItemdiscountCreateactivityAPIRequest(v *AlibabaHmMarketingItemdiscountCreateactivityAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingItemdiscountCreateactivityAPIRequest.Put(v)
}
