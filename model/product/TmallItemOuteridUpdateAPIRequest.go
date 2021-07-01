package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫商品/SKU商家编码更新接口 API请求
tmall.item.outerid.update

天猫商品/SKU商家编码更新接口；支持商品、SKU的商家编码同时更新；支持同一商品下的SKU批量更新。（感谢sample小雨提供接口命名）
*/
type TmallItemOuteridUpdateAPIRequest struct {
    model.Params
    // 商品ID
    _itemId   int64
    // 商品维度商家编码，如果不修改可以不传；清空请设置空串
    _outerId   string
    // 商品SKU更新OuterId时候用的数据
    _skuOuters   []UpdateSkuOuterId
}

// 初始化TmallItemOuteridUpdateAPIRequest对象
func NewTmallItemOuteridUpdateRequest() *TmallItemOuteridUpdateAPIRequest{
    return &TmallItemOuteridUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemOuteridUpdateAPIRequest) GetApiMethodName() string {
    return "tmall.item.outerid.update"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemOuteridUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品ID
func (r *TmallItemOuteridUpdateAPIRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TmallItemOuteridUpdateAPIRequest) GetItemId() int64 {
    return r._itemId
}
// OuterId Setter
// 商品维度商家编码，如果不修改可以不传；清空请设置空串
func (r *TmallItemOuteridUpdateAPIRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r TmallItemOuteridUpdateAPIRequest) GetOuterId() string {
    return r._outerId
}
// SkuOuters Setter
// 商品SKU更新OuterId时候用的数据
func (r *TmallItemOuteridUpdateAPIRequest) SetSkuOuters(_skuOuters []UpdateSkuOuterId) error {
    r._skuOuters = _skuOuters
    r.Set("sku_outers", _skuOuters)
    return nil
}

// SkuOuters Getter
func (r TmallItemOuteridUpdateAPIRequest) GetSkuOuters() []UpdateSkuOuterId {
    return r._skuOuters
}
