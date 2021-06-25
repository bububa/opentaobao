package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品sku编辑接口 APIRequest
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

func NewTaobaoFenxiaoProductSkuUpdateRequest() *TaobaoFenxiaoProductSkuUpdateRequest{
    return &TaobaoFenxiaoProductSkuUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFenxiaoProductSkuUpdateRequest) GetApiMethodName() string {
    return "taobao.fenxiao.product.sku.update"
}

func (r TaobaoFenxiaoProductSkuUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFenxiaoProductSkuUpdateRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r TaobaoFenxiaoProductSkuUpdateRequest) GetProductId() int64 {
    return r.productId
}

func (r *TaobaoFenxiaoProductSkuUpdateRequest) SetQuantity(quantity int64) error {
    r.quantity = quantity
    r.Set("quantity", quantity)
    return nil
}

func (r TaobaoFenxiaoProductSkuUpdateRequest) GetQuantity() int64 {
    return r.quantity
}

func (r *TaobaoFenxiaoProductSkuUpdateRequest) SetStandardPrice(standardPrice string) error {
    r.standardPrice = standardPrice
    r.Set("standard_price", standardPrice)
    return nil
}

func (r TaobaoFenxiaoProductSkuUpdateRequest) GetStandardPrice() string {
    return r.standardPrice
}

func (r *TaobaoFenxiaoProductSkuUpdateRequest) SetAgentCostPrice(agentCostPrice string) error {
    r.agentCostPrice = agentCostPrice
    r.Set("agent_cost_price", agentCostPrice)
    return nil
}

func (r TaobaoFenxiaoProductSkuUpdateRequest) GetAgentCostPrice() string {
    return r.agentCostPrice
}

func (r *TaobaoFenxiaoProductSkuUpdateRequest) SetProperties(properties string) error {
    r.properties = properties
    r.Set("properties", properties)
    return nil
}

func (r TaobaoFenxiaoProductSkuUpdateRequest) GetProperties() string {
    return r.properties
}

func (r *TaobaoFenxiaoProductSkuUpdateRequest) SetSkuNumber(skuNumber string) error {
    r.skuNumber = skuNumber
    r.Set("sku_number", skuNumber)
    return nil
}

func (r TaobaoFenxiaoProductSkuUpdateRequest) GetSkuNumber() string {
    return r.skuNumber
}

func (r *TaobaoFenxiaoProductSkuUpdateRequest) SetDealerCostPrice(dealerCostPrice string) error {
    r.dealerCostPrice = dealerCostPrice
    r.Set("dealer_cost_price", dealerCostPrice)
    return nil
}

func (r TaobaoFenxiaoProductSkuUpdateRequest) GetDealerCostPrice() string {
    return r.dealerCostPrice
}

