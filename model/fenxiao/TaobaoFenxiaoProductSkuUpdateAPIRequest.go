package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductSkuUpdateAPIRequest 产品sku编辑接口 API请求
// taobao.fenxiao.product.sku.update
//
// 产品SKU信息更新
type TaobaoFenxiaoProductSkuUpdateAPIRequest struct {
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

// NewTaobaoFenxiaoProductSkuUpdateRequest 初始化TaobaoFenxiaoProductSkuUpdateAPIRequest对象
func NewTaobaoFenxiaoProductSkuUpdateRequest() *TaobaoFenxiaoProductSkuUpdateAPIRequest {
	return &TaobaoFenxiaoProductSkuUpdateAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFenxiaoProductSkuUpdateAPIRequest) Reset() {
	r._standardPrice = ""
	r._agentCostPrice = ""
	r._dealerCostPrice = ""
	r._skuNumber = ""
	r._properties = ""
	r._productId = 0
	r._quantity = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductSkuUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.sku.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductSkuUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFenxiaoProductSkuUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStandardPrice is StandardPrice Setter
// 采购基准价
func (r *TaobaoFenxiaoProductSkuUpdateAPIRequest) SetStandardPrice(_standardPrice string) error {
	r._standardPrice = _standardPrice
	r.Set("standard_price", _standardPrice)
	return nil
}

// GetStandardPrice StandardPrice Getter
func (r TaobaoFenxiaoProductSkuUpdateAPIRequest) GetStandardPrice() string {
	return r._standardPrice
}

// SetAgentCostPrice is AgentCostPrice Setter
// 代销采购价
func (r *TaobaoFenxiaoProductSkuUpdateAPIRequest) SetAgentCostPrice(_agentCostPrice string) error {
	r._agentCostPrice = _agentCostPrice
	r.Set("agent_cost_price", _agentCostPrice)
	return nil
}

// GetAgentCostPrice AgentCostPrice Getter
func (r TaobaoFenxiaoProductSkuUpdateAPIRequest) GetAgentCostPrice() string {
	return r._agentCostPrice
}

// SetDealerCostPrice is DealerCostPrice Setter
// 经销采购价
func (r *TaobaoFenxiaoProductSkuUpdateAPIRequest) SetDealerCostPrice(_dealerCostPrice string) error {
	r._dealerCostPrice = _dealerCostPrice
	r.Set("dealer_cost_price", _dealerCostPrice)
	return nil
}

// GetDealerCostPrice DealerCostPrice Getter
func (r TaobaoFenxiaoProductSkuUpdateAPIRequest) GetDealerCostPrice() string {
	return r._dealerCostPrice
}

// SetSkuNumber is SkuNumber Setter
// 商家编码
func (r *TaobaoFenxiaoProductSkuUpdateAPIRequest) SetSkuNumber(_skuNumber string) error {
	r._skuNumber = _skuNumber
	r.Set("sku_number", _skuNumber)
	return nil
}

// GetSkuNumber SkuNumber Getter
func (r TaobaoFenxiaoProductSkuUpdateAPIRequest) GetSkuNumber() string {
	return r._skuNumber
}

// SetProperties is Properties Setter
// sku属性
func (r *TaobaoFenxiaoProductSkuUpdateAPIRequest) SetProperties(_properties string) error {
	r._properties = _properties
	r.Set("properties", _properties)
	return nil
}

// GetProperties Properties Getter
func (r TaobaoFenxiaoProductSkuUpdateAPIRequest) GetProperties() string {
	return r._properties
}

// SetProductId is ProductId Setter
// 产品ID
func (r *TaobaoFenxiaoProductSkuUpdateAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaoFenxiaoProductSkuUpdateAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetQuantity is Quantity Setter
// 产品SKU库存
func (r *TaobaoFenxiaoProductSkuUpdateAPIRequest) SetQuantity(_quantity int64) error {
	r._quantity = _quantity
	r.Set("quantity", _quantity)
	return nil
}

// GetQuantity Quantity Getter
func (r TaobaoFenxiaoProductSkuUpdateAPIRequest) GetQuantity() int64 {
	return r._quantity
}

var poolTaobaoFenxiaoProductSkuUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFenxiaoProductSkuUpdateRequest()
	},
}

// GetTaobaoFenxiaoProductSkuUpdateRequest 从 sync.Pool 获取 TaobaoFenxiaoProductSkuUpdateAPIRequest
func GetTaobaoFenxiaoProductSkuUpdateAPIRequest() *TaobaoFenxiaoProductSkuUpdateAPIRequest {
	return poolTaobaoFenxiaoProductSkuUpdateAPIRequest.Get().(*TaobaoFenxiaoProductSkuUpdateAPIRequest)
}

// ReleaseTaobaoFenxiaoProductSkuUpdateAPIRequest 将 TaobaoFenxiaoProductSkuUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoFenxiaoProductSkuUpdateAPIRequest(v *TaobaoFenxiaoProductSkuUpdateAPIRequest) {
	v.Reset()
	poolTaobaoFenxiaoProductSkuUpdateAPIRequest.Put(v)
}
