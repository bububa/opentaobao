package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
编辑商品 API请求
taobao.banamadpc.item.update

巴拿马供应商通过此接口编辑商品
*/
type TaobaoBanamadpcItemUpdateRequest struct {
    model.Params
    // 商品的schema xml
    _xml   string
    // 商品id
    _itemId   int64
}

// 初始化TaobaoBanamadpcItemUpdateRequest对象
func NewTaobaoBanamadpcItemUpdateRequest() *TaobaoBanamadpcItemUpdateRequest{
    return &TaobaoBanamadpcItemUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBanamadpcItemUpdateRequest) GetApiMethodName() string {
    return "taobao.banamadpc.item.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBanamadpcItemUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Xml Setter
// 商品的schema xml
func (r *TaobaoBanamadpcItemUpdateRequest) SetXml(_xml string) error {
    r._xml = _xml
    r.Set("xml", _xml)
    return nil
}

// Xml Getter
func (r TaobaoBanamadpcItemUpdateRequest) GetXml() string {
    return r._xml
}
// ItemId Setter
// 商品id
func (r *TaobaoBanamadpcItemUpdateRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoBanamadpcItemUpdateRequest) GetItemId() int64 {
    return r._itemId
}
