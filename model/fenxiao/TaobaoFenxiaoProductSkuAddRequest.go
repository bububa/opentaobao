package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品sku添加接口 APIRequest
taobao.fenxiao.product.sku.add

添加产品SKU信息
*/
type TaobaoFenxiaoProductSkuAddRequest struct {
    model.Params

    // 产品ID
    productId   int64 

    // sku产品库存，在0到1000000之间，如果不传，则库存为0
    quantity   int64 

    // 采购基准价，最大值1000000000
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

func NewTaobaoFenxiaoProductSkuAddRequest() *TaobaoFenxiaoProductSkuAddRequest{
    return &TaobaoFenxiaoProductSkuAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFenxiaoProductSkuAddRequest) GetApiMethodName() string {
    return "taobao.fenxiao.product.sku.add"
}

func (r TaobaoFenxiaoProductSkuAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFenxiaoProductSkuAddRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r TaobaoFenxiaoProductSkuAddRequest) GetProductId() int64 {
    return r.productId
}

func (r *TaobaoFenxiaoProductSkuAddRequest) SetQuantity(quantity int64) error {
    r.quantity = quantity
    r.Set("quantity", quantity)
    return nil
}

func (r TaobaoFenxiaoProductSkuAddRequest) GetQuantity() int64 {
    return r.quantity
}

func (r *TaobaoFenxiaoProductSkuAddRequest) SetStandardPrice(standardPrice string) error {
    r.standardPrice = standardPrice
    r.Set("standard_price", standardPrice)
    return nil
}

func (r TaobaoFenxiaoProductSkuAddRequest) GetStandardPrice() string {
    return r.standardPrice
}

func (r *TaobaoFenxiaoProductSkuAddRequest) SetAgentCostPrice(agentCostPrice string) error {
    r.agentCostPrice = agentCostPrice
    r.Set("agent_cost_price", agentCostPrice)
    return nil
}

func (r TaobaoFenxiaoProductSkuAddRequest) GetAgentCostPrice() string {
    return r.agentCostPrice
}

func (r *TaobaoFenxiaoProductSkuAddRequest) SetProperties(properties string) error {
    r.properties = properties
    r.Set("properties", properties)
    return nil
}

func (r TaobaoFenxiaoProductSkuAddRequest) GetProperties() string {
    return r.properties
}

func (r *TaobaoFenxiaoProductSkuAddRequest) SetSkuNumber(skuNumber string) error {
    r.skuNumber = skuNumber
    r.Set("sku_number", skuNumber)
    return nil
}

func (r TaobaoFenxiaoProductSkuAddRequest) GetSkuNumber() string {
    return r.skuNumber
}

func (r *TaobaoFenxiaoProductSkuAddRequest) SetDealerCostPrice(dealerCostPrice string) error {
    r.dealerCostPrice = dealerCostPrice
    r.Set("dealer_cost_price", dealerCostPrice)
    return nil
}

func (r TaobaoFenxiaoProductSkuAddRequest) GetDealerCostPrice() string {
    return r.dealerCostPrice
}

