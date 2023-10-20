package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItempoolSkuQueryAPIRequest 查询商品池活动商品【同城零售】 API请求
// alibaba.retail.marketing.itempool.sku.query
//
// 查询商品池活动商品【同城零售】
type AlibabaRetailMarketingItempoolSkuQueryAPIRequest struct {
	model.Params
	// 请求入参
	_param0 *ItemPoolActivitySkuQueryRequest
}

// NewAlibabaRetailMarketingItempoolSkuQueryRequest 初始化AlibabaRetailMarketingItempoolSkuQueryAPIRequest对象
func NewAlibabaRetailMarketingItempoolSkuQueryRequest() *AlibabaRetailMarketingItempoolSkuQueryAPIRequest {
	return &AlibabaRetailMarketingItempoolSkuQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaRetailMarketingItempoolSkuQueryAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItempoolSkuQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itempool.sku.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItempoolSkuQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailMarketingItempoolSkuQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 请求入参
func (r *AlibabaRetailMarketingItempoolSkuQueryAPIRequest) SetParam0(_param0 *ItemPoolActivitySkuQueryRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaRetailMarketingItempoolSkuQueryAPIRequest) GetParam0() *ItemPoolActivitySkuQueryRequest {
	return r._param0
}

var poolAlibabaRetailMarketingItempoolSkuQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaRetailMarketingItempoolSkuQueryRequest()
	},
}

// GetAlibabaRetailMarketingItempoolSkuQueryRequest 从 sync.Pool 获取 AlibabaRetailMarketingItempoolSkuQueryAPIRequest
func GetAlibabaRetailMarketingItempoolSkuQueryAPIRequest() *AlibabaRetailMarketingItempoolSkuQueryAPIRequest {
	return poolAlibabaRetailMarketingItempoolSkuQueryAPIRequest.Get().(*AlibabaRetailMarketingItempoolSkuQueryAPIRequest)
}

// ReleaseAlibabaRetailMarketingItempoolSkuQueryAPIRequest 将 AlibabaRetailMarketingItempoolSkuQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaRetailMarketingItempoolSkuQueryAPIRequest(v *AlibabaRetailMarketingItempoolSkuQueryAPIRequest) {
	v.Reset()
	poolAlibabaRetailMarketingItempoolSkuQueryAPIRequest.Put(v)
}
