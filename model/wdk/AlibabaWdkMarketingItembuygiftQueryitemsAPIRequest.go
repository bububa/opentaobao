package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItembuygiftQueryitemsAPIRequest 查询买赠活动下的商品 API请求
// alibaba.wdk.marketing.itembuygift.queryitems
//
// 查询买赠活动下的商品
type AlibabaWdkMarketingItembuygiftQueryitemsAPIRequest struct {
	model.Params
	// 查询入参
	_param *ActivitySkuQuery
}

// NewAlibabaWdkMarketingItembuygiftQueryitemsRequest 初始化AlibabaWdkMarketingItembuygiftQueryitemsAPIRequest对象
func NewAlibabaWdkMarketingItembuygiftQueryitemsRequest() *AlibabaWdkMarketingItembuygiftQueryitemsAPIRequest {
	return &AlibabaWdkMarketingItembuygiftQueryitemsAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingItembuygiftQueryitemsAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItembuygiftQueryitemsAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itembuygift.queryitems"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItembuygiftQueryitemsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingItembuygiftQueryitemsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 查询入参
func (r *AlibabaWdkMarketingItembuygiftQueryitemsAPIRequest) SetParam(_param *ActivitySkuQuery) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkMarketingItembuygiftQueryitemsAPIRequest) GetParam() *ActivitySkuQuery {
	return r._param
}

var poolAlibabaWdkMarketingItembuygiftQueryitemsAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingItembuygiftQueryitemsRequest()
	},
}

// GetAlibabaWdkMarketingItembuygiftQueryitemsRequest 从 sync.Pool 获取 AlibabaWdkMarketingItembuygiftQueryitemsAPIRequest
func GetAlibabaWdkMarketingItembuygiftQueryitemsAPIRequest() *AlibabaWdkMarketingItembuygiftQueryitemsAPIRequest {
	return poolAlibabaWdkMarketingItembuygiftQueryitemsAPIRequest.Get().(*AlibabaWdkMarketingItembuygiftQueryitemsAPIRequest)
}

// ReleaseAlibabaWdkMarketingItembuygiftQueryitemsAPIRequest 将 AlibabaWdkMarketingItembuygiftQueryitemsAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingItembuygiftQueryitemsAPIRequest(v *AlibabaWdkMarketingItembuygiftQueryitemsAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingItembuygiftQueryitemsAPIRequest.Put(v)
}
