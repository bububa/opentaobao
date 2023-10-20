package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItempoolQueryitemsAPIRequest 查询商品池活动下的商品 API请求
// alibaba.wdk.marketing.itempool.queryitems
//
// 查询商品池活动下面的商品
type AlibabaWdkMarketingItempoolQueryitemsAPIRequest struct {
	model.Params
	// 查询入参
	_param *ActivitySkuQuery
}

// NewAlibabaWdkMarketingItempoolQueryitemsRequest 初始化AlibabaWdkMarketingItempoolQueryitemsAPIRequest对象
func NewAlibabaWdkMarketingItempoolQueryitemsRequest() *AlibabaWdkMarketingItempoolQueryitemsAPIRequest {
	return &AlibabaWdkMarketingItempoolQueryitemsAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingItempoolQueryitemsAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolQueryitemsAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itempool.queryitems"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolQueryitemsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingItempoolQueryitemsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 查询入参
func (r *AlibabaWdkMarketingItempoolQueryitemsAPIRequest) SetParam(_param *ActivitySkuQuery) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkMarketingItempoolQueryitemsAPIRequest) GetParam() *ActivitySkuQuery {
	return r._param
}

var poolAlibabaWdkMarketingItempoolQueryitemsAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingItempoolQueryitemsRequest()
	},
}

// GetAlibabaWdkMarketingItempoolQueryitemsRequest 从 sync.Pool 获取 AlibabaWdkMarketingItempoolQueryitemsAPIRequest
func GetAlibabaWdkMarketingItempoolQueryitemsAPIRequest() *AlibabaWdkMarketingItempoolQueryitemsAPIRequest {
	return poolAlibabaWdkMarketingItempoolQueryitemsAPIRequest.Get().(*AlibabaWdkMarketingItempoolQueryitemsAPIRequest)
}

// ReleaseAlibabaWdkMarketingItempoolQueryitemsAPIRequest 将 AlibabaWdkMarketingItempoolQueryitemsAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingItempoolQueryitemsAPIRequest(v *AlibabaWdkMarketingItempoolQueryitemsAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingItempoolQueryitemsAPIRequest.Put(v)
}
