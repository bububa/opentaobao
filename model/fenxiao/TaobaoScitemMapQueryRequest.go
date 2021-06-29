package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查找IC商品或分销商品与后端商品的关联信息 API请求
taobao.scitem.map.query

查找IC商品或分销商品与后端商品的关联信息。skuId如果不传就查找该itemId下所有的sku
*/
type TaobaoScitemMapQueryRequest struct {
    model.Params
    // map_type为1：前台ic商品id<br/>map_type为2：分销productid
    _itemId   int64
    // map_type为1：前台ic商品skuid <br/>map_type为2：分销商品的skuid
    _skuId   int64
}

// 初始化TaobaoScitemMapQueryRequest对象
func NewTaobaoScitemMapQueryRequest() *TaobaoScitemMapQueryRequest{
    return &TaobaoScitemMapQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoScitemMapQueryRequest) GetApiMethodName() string {
    return "taobao.scitem.map.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoScitemMapQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// map_type为1：前台ic商品id<br/>map_type为2：分销productid
func (r *TaobaoScitemMapQueryRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoScitemMapQueryRequest) GetItemId() int64 {
    return r._itemId
}
// SkuId Setter
// map_type为1：前台ic商品skuid <br/>map_type为2：分销商品的skuid
func (r *TaobaoScitemMapQueryRequest) SetSkuId(_skuId int64) error {
    r._skuId = _skuId
    r.Set("sku_id", _skuId)
    return nil
}

// SkuId Getter
func (r TaobaoScitemMapQueryRequest) GetSkuId() int64 {
    return r._skuId
}
