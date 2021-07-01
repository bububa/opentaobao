package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest
特价活动新增商品 API请求
alibaba.retail.marketing.itemdiscount.activity.sku.add

新增或更新活动商品信息【同城零售】 */
type AlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest struct {
	model.Params
	// 添加活动商品参数
	_param *ItemDiscountActivityElementOperateRequest
}

// NewAlibabaRetailMarketingItemdiscountActivitySkuAddRequest 初始化AlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest对象
func NewAlibabaRetailMarketingItemdiscountActivitySkuAddRequest() *AlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest {
	return &AlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itemdiscount.activity.sku.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 添加活动商品参数
func (r *AlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest) SetParam(_param *ItemDiscountActivityElementOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r AlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest) GetParam() *ItemDiscountActivityElementOperateRequest {
	return r._param
}
