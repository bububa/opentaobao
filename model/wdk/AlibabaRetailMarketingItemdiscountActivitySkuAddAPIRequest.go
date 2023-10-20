package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest 特价活动新增商品 API请求
// alibaba.retail.marketing.itemdiscount.activity.sku.add
//
// 新增或更新活动商品信息【同城零售】
type AlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest struct {
	model.Params
	// 添加活动商品参数
	_param *ItemDiscountActivityElementOperateRequest
}

// NewAlibabaRetailMarketingItemdiscountActivitySkuAddRequest 初始化AlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest对象
func NewAlibabaRetailMarketingItemdiscountActivitySkuAddRequest() *AlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest {
	return &AlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itemdiscount.activity.sku.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 添加活动商品参数
func (r *AlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest) SetParam(_param *ItemDiscountActivityElementOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest) GetParam() *ItemDiscountActivityElementOperateRequest {
	return r._param
}

var poolAlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaRetailMarketingItemdiscountActivitySkuAddRequest()
	},
}

// GetAlibabaRetailMarketingItemdiscountActivitySkuAddRequest 从 sync.Pool 获取 AlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest
func GetAlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest() *AlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest {
	return poolAlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest.Get().(*AlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest)
}

// ReleaseAlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest 将 AlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest 放入 sync.Pool
func ReleaseAlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest(v *AlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest) {
	v.Reset()
	poolAlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest.Put(v)
}
