package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingFullrangeQueryitemAPIRequest 全场活动查询换购品 API请求
// alibaba.hm.marketing.fullrange.queryitem
//
// 全场活动查询换购品
type AlibabaHmMarketingFullrangeQueryitemAPIRequest struct {
	model.Params
	// 换购商品查询参数
	_param0 *ActivitySkuQuery
}

// NewAlibabaHmMarketingFullrangeQueryitemRequest 初始化AlibabaHmMarketingFullrangeQueryitemAPIRequest对象
func NewAlibabaHmMarketingFullrangeQueryitemRequest() *AlibabaHmMarketingFullrangeQueryitemAPIRequest {
	return &AlibabaHmMarketingFullrangeQueryitemAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingFullrangeQueryitemAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingFullrangeQueryitemAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.fullrange.queryitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingFullrangeQueryitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingFullrangeQueryitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 换购商品查询参数
func (r *AlibabaHmMarketingFullrangeQueryitemAPIRequest) SetParam0(_param0 *ActivitySkuQuery) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaHmMarketingFullrangeQueryitemAPIRequest) GetParam0() *ActivitySkuQuery {
	return r._param0
}

var poolAlibabaHmMarketingFullrangeQueryitemAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingFullrangeQueryitemRequest()
	},
}

// GetAlibabaHmMarketingFullrangeQueryitemRequest 从 sync.Pool 获取 AlibabaHmMarketingFullrangeQueryitemAPIRequest
func GetAlibabaHmMarketingFullrangeQueryitemAPIRequest() *AlibabaHmMarketingFullrangeQueryitemAPIRequest {
	return poolAlibabaHmMarketingFullrangeQueryitemAPIRequest.Get().(*AlibabaHmMarketingFullrangeQueryitemAPIRequest)
}

// ReleaseAlibabaHmMarketingFullrangeQueryitemAPIRequest 将 AlibabaHmMarketingFullrangeQueryitemAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingFullrangeQueryitemAPIRequest(v *AlibabaHmMarketingFullrangeQueryitemAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingFullrangeQueryitemAPIRequest.Put(v)
}
