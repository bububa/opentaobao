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
type TaobaoFenxiaoProductSkuUpdateRequest struct {
    model.Params
    // 产品ID
    productId   int64
    // 产品SKU库存
    quantity   int64
    // 采购基准价
    standardPrice   string
    // 代销采购价
    agentCostPrice   string
    // sku属性
    properties   string
    // 商家编码
    skuNumber   string
    // 经销采购价
    dealerCostPrice   string
}

// 初始化TaobaoFenxiaoProductSkuUpdateRequest对象
func NewTaobaoFenxiaoProductSkuUpdateRequest() *TaobaoFenxiaoProductSkuUpdateRequest{
    return &TaobaoFenxiaoProductSkuUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductSkuUpdateRequest) GetApiMethodName() string {
    return "taobao.fenxiao.product.sku.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductSkuUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 产品ID
func (r *TaobaoFenxiaoProductSkuUpdateRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r TaobaoFenxiaoProductSkuUpdateRequest) GetProductId() int64 {
    return r.productId
}
// Quantity Setter
// 产品SKU库存
func (r *TaobaoFenxiaoProductSkuUpdateRequest) SetQuantity(quantity int64) error {
    r.quantity = quantity
    r.Set("quantity", quantity)
    return nil
}

// Quantity Getter
func (r TaobaoFenxiaoProductSkuUpdateRequest) GetQuantity() int64 {
    return r.quantity
}
// StandardPrice Setter
// 采购基准价
func (r *TaobaoFenxiaoProductSkuUpdateRequest) SetStandardPrice(standardPrice string) error {
    r.standardPrice = standardPrice
    r.Set("standard_price", standardPrice)
    return nil
}

// StandardPrice Getter
func (r TaobaoFenxiaoProductSkuUpdateRequest) GetStandardPrice() string {
    return r.standardPrice
}
// AgentCostPrice Setter
// 代销采购价
func (r *TaobaoFenxiaoProductSkuUpdateRequest) SetAgentCostPrice(agentCostPrice string) error {
    r.agentCostPrice = agentCostPrice
    r.Set("agent_cost_price", agentCostPrice)
    return nil
}

// AgentCostPrice Getter
func (r TaobaoFenxiaoProductSkuUpdateRequest) GetAgentCostPrice() string {
    return r.agentCostPrice
}
// Properties Setter
// sku属性
func (r *TaobaoFenxiaoProductSkuUpdateRequest) SetProperties(properties string) error {
    r.properties = properties
    r.Set("properties", properties)
    return nil
}

// Properties Getter
func (r TaobaoFenxiaoProductSkuUpdateRequest) GetProperties() string {
    return r.properties
}
// SkuNumber Setter
// 商家编码
func (r *TaobaoFenxiaoProductSkuUpdateRequest) SetSkuNumber(skuNumber string) error {
    r.skuNumber = skuNumber
    r.Set("sku_number", skuNumber)
    return nil
}

// SkuNumber Getter
func (r TaobaoFenxiaoProductSkuUpdateRequest) GetSkuNumber() string {
    return r.skuNumber
}
// DealerCostPrice Setter
// 经销采购价
func (r *TaobaoFenxiaoProductSkuUpdateRequest) SetDealerCostPrice(dealerCostPrice string) error {
    r.dealerCostPrice = dealerCostPrice
    r.Set("dealer_cost_price", dealerCostPrice)
    return nil
}

// DealerCostPrice Getter
func (r TaobaoFenxiaoProductSkuUpdateRequest) GetDealerCostPrice() string {
    return r.dealerCostPrice
}
