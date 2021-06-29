package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
编辑商品发布页 API请求
taobao.banamadpc.item.edit.render

巴拿马供应商通过此接口获取编辑商品发布页
*/
type TaobaoBanamadpcItemEditRenderRequest struct {
    model.Params
    // 商品id
    itemId   int64
}

// 初始化TaobaoBanamadpcItemEditRenderRequest对象
func NewTaobaoBanamadpcItemEditRenderRequest() *TaobaoBanamadpcItemEditRenderRequest{
    return &TaobaoBanamadpcItemEditRenderRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBanamadpcItemEditRenderRequest) GetApiMethodName() string {
    return "taobao.banamadpc.item.edit.render"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBanamadpcItemEditRenderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *TaobaoBanamadpcItemEditRenderRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoBanamadpcItemEditRenderRequest) GetItemId() int64 {
    return r.itemId
}
