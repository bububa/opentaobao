package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductSkuAddAPIRequest 产品sku添加接口 API请求
// taobao.fenxiao.product.sku.add
//
// 添加产品SKU信息
type TaobaoFenxiaoProductSkuAddAPIRequest struct {
	model.Params
	// 采购基准价，最大值1000000000
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
	// sku产品库存，在0到1000000之间，如果不传，则库存为0
	_quantity int64
}

// NewTaobaoFenxiaoProductSkuAddRequest 初始化TaobaoFenxiaoProductSkuAddAPIRequest对象
func NewTaobaoFenxiaoProductSkuAddRequest() *TaobaoFenxiaoProductSkuAddAPIRequest {
	return &TaobaoFenxiaoProductSkuAddAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFenxiaoProductSkuAddAPIRequest) Reset() {
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
func (r TaobaoFenxiaoProductSkuAddAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.sku.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductSkuAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFenxiaoProductSkuAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStandardPrice is StandardPrice Setter
// 采购基准价，最大值1000000000
func (r *TaobaoFenxiaoProductSkuAddAPIRequest) SetStandardPrice(_standardPrice string) error {
	r._standardPrice = _standardPrice
	r.Set("standard_price", _standardPrice)
	return nil
}

// GetStandardPrice StandardPrice Getter
func (r TaobaoFenxiaoProductSkuAddAPIRequest) GetStandardPrice() string {
	return r._standardPrice
}

// SetAgentCostPrice is AgentCostPrice Setter
// 代销采购价
func (r *TaobaoFenxiaoProductSkuAddAPIRequest) SetAgentCostPrice(_agentCostPrice string) error {
	r._agentCostPrice = _agentCostPrice
	r.Set("agent_cost_price", _agentCostPrice)
	return nil
}

// GetAgentCostPrice AgentCostPrice Getter
func (r TaobaoFenxiaoProductSkuAddAPIRequest) GetAgentCostPrice() string {
	return r._agentCostPrice
}

// SetDealerCostPrice is DealerCostPrice Setter
// 经销采购价
func (r *TaobaoFenxiaoProductSkuAddAPIRequest) SetDealerCostPrice(_dealerCostPrice string) error {
	r._dealerCostPrice = _dealerCostPrice
	r.Set("dealer_cost_price", _dealerCostPrice)
	return nil
}

// GetDealerCostPrice DealerCostPrice Getter
func (r TaobaoFenxiaoProductSkuAddAPIRequest) GetDealerCostPrice() string {
	return r._dealerCostPrice
}

// SetSkuNumber is SkuNumber Setter
// 商家编码
func (r *TaobaoFenxiaoProductSkuAddAPIRequest) SetSkuNumber(_skuNumber string) error {
	r._skuNumber = _skuNumber
	r.Set("sku_number", _skuNumber)
	return nil
}

// GetSkuNumber SkuNumber Getter
func (r TaobaoFenxiaoProductSkuAddAPIRequest) GetSkuNumber() string {
	return r._skuNumber
}

// SetProperties is Properties Setter
// sku属性
func (r *TaobaoFenxiaoProductSkuAddAPIRequest) SetProperties(_properties string) error {
	r._properties = _properties
	r.Set("properties", _properties)
	return nil
}

// GetProperties Properties Getter
func (r TaobaoFenxiaoProductSkuAddAPIRequest) GetProperties() string {
	return r._properties
}

// SetProductId is ProductId Setter
// 产品ID
func (r *TaobaoFenxiaoProductSkuAddAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaoFenxiaoProductSkuAddAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetQuantity is Quantity Setter
// sku产品库存，在0到1000000之间，如果不传，则库存为0
func (r *TaobaoFenxiaoProductSkuAddAPIRequest) SetQuantity(_quantity int64) error {
	r._quantity = _quantity
	r.Set("quantity", _quantity)
	return nil
}

// GetQuantity Quantity Getter
func (r TaobaoFenxiaoProductSkuAddAPIRequest) GetQuantity() int64 {
	return r._quantity
}

var poolTaobaoFenxiaoProductSkuAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFenxiaoProductSkuAddRequest()
	},
}

// GetTaobaoFenxiaoProductSkuAddRequest 从 sync.Pool 获取 TaobaoFenxiaoProductSkuAddAPIRequest
func GetTaobaoFenxiaoProductSkuAddAPIRequest() *TaobaoFenxiaoProductSkuAddAPIRequest {
	return poolTaobaoFenxiaoProductSkuAddAPIRequest.Get().(*TaobaoFenxiaoProductSkuAddAPIRequest)
}

// ReleaseTaobaoFenxiaoProductSkuAddAPIRequest 将 TaobaoFenxiaoProductSkuAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoFenxiaoProductSkuAddAPIRequest(v *TaobaoFenxiaoProductSkuAddAPIRequest) {
	v.Reset()
	poolTaobaoFenxiaoProductSkuAddAPIRequest.Put(v)
}
