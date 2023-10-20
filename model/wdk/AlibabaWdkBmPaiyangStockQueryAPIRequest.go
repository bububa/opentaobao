package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkbmpaiyangstockqueryAPIRequest 派样商品门店库存查询接口 API请求
// alibaba.wdk.bm.paiyang.stock.query
//
// 淘鲜达接入第三方进行派样，第三方查询派样商品的门店库存信息。
type AlibabawdkbmpaiyangstockqueryAPIRequest struct {
	model.Params
	// 请求入参
	_isvShopStockParam *IsvShopStockParam
}

// NewAlibabawdkbmpaiyangstockqueryRequest 初始化AlibabawdkbmpaiyangstockqueryAPIRequest对象
func NewAlibabawdkbmpaiyangstockqueryRequest() *AlibabawdkbmpaiyangstockqueryAPIRequest {
	return &AlibabawdkbmpaiyangstockqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkbmpaiyangstockqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.bm.paiyang.stock.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkbmpaiyangstockqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkbmpaiyangstockqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIsvShopStockParam is IsvShopStockParam Setter
// 请求入参
func (r *AlibabawdkbmpaiyangstockqueryAPIRequest) SetIsvShopStockParam(_isvShopStockParam *IsvShopStockParam) error {
	r._isvShopStockParam = _isvShopStockParam
	r.Set("isv_shop_stock_param", _isvShopStockParam)
	return nil
}

// GetIsvShopStockParam IsvShopStockParam Getter
func (r AlibabawdkbmpaiyangstockqueryAPIRequest) GetIsvShopStockParam() *IsvShopStockParam {
	return r._isvShopStockParam
}
