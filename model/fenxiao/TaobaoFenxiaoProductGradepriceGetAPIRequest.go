package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaoproductgradepricegetAPIRequest 等级折扣查询 API请求
// taobao.fenxiao.product.gradeprice.get
//
// 等级折扣查询
type TaobaofenxiaoproductgradepricegetAPIRequest struct {
	model.Params
	// 产品id
	_productId int64
	// skuId
	_skuId int64
	// 经、代销模式（1：代销、2：经销）
	_tradeType int64
}

// NewTaobaofenxiaoproductgradepricegetRequest 初始化TaobaofenxiaoproductgradepricegetAPIRequest对象
func NewTaobaofenxiaoproductgradepricegetRequest() *TaobaofenxiaoproductgradepricegetAPIRequest {
	return &TaobaofenxiaoproductgradepricegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaoproductgradepricegetAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.gradeprice.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaoproductgradepricegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaoproductgradepricegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// 产品id
func (r *TaobaofenxiaoproductgradepricegetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaofenxiaoproductgradepricegetAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetSkuId is SkuId Setter
// skuId
func (r *TaobaofenxiaoproductgradepricegetAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TaobaofenxiaoproductgradepricegetAPIRequest) GetSkuId() int64 {
	return r._skuId
}

// SetTradeType is TradeType Setter
// 经、代销模式（1：代销、2：经销）
func (r *TaobaofenxiaoproductgradepricegetAPIRequest) SetTradeType(_tradeType int64) error {
	r._tradeType = _tradeType
	r.Set("trade_type", _tradeType)
	return nil
}

// GetTradeType TradeType Getter
func (r TaobaofenxiaoproductgradepricegetAPIRequest) GetTradeType() int64 {
	return r._tradeType
}
