package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新交易的销售属性 API请求
taobao.trade.ordersku.update

只能更新发货前子订单的销售属性 <br/>只能更新价格相同的销售属性。对于拍下减库存的交易会同步更新销售属性的库存量。对于旺店的交易，要使用商品扩展信息中的SKU价格来比较。 <br/>必须使用sku_id或sku_props中的一个参数来更新，如果两个都传的话，sku_id优先
*/
type TaobaoTradeOrderskuUpdateRequest struct {
    model.Params
    // 子订单编号（对于单笔订单的交易可以传交易编号）。
    oid   int64
    // 销售属性编号，可以通过taobao.item.skus.get获取订单对应的商品的所有销售属性。
    skuId   int64
    // 销售属性组合串，格式：p1:v1;p2:v2，如：1627207:28329;20509:28314。可以通过taobao.item.skus.get获取订单对应的商品的所有销售属性。
    skuProps   string
}

// 初始化TaobaoTradeOrderskuUpdateRequest对象
func NewTaobaoTradeOrderskuUpdateRequest() *TaobaoTradeOrderskuUpdateRequest{
    return &TaobaoTradeOrderskuUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTradeOrderskuUpdateRequest) GetApiMethodName() string {
    return "taobao.trade.ordersku.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTradeOrderskuUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Oid Setter
// 子订单编号（对于单笔订单的交易可以传交易编号）。
func (r *TaobaoTradeOrderskuUpdateRequest) SetOid(oid int64) error {
    r.oid = oid
    r.Set("oid", oid)
    return nil
}

// Oid Getter
func (r TaobaoTradeOrderskuUpdateRequest) GetOid() int64 {
    return r.oid
}
// SkuId Setter
// 销售属性编号，可以通过taobao.item.skus.get获取订单对应的商品的所有销售属性。
func (r *TaobaoTradeOrderskuUpdateRequest) SetSkuId(skuId int64) error {
    r.skuId = skuId
    r.Set("sku_id", skuId)
    return nil
}

// SkuId Getter
func (r TaobaoTradeOrderskuUpdateRequest) GetSkuId() int64 {
    return r.skuId
}
// SkuProps Setter
// 销售属性组合串，格式：p1:v1;p2:v2，如：1627207:28329;20509:28314。可以通过taobao.item.skus.get获取订单对应的商品的所有销售属性。
func (r *TaobaoTradeOrderskuUpdateRequest) SetSkuProps(skuProps string) error {
    r.skuProps = skuProps
    r.Set("sku_props", skuProps)
    return nil
}

// SkuProps Getter
func (r TaobaoTradeOrderskuUpdateRequest) GetSkuProps() string {
    return r.skuProps
}
