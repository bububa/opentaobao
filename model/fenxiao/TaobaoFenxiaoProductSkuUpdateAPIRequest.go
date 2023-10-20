package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaoproductskuupdateAPIRequest 产品sku编辑接口 API请求
// taobao.fenxiao.product.sku.update
//
// 产品SKU信息更新
type TaobaofenxiaoproductskuupdateAPIRequest struct {
	model.Params
	// 采购基准价
	_standardPrice string
	// 代销采购价
	_agentCostPrice string
	// 经销采购价
	_dealerCostPrice string
	// 商家编码
	_skuNumber string
	// sku属性
	_properties string
	// 产品ID
	_productId int64
	// 产品SKU库存
	_quantity int64
}

// NewTaobaofenxiaoproductskuupdateRequest 初始化TaobaofenxiaoproductskuupdateAPIRequest对象
func NewTaobaofenxiaoproductskuupdateRequest() *TaobaofenxiaoproductskuupdateAPIRequest {
	return &TaobaofenxiaoproductskuupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaoproductskuupdateAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.sku.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaoproductskuupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaoproductskuupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStandardPrice is StandardPrice Setter
// 采购基准价
func (r *TaobaofenxiaoproductskuupdateAPIRequest) SetStandardPrice(_standardPrice string) error {
	r._standardPrice = _standardPrice
	r.Set("standard_price", _standardPrice)
	return nil
}

// GetStandardPrice StandardPrice Getter
func (r TaobaofenxiaoproductskuupdateAPIRequest) GetStandardPrice() string {
	return r._standardPrice
}

// SetAgentCostPrice is AgentCostPrice Setter
// 代销采购价
func (r *TaobaofenxiaoproductskuupdateAPIRequest) SetAgentCostPrice(_agentCostPrice string) error {
	r._agentCostPrice = _agentCostPrice
	r.Set("agent_cost_price", _agentCostPrice)
	return nil
}

// GetAgentCostPrice AgentCostPrice Getter
func (r TaobaofenxiaoproductskuupdateAPIRequest) GetAgentCostPrice() string {
	return r._agentCostPrice
}

// SetDealerCostPrice is DealerCostPrice Setter
// 经销采购价
func (r *TaobaofenxiaoproductskuupdateAPIRequest) SetDealerCostPrice(_dealerCostPrice string) error {
	r._dealerCostPrice = _dealerCostPrice
	r.Set("dealer_cost_price", _dealerCostPrice)
	return nil
}

// GetDealerCostPrice DealerCostPrice Getter
func (r TaobaofenxiaoproductskuupdateAPIRequest) GetDealerCostPrice() string {
	return r._dealerCostPrice
}

// SetSkuNumber is SkuNumber Setter
// 商家编码
func (r *TaobaofenxiaoproductskuupdateAPIRequest) SetSkuNumber(_skuNumber string) error {
	r._skuNumber = _skuNumber
	r.Set("sku_number", _skuNumber)
	return nil
}

// GetSkuNumber SkuNumber Getter
func (r TaobaofenxiaoproductskuupdateAPIRequest) GetSkuNumber() string {
	return r._skuNumber
}

// SetProperties is Properties Setter
// sku属性
func (r *TaobaofenxiaoproductskuupdateAPIRequest) SetProperties(_properties string) error {
	r._properties = _properties
	r.Set("properties", _properties)
	return nil
}

// GetProperties Properties Getter
func (r TaobaofenxiaoproductskuupdateAPIRequest) GetProperties() string {
	return r._properties
}

// SetProductId is ProductId Setter
// 产品ID
func (r *TaobaofenxiaoproductskuupdateAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaofenxiaoproductskuupdateAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetQuantity is Quantity Setter
// 产品SKU库存
func (r *TaobaofenxiaoproductskuupdateAPIRequest) SetQuantity(_quantity int64) error {
	r._quantity = _quantity
	r.Set("quantity", _quantity)
	return nil
}

// GetQuantity Quantity Getter
func (r TaobaofenxiaoproductskuupdateAPIRequest) GetQuantity() int64 {
	return r._quantity
}
