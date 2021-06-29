package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发品资质校验 API请求
taobao.item.permit.check

对淘宝商品发品、编辑前的预校验接口
*/
type TaobaoItemPermitCheckRequest struct {
    model.Params
    // 商品id
    _itemId   int64
    // 类目id
    _cid   int64
    // 发布类型。可选值:fixed(一口价),auction(拍卖)
    _type   string
}

// 初始化TaobaoItemPermitCheckRequest对象
func NewTaobaoItemPermitCheckRequest() *TaobaoItemPermitCheckRequest{
    return &TaobaoItemPermitCheckRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItemPermitCheckRequest) GetApiMethodName() string {
    return "taobao.item.permit.check"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItemPermitCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *TaobaoItemPermitCheckRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoItemPermitCheckRequest) GetItemId() int64 {
    return r._itemId
}
// Cid Setter
// 类目id
func (r *TaobaoItemPermitCheckRequest) SetCid(_cid int64) error {
    r._cid = _cid
    r.Set("cid", _cid)
    return nil
}

// Cid Getter
func (r TaobaoItemPermitCheckRequest) GetCid() int64 {
    return r._cid
}
// Type Setter
// 发布类型。可选值:fixed(一口价),auction(拍卖)
func (r *TaobaoItemPermitCheckRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoItemPermitCheckRequest) GetType() string {
    return r._type
}
