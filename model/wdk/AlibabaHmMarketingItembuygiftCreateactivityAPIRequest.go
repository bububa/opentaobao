package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItembuygiftCreateactivityAPIRequest 创建买赠活动 API请求
// alibaba.hm.marketing.itembuygift.createactivity
//
// 创建买赠活动
type AlibabaHmMarketingItembuygiftCreateactivityAPIRequest struct {
	model.Params
	// 创建活动请求入参
	_param *ItemBuyGiftActivity
}

// NewAlibabaHmMarketingItembuygiftCreateactivityRequest 初始化AlibabaHmMarketingItembuygiftCreateactivityAPIRequest对象
func NewAlibabaHmMarketingItembuygiftCreateactivityRequest() *AlibabaHmMarketingItembuygiftCreateactivityAPIRequest {
	return &AlibabaHmMarketingItembuygiftCreateactivityAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingItembuygiftCreateactivityAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingItembuygiftCreateactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itembuygift.createactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingItembuygiftCreateactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingItembuygiftCreateactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建活动请求入参
func (r *AlibabaHmMarketingItembuygiftCreateactivityAPIRequest) SetParam(_param *ItemBuyGiftActivity) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaHmMarketingItembuygiftCreateactivityAPIRequest) GetParam() *ItemBuyGiftActivity {
	return r._param
}

var poolAlibabaHmMarketingItembuygiftCreateactivityAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingItembuygiftCreateactivityRequest()
	},
}

// GetAlibabaHmMarketingItembuygiftCreateactivityRequest 从 sync.Pool 获取 AlibabaHmMarketingItembuygiftCreateactivityAPIRequest
func GetAlibabaHmMarketingItembuygiftCreateactivityAPIRequest() *AlibabaHmMarketingItembuygiftCreateactivityAPIRequest {
	return poolAlibabaHmMarketingItembuygiftCreateactivityAPIRequest.Get().(*AlibabaHmMarketingItembuygiftCreateactivityAPIRequest)
}

// ReleaseAlibabaHmMarketingItembuygiftCreateactivityAPIRequest 将 AlibabaHmMarketingItembuygiftCreateactivityAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingItembuygiftCreateactivityAPIRequest(v *AlibabaHmMarketingItembuygiftCreateactivityAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingItembuygiftCreateactivityAPIRequest.Put(v)
}
