package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest 删除单品买赠活动商品 API请求
// alibaba.retail.marketing.buygift.activity.sku.delete
//
// 删除单品买赠活动商品信息【同城零售】
type AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest struct {
	model.Params
	// 删除买赠活动商品参数
	_param *BuyGiftActivitySkuOperateRequest
}

// NewAlibabaRetailMarketingBuygiftActivitySkuDeleteRequest 初始化AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest对象
func NewAlibabaRetailMarketingBuygiftActivitySkuDeleteRequest() *AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest {
	return &AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.buygift.activity.sku.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 删除买赠活动商品参数
func (r *AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest) SetParam(_param *BuyGiftActivitySkuOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest) GetParam() *BuyGiftActivitySkuOperateRequest {
	return r._param
}
