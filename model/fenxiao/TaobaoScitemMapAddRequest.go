package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建IC商品与后端商品的映射关系 API请求
taobao.scitem.map.add

创建IC商品或分销商品与后端商品的映射关系
*/
type TaobaoScitemMapAddRequest struct {
    model.Params
    // 前台ic商品id
    _itemId   int64
    // 前台ic商品skuid
    _skuId   int64
    // sc_item_id和outer_code 其中一个不能为空
    _scItemId   int64
    // sc_item_id和outer_code 其中一个不能为空
    _outerCode   string
    // 默认值为false<br/>true:进行高级校验,前端商品或SKU的商家编码必须与后端商品的商家编码一致，否则会拒绝关联<br/>false:不进行高级校验
    _needCheck   bool
}

// 初始化TaobaoScitemMapAddRequest对象
func NewTaobaoScitemMapAddRequest() *TaobaoScitemMapAddRequest{
    return &TaobaoScitemMapAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoScitemMapAddRequest) GetApiMethodName() string {
    return "taobao.scitem.map.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoScitemMapAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 前台ic商品id
func (r *TaobaoScitemMapAddRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoScitemMapAddRequest) GetItemId() int64 {
    return r._itemId
}
// SkuId Setter
// 前台ic商品skuid
func (r *TaobaoScitemMapAddRequest) SetSkuId(_skuId int64) error {
    r._skuId = _skuId
    r.Set("sku_id", _skuId)
    return nil
}

// SkuId Getter
func (r TaobaoScitemMapAddRequest) GetSkuId() int64 {
    return r._skuId
}
// ScItemId Setter
// sc_item_id和outer_code 其中一个不能为空
func (r *TaobaoScitemMapAddRequest) SetScItemId(_scItemId int64) error {
    r._scItemId = _scItemId
    r.Set("sc_item_id", _scItemId)
    return nil
}

// ScItemId Getter
func (r TaobaoScitemMapAddRequest) GetScItemId() int64 {
    return r._scItemId
}
// OuterCode Setter
// sc_item_id和outer_code 其中一个不能为空
func (r *TaobaoScitemMapAddRequest) SetOuterCode(_outerCode string) error {
    r._outerCode = _outerCode
    r.Set("outer_code", _outerCode)
    return nil
}

// OuterCode Getter
func (r TaobaoScitemMapAddRequest) GetOuterCode() string {
    return r._outerCode
}
// NeedCheck Setter
// 默认值为false<br/>true:进行高级校验,前端商品或SKU的商家编码必须与后端商品的商家编码一致，否则会拒绝关联<br/>false:不进行高级校验
func (r *TaobaoScitemMapAddRequest) SetNeedCheck(_needCheck bool) error {
    r._needCheck = _needCheck
    r.Set("need_check", _needCheck)
    return nil
}

// NeedCheck Getter
func (r TaobaoScitemMapAddRequest) GetNeedCheck() bool {
    return r._needCheck
}
