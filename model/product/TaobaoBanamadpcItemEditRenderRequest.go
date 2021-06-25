package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
编辑商品发布页 APIRequest
taobao.banamadpc.item.edit.render

巴拿马供应商通过此接口获取编辑商品发布页
*/
type TaobaoBanamadpcItemEditRenderRequest struct {
    model.Params

    // 商品id
    itemId   int64 

}

func NewTaobaoBanamadpcItemEditRenderRequest() *TaobaoBanamadpcItemEditRenderRequest{
    return &TaobaoBanamadpcItemEditRenderRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBanamadpcItemEditRenderRequest) GetApiMethodName() string {
    return "taobao.banamadpc.item.edit.render"
}

func (r TaobaoBanamadpcItemEditRenderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBanamadpcItemEditRenderRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoBanamadpcItemEditRenderRequest) GetItemId() int64 {
    return r.itemId
}

