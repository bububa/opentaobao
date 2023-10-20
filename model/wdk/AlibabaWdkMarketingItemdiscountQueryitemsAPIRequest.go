package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItemdiscountQueryitemsAPIRequest 查询特价商品 API请求
// alibaba.wdk.marketing.itemdiscount.queryitems
//
// 查询参加特价活动的商品优惠详情
type AlibabaWdkMarketingItemdiscountQueryitemsAPIRequest struct {
	model.Params
	// 查询入参
	_param *ActivitySkuQuery
}

// NewAlibabaWdkMarketingItemdiscountQueryitemsRequest 初始化AlibabaWdkMarketingItemdiscountQueryitemsAPIRequest对象
func NewAlibabaWdkMarketingItemdiscountQueryitemsRequest() *AlibabaWdkMarketingItemdiscountQueryitemsAPIRequest {
	return &AlibabaWdkMarketingItemdiscountQueryitemsAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingItemdiscountQueryitemsAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItemdiscountQueryitemsAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itemdiscount.queryitems"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItemdiscountQueryitemsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingItemdiscountQueryitemsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 查询入参
func (r *AlibabaWdkMarketingItemdiscountQueryitemsAPIRequest) SetParam(_param *ActivitySkuQuery) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkMarketingItemdiscountQueryitemsAPIRequest) GetParam() *ActivitySkuQuery {
	return r._param
}

var poolAlibabaWdkMarketingItemdiscountQueryitemsAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingItemdiscountQueryitemsRequest()
	},
}

// GetAlibabaWdkMarketingItemdiscountQueryitemsRequest 从 sync.Pool 获取 AlibabaWdkMarketingItemdiscountQueryitemsAPIRequest
func GetAlibabaWdkMarketingItemdiscountQueryitemsAPIRequest() *AlibabaWdkMarketingItemdiscountQueryitemsAPIRequest {
	return poolAlibabaWdkMarketingItemdiscountQueryitemsAPIRequest.Get().(*AlibabaWdkMarketingItemdiscountQueryitemsAPIRequest)
}

// ReleaseAlibabaWdkMarketingItemdiscountQueryitemsAPIRequest 将 AlibabaWdkMarketingItemdiscountQueryitemsAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingItemdiscountQueryitemsAPIRequest(v *AlibabaWdkMarketingItemdiscountQueryitemsAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingItemdiscountQueryitemsAPIRequest.Put(v)
}
