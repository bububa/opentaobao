package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItemdiscountQueryitemsAPIRequest 查询特价商品 API请求
// alibaba.hm.marketing.itemdiscount.queryitems
//
// 查询参加特价活动的商品优惠详情
type AlibabaHmMarketingItemdiscountQueryitemsAPIRequest struct {
	model.Params
	// 查询入参
	_param *ActivitySkuQuery
}

// NewAlibabaHmMarketingItemdiscountQueryitemsRequest 初始化AlibabaHmMarketingItemdiscountQueryitemsAPIRequest对象
func NewAlibabaHmMarketingItemdiscountQueryitemsRequest() *AlibabaHmMarketingItemdiscountQueryitemsAPIRequest {
	return &AlibabaHmMarketingItemdiscountQueryitemsAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingItemdiscountQueryitemsAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingItemdiscountQueryitemsAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itemdiscount.queryitems"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingItemdiscountQueryitemsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingItemdiscountQueryitemsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 查询入参
func (r *AlibabaHmMarketingItemdiscountQueryitemsAPIRequest) SetParam(_param *ActivitySkuQuery) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaHmMarketingItemdiscountQueryitemsAPIRequest) GetParam() *ActivitySkuQuery {
	return r._param
}

var poolAlibabaHmMarketingItemdiscountQueryitemsAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingItemdiscountQueryitemsRequest()
	},
}

// GetAlibabaHmMarketingItemdiscountQueryitemsRequest 从 sync.Pool 获取 AlibabaHmMarketingItemdiscountQueryitemsAPIRequest
func GetAlibabaHmMarketingItemdiscountQueryitemsAPIRequest() *AlibabaHmMarketingItemdiscountQueryitemsAPIRequest {
	return poolAlibabaHmMarketingItemdiscountQueryitemsAPIRequest.Get().(*AlibabaHmMarketingItemdiscountQueryitemsAPIRequest)
}

// ReleaseAlibabaHmMarketingItemdiscountQueryitemsAPIRequest 将 AlibabaHmMarketingItemdiscountQueryitemsAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingItemdiscountQueryitemsAPIRequest(v *AlibabaHmMarketingItemdiscountQueryitemsAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingItemdiscountQueryitemsAPIRequest.Put(v)
}
