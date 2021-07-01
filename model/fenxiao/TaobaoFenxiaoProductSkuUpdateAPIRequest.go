package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品sku编辑接口 API请求
taobao.fenxiao.product.sku.update

产品SKU信息更新
*/
type TaobaoFenxiaoProductSkuUpdateAPIRequest struct {
    model.Params
    // 产品ID
    _productId   int64
    // 产品SKU库存
    _quantity   int64
    // 采购基准价
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

// 初始化TaobaoFenxiaoProductSkuUpdateAPIRequest对象
func NewTaobaoFenxiaoProductSkuUpdateRequest() *TaobaoFenxiaoProductSkuUpdateAPIRequest{
    return &TaobaoFenxiaoProductSkuUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductSkuUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.fenxiao.product.sku.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductSkuUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 产品ID
func (r *TaobaoFenxiaoProductSkuUpdateAPIRequest) SetProductId(_productId int64) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r TaobaoFenxiaoProductSkuUpdateAPIRequest) GetProductId() int64 {
    return r._productId
}
// Quantity Setter
// 产品SKU库存
func (r *TaobaoFenxiaoProductSkuUpdateAPIRequest) SetQuantity(_quantity int64) error {
    r._quantity = _quantity
    r.Set("quantity", _quantity)
    return nil
}

// Quantity Getter
func (r TaobaoFenxiaoProductSkuUpdateAPIRequest) GetQuantity() int64 {
    return r._quantity
}
// StandardPrice Setter
// 采购基准价
func (r *TaobaoFenxiaoProductSkuUpdateAPIRequest) SetStandardPrice(_standardPrice string) error {
    r._standardPrice = _standardPrice
    r.Set("standard_price", _standardPrice)
    return nil
}

// StandardPrice Getter
func (r TaobaoFenxiaoProductSkuUpdateAPIRequest) GetStandardPrice() string {
    return r._standardPrice
}
// AgentCostPrice Setter
// 代销采购价
func (r *TaobaoFenxiaoProductSkuUpdateAPIRequest) SetAgentCostPrice(_agentCostPrice string) error {
    r._agentCostPrice = _agentCostPrice
    r.Set("agent_cost_price", _agentCostPrice)
    return nil
}

// AgentCostPrice Getter
func (r TaobaoFenxiaoProductSkuUpdateAPIRequest) GetAgentCostPrice() string {
    return r._agentCostPrice
}
// Properties Setter
// sku属性
func (r *TaobaoFenxiaoProductSkuUpdateAPIRequest) SetProperties(_properties string) error {
    r._properties = _properties
    r.Set("properties", _properties)
    return nil
}

// Properties Getter
func (r TaobaoFenxiaoProductSkuUpdateAPIRequest) GetProperties() string {
    return r._properties
}
// SkuNumber Setter
// 商家编码
func (r *TaobaoFenxiaoProductSkuUpdateAPIRequest) SetSkuNumber(_skuNumber string) error {
    r._skuNumber = _skuNumber
    r.Set("sku_number", _skuNumber)
    return nil
}

// SkuNumber Getter
func (r TaobaoFenxiaoProductSkuUpdateAPIRequest) GetSkuNumber() string {
    return r._skuNumber
}
// DealerCostPrice Setter
// 经销采购价
func (r *TaobaoFenxiaoProductSkuUpdateAPIRequest) SetDealerCostPrice(_dealerCostPrice string) error {
    r._dealerCostPrice = _dealerCostPrice
    r.Set("dealer_cost_price", _dealerCostPrice)
    return nil
}

// DealerCostPrice Getter
func (r TaobaoFenxiaoProductSkuUpdateAPIRequest) GetDealerCostPrice() string {
    return r._dealerCostPrice
}
