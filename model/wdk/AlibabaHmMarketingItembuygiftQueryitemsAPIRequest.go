package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItembuygiftQueryitemsAPIRequest 查询买赠活动下的商品 API请求
// alibaba.hm.marketing.itembuygift.queryitems
//
// 查询买赠活动下的商品
type AlibabaHmMarketingItembuygiftQueryitemsAPIRequest struct {
	model.Params
	// 查询入参
	_param *ActivitySkuQuery
}

// NewAlibabaHmMarketingItembuygiftQueryitemsRequest 初始化AlibabaHmMarketingItembuygiftQueryitemsAPIRequest对象
func NewAlibabaHmMarketingItembuygiftQueryitemsRequest() *AlibabaHmMarketingItembuygiftQueryitemsAPIRequest {
	return &AlibabaHmMarketingItembuygiftQueryitemsAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingItembuygiftQueryitemsAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingItembuygiftQueryitemsAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itembuygift.queryitems"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingItembuygiftQueryitemsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingItembuygiftQueryitemsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 查询入参
func (r *AlibabaHmMarketingItembuygiftQueryitemsAPIRequest) SetParam(_param *ActivitySkuQuery) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaHmMarketingItembuygiftQueryitemsAPIRequest) GetParam() *ActivitySkuQuery {
	return r._param
}

var poolAlibabaHmMarketingItembuygiftQueryitemsAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingItembuygiftQueryitemsRequest()
	},
}

// GetAlibabaHmMarketingItembuygiftQueryitemsRequest 从 sync.Pool 获取 AlibabaHmMarketingItembuygiftQueryitemsAPIRequest
func GetAlibabaHmMarketingItembuygiftQueryitemsAPIRequest() *AlibabaHmMarketingItembuygiftQueryitemsAPIRequest {
	return poolAlibabaHmMarketingItembuygiftQueryitemsAPIRequest.Get().(*AlibabaHmMarketingItembuygiftQueryitemsAPIRequest)
}

// ReleaseAlibabaHmMarketingItembuygiftQueryitemsAPIRequest 将 AlibabaHmMarketingItembuygiftQueryitemsAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingItembuygiftQueryitemsAPIRequest(v *AlibabaHmMarketingItembuygiftQueryitemsAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingItembuygiftQueryitemsAPIRequest.Put(v)
}
