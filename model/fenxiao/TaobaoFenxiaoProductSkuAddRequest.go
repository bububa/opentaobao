package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品sku添加接口 API请求
taobao.fenxiao.product.sku.add

添加产品SKU信息
*/
type TaobaoFenxiaoProductSkuAddRequest struct {
    model.Params
    // 产品ID
    _productId   int64
    // sku产品库存，在0到1000000之间，如果不传，则库存为0
    _quantity   int64
    // 采购基准价，最大值1000000000
    _standardPrice   string
    // 代销采购价
    _agentCostPrice   string
    // sku属性
    _properties   string
    // 商家编码
    _skuNumber   string
    // 经销采购价
    _dealerCostPrice   string
}

// 初始化TaobaoFenxiaoProductSkuAddRequest对象
func NewTaobaoFenxiaoProductSkuAddRequest() *TaobaoFenxiaoProductSkuAddRequest{
    return &TaobaoFenxiaoProductSkuAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductSkuAddRequest) GetApiMethodName() string {
    return "taobao.fenxiao.product.sku.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductSkuAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 产品ID
func (r *TaobaoFenxiaoProductSkuAddRequest) SetProductId(_productId int64) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r TaobaoFenxiaoProductSkuAddRequest) GetProductId() int64 {
    return r._productId
}
// Quantity Setter
// sku产品库存，在0到1000000之间，如果不传，则库存为0
func (r *TaobaoFenxiaoProductSkuAddRequest) SetQuantity(_quantity int64) error {
    r._quantity = _quantity
    r.Set("quantity", _quantity)
    return nil
}

// Quantity Getter
func (r TaobaoFenxiaoProductSkuAddRequest) GetQuantity() int64 {
    return r._quantity
}
// StandardPrice Setter
// 采购基准价，最大值1000000000
func (r *TaobaoFenxiaoProductSkuAddRequest) SetStandardPrice(_standardPrice string) error {
    r._standardPrice = _standardPrice
    r.Set("standard_price", _standardPrice)
    return nil
}

// StandardPrice Getter
func (r TaobaoFenxiaoProductSkuAddRequest) GetStandardPrice() string {
    return r._standardPrice
}
// AgentCostPrice Setter
// 代销采购价
func (r *TaobaoFenxiaoProductSkuAddRequest) SetAgentCostPrice(_agentCostPrice string) error {
    r._agentCostPrice = _agentCostPrice
    r.Set("agent_cost_price", _agentCostPrice)
    return nil
}

// AgentCostPrice Getter
func (r TaobaoFenxiaoProductSkuAddRequest) GetAgentCostPrice() string {
    return r._agentCostPrice
}
// Properties Setter
// sku属性
func (r *TaobaoFenxiaoProductSkuAddRequest) SetProperties(_properties string) error {
    r._properties = _properties
    r.Set("properties", _properties)
    return nil
}

// Properties Getter
func (r TaobaoFenxiaoProductSkuAddRequest) GetProperties() string {
    return r._properties
}
// SkuNumber Setter
// 商家编码
func (r *TaobaoFenxiaoProductSkuAddRequest) SetSkuNumber(_skuNumber string) error {
    r._skuNumber = _skuNumber
    r.Set("sku_number", _skuNumber)
    return nil
}

// SkuNumber Getter
func (r TaobaoFenxiaoProductSkuAddRequest) GetSkuNumber() string {
    return r._skuNumber
}
// DealerCostPrice Setter
// 经销采购价
func (r *TaobaoFenxiaoProductSkuAddRequest) SetDealerCostPrice(_dealerCostPrice string) error {
    r._dealerCostPrice = _dealerCostPrice
    r.Set("dealer_cost_price", _dealerCostPrice)
    return nil
}

// DealerCostPrice Getter
func (r TaobaoFenxiaoProductSkuAddRequest) GetDealerCostPrice() string {
    return r._dealerCostPrice
}
