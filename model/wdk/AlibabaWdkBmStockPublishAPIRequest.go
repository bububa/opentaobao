package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkBmStockPublishAPIRequest 品牌营销涉及到的商品的库存同步接口 API请求
// alibaba.wdk.bm.stock.publish
//
// 用于操作sku的库存
type AlibabaWdkBmStockPublishAPIRequest struct {
	model.Params
	// 批量入参
	_skuStockPublishParamList []SkuStockPublishParamDo
}

// NewAlibabaWdkBmStockPublishRequest 初始化AlibabaWdkBmStockPublishAPIRequest对象
func NewAlibabaWdkBmStockPublishRequest() *AlibabaWdkBmStockPublishAPIRequest {
	return &AlibabaWdkBmStockPublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkBmStockPublishAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.bm.stock.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkBmStockPublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkBmStockPublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuStockPublishParamList is SkuStockPublishParamList Setter
// 批量入参
func (r *AlibabaWdkBmStockPublishAPIRequest) SetSkuStockPublishParamList(_skuStockPublishParamList []SkuStockPublishParamDo) error {
	r._skuStockPublishParamList = _skuStockPublishParamList
	r.Set("sku_stock_publish_param_list", _skuStockPublishParamList)
	return nil
}

// GetSkuStockPublishParamList SkuStockPublishParamList Getter
func (r AlibabaWdkBmStockPublishAPIRequest) GetSkuStockPublishParamList() []SkuStockPublishParamDo {
	return r._skuStockPublishParamList
}
