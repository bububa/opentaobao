package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkBmPaiyangStockQueryAPIRequest 派样商品门店库存查询接口 API请求
// alibaba.wdk.bm.paiyang.stock.query
//
// 淘鲜达接入第三方进行派样，第三方查询派样商品的门店库存信息。
type AlibabaWdkBmPaiyangStockQueryAPIRequest struct {
	model.Params
	// 请求入参
	_isvShopStockParam *IsvShopStockParam
}

// NewAlibabaWdkBmPaiyangStockQueryRequest 初始化AlibabaWdkBmPaiyangStockQueryAPIRequest对象
func NewAlibabaWdkBmPaiyangStockQueryRequest() *AlibabaWdkBmPaiyangStockQueryAPIRequest {
	return &AlibabaWdkBmPaiyangStockQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkBmPaiyangStockQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.bm.paiyang.stock.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkBmPaiyangStockQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is IsvShopStockParam Setter
// 请求入参
func (r *AlibabaWdkBmPaiyangStockQueryAPIRequest) SetIsvShopStockParam(_isvShopStockParam *IsvShopStockParam) error {
	r._isvShopStockParam = _isvShopStockParam
	r.Set("isv_shop_stock_param", _isvShopStockParam)
	return nil
}

// Get IsvShopStockParam Getter
func (r AlibabaWdkBmPaiyangStockQueryAPIRequest) GetIsvShopStockParam() *IsvShopStockParam {
	return r._isvShopStockParam
}
