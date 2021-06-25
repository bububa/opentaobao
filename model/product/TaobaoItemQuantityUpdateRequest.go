package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
宝贝/SKU库存修改 APIRequest
taobao.item.quantity.update

提供按照全量或增量形式修改宝贝/SKU库存的功能
*/
type TaobaoItemQuantityUpdateRequest struct {
    model.Params

    // 商品数字ID，必填参数
    numIid   int64 

    // 要操作的SKU的数字ID，可选。如果不填默认修改宝贝的库存，如果填上则修改该SKU的库存
    skuId   int64 

    // SKU的商家编码，可选参数。如果不填则默认修改宝贝的库存，如果填了则按照商家编码搜索出对应的SKU并修改库存。当sku_id和本字段都填写时以sku_id为准搜索对应SKU
    outerId   string 

    // 库存修改值，必选。当全量更新库存时，quantity必须为大于等于0的正整数；当增量更新库存时，quantity为整数，可小于等于0。若增量更新时传入的库存为负数，则负数与实际库存之和不能小于0。比如当前实际库存为1，传入增量更新quantity=-1，库存改为0
    quantity   int64 

    // 库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新
    type   int64 

}

func NewTaobaoItemQuantityUpdateRequest() *TaobaoItemQuantityUpdateRequest{
    return &TaobaoItemQuantityUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoItemQuantityUpdateRequest) GetApiMethodName() string {
    return "taobao.item.quantity.update"
}

func (r TaobaoItemQuantityUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoItemQuantityUpdateRequest) SetNumIid(numIid int64) error {
    r.numIid = numIid
    r.Set("num_iid", numIid)
    return nil
}

func (r TaobaoItemQuantityUpdateRequest) GetNumIid() int64 {
    return r.numIid
}

func (r *TaobaoItemQuantityUpdateRequest) SetSkuId(skuId int64) error {
    r.skuId = skuId
    r.Set("sku_id", skuId)
    return nil
}

func (r TaobaoItemQuantityUpdateRequest) GetSkuId() int64 {
    return r.skuId
}

func (r *TaobaoItemQuantityUpdateRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r TaobaoItemQuantityUpdateRequest) GetOuterId() string {
    return r.outerId
}

func (r *TaobaoItemQuantityUpdateRequest) SetQuantity(quantity int64) error {
    r.quantity = quantity
    r.Set("quantity", quantity)
    return nil
}

func (r TaobaoItemQuantityUpdateRequest) GetQuantity() int64 {
    return r.quantity
}

func (r *TaobaoItemQuantityUpdateRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoItemQuantityUpdateRequest) GetType() int64 {
    return r.type
}

