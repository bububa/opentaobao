package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingBuygiftSkuQueryAPIRequest 查询买赠活动商品【同城零售】 API请求
// alibaba.retail.marketing.buygift.sku.query
//
// 查询买赠活动商品【同城零售】
type AlibabaRetailMarketingBuygiftSkuQueryAPIRequest struct {
	model.Params
	// 买赠商品查询入参
	_param0 *BuyGiftActivitySkuQueryRequest
}

// NewAlibabaRetailMarketingBuygiftSkuQueryRequest 初始化AlibabaRetailMarketingBuygiftSkuQueryAPIRequest对象
func NewAlibabaRetailMarketingBuygiftSkuQueryRequest() *AlibabaRetailMarketingBuygiftSkuQueryAPIRequest {
	return &AlibabaRetailMarketingBuygiftSkuQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaRetailMarketingBuygiftSkuQueryAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingBuygiftSkuQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.buygift.sku.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingBuygiftSkuQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailMarketingBuygiftSkuQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 买赠商品查询入参
func (r *AlibabaRetailMarketingBuygiftSkuQueryAPIRequest) SetParam0(_param0 *BuyGiftActivitySkuQueryRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaRetailMarketingBuygiftSkuQueryAPIRequest) GetParam0() *BuyGiftActivitySkuQueryRequest {
	return r._param0
}

var poolAlibabaRetailMarketingBuygiftSkuQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaRetailMarketingBuygiftSkuQueryRequest()
	},
}

// GetAlibabaRetailMarketingBuygiftSkuQueryRequest 从 sync.Pool 获取 AlibabaRetailMarketingBuygiftSkuQueryAPIRequest
func GetAlibabaRetailMarketingBuygiftSkuQueryAPIRequest() *AlibabaRetailMarketingBuygiftSkuQueryAPIRequest {
	return poolAlibabaRetailMarketingBuygiftSkuQueryAPIRequest.Get().(*AlibabaRetailMarketingBuygiftSkuQueryAPIRequest)
}

// ReleaseAlibabaRetailMarketingBuygiftSkuQueryAPIRequest 将 AlibabaRetailMarketingBuygiftSkuQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaRetailMarketingBuygiftSkuQueryAPIRequest(v *AlibabaRetailMarketingBuygiftSkuQueryAPIRequest) {
	v.Reset()
	poolAlibabaRetailMarketingBuygiftSkuQueryAPIRequest.Put(v)
}
