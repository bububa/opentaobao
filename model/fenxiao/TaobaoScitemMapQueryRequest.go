package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查找IC商品或分销商品与后端商品的关联信息 APIRequest
taobao.scitem.map.query

查找IC商品或分销商品与后端商品的关联信息。skuId如果不传就查找该itemId下所有的sku
*/
type TaobaoScitemMapQueryRequest struct {
    model.Params

    // map_type为1：前台ic商品id<br/>map_type为2：分销productid
    itemId   int64 

    // map_type为1：前台ic商品skuid <br/>map_type为2：分销商品的skuid
    skuId   int64 

}

func NewTaobaoScitemMapQueryRequest() *TaobaoScitemMapQueryRequest{
    return &TaobaoScitemMapQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoScitemMapQueryRequest) GetApiMethodName() string {
    return "taobao.scitem.map.query"
}

func (r TaobaoScitemMapQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoScitemMapQueryRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoScitemMapQueryRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TaobaoScitemMapQueryRequest) SetSkuId(skuId int64) error {
    r.skuId = skuId
    r.Set("sku_id", skuId)
    return nil
}

func (r TaobaoScitemMapQueryRequest) GetSkuId() int64 {
    return r.skuId
}

