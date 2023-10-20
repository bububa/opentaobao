package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingFullrangeQueryitemAPIRequest 全场活动查询换购品 API请求
// alibaba.wdk.marketing.fullrange.queryitem
//
// 全场活动查询换购品
type AlibabaWdkMarketingFullrangeQueryitemAPIRequest struct {
	model.Params
	// 换购商品查询参数
	_param0 *ActivitySkuQuery
}

// NewAlibabaWdkMarketingFullrangeQueryitemRequest 初始化AlibabaWdkMarketingFullrangeQueryitemAPIRequest对象
func NewAlibabaWdkMarketingFullrangeQueryitemRequest() *AlibabaWdkMarketingFullrangeQueryitemAPIRequest {
	return &AlibabaWdkMarketingFullrangeQueryitemAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingFullrangeQueryitemAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingFullrangeQueryitemAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.fullrange.queryitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingFullrangeQueryitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingFullrangeQueryitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 换购商品查询参数
func (r *AlibabaWdkMarketingFullrangeQueryitemAPIRequest) SetParam0(_param0 *ActivitySkuQuery) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaWdkMarketingFullrangeQueryitemAPIRequest) GetParam0() *ActivitySkuQuery {
	return r._param0
}

var poolAlibabaWdkMarketingFullrangeQueryitemAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingFullrangeQueryitemRequest()
	},
}

// GetAlibabaWdkMarketingFullrangeQueryitemRequest 从 sync.Pool 获取 AlibabaWdkMarketingFullrangeQueryitemAPIRequest
func GetAlibabaWdkMarketingFullrangeQueryitemAPIRequest() *AlibabaWdkMarketingFullrangeQueryitemAPIRequest {
	return poolAlibabaWdkMarketingFullrangeQueryitemAPIRequest.Get().(*AlibabaWdkMarketingFullrangeQueryitemAPIRequest)
}

// ReleaseAlibabaWdkMarketingFullrangeQueryitemAPIRequest 将 AlibabaWdkMarketingFullrangeQueryitemAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingFullrangeQueryitemAPIRequest(v *AlibabaWdkMarketingFullrangeQueryitemAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingFullrangeQueryitemAPIRequest.Put(v)
}
