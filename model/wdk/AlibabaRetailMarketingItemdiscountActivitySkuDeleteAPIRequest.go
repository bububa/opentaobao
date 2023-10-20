package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIRequest 删除特价活动商品 API请求
// alibaba.retail.marketing.itemdiscount.activity.sku.delete
//
// 删除活动商品信息【同城零售】
type AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIRequest struct {
	model.Params
	// 添加活动商品参数
	_param *ItemDiscountActivityElementOperateRequest
}

// NewAlibabaRetailMarketingItemdiscountActivitySkuDeleteRequest 初始化AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIRequest对象
func NewAlibabaRetailMarketingItemdiscountActivitySkuDeleteRequest() *AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIRequest {
	return &AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itemdiscount.activity.sku.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 添加活动商品参数
func (r *AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIRequest) SetParam(_param *ItemDiscountActivityElementOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIRequest) GetParam() *ItemDiscountActivityElementOperateRequest {
	return r._param
}

var poolAlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaRetailMarketingItemdiscountActivitySkuDeleteRequest()
	},
}

// GetAlibabaRetailMarketingItemdiscountActivitySkuDeleteRequest 从 sync.Pool 获取 AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIRequest
func GetAlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIRequest() *AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIRequest {
	return poolAlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIRequest.Get().(*AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIRequest)
}

// ReleaseAlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIRequest 将 AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIRequest 放入 sync.Pool
func ReleaseAlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIRequest(v *AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIRequest) {
	v.Reset()
	poolAlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIRequest.Put(v)
}
