package wdk

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itemdiscount.activity.sku.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
