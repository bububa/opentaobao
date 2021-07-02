package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest 添加单品买赠活动商品 API请求
// alibaba.retail.marketing.buygift.activity.sku.add
//
// 新增或更新单品买赠活动商品信息【同城零售】
type AlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest struct {
	model.Params
	// 添加活动商品参数
	_param *BuyGiftActivitySkuOperateRequest
}

// NewAlibabaRetailMarketingBuygiftActivitySkuAddRequest 初始化AlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest对象
func NewAlibabaRetailMarketingBuygiftActivitySkuAddRequest() *AlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest {
	return &AlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.buygift.activity.sku.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 添加活动商品参数
func (r *AlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest) SetParam(_param *BuyGiftActivitySkuOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r AlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest) GetParam() *BuyGiftActivitySkuOperateRequest {
	return r._param
}
