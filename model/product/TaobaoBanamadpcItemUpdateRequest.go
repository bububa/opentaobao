package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
编辑商品 APIRequest
taobao.banamadpc.item.update

巴拿马供应商通过此接口编辑商品
*/
type TaobaoBanamadpcItemUpdateRequest struct {
    model.Params

    // 商品的schema xml
    xml   string 

    // 商品id
    itemId   int64 

}

func NewTaobaoBanamadpcItemUpdateRequest() *TaobaoBanamadpcItemUpdateRequest{
    return &TaobaoBanamadpcItemUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBanamadpcItemUpdateRequest) GetApiMethodName() string {
    return "taobao.banamadpc.item.update"
}

func (r TaobaoBanamadpcItemUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBanamadpcItemUpdateRequest) SetXml(xml string) error {
    r.xml = xml
    r.Set("xml", xml)
    return nil
}

func (r TaobaoBanamadpcItemUpdateRequest) GetXml() string {
    return r.xml
}

func (r *TaobaoBanamadpcItemUpdateRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoBanamadpcItemUpdateRequest) GetItemId() int64 {
    return r.itemId
}

