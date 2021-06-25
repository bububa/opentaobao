package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发品资质校验 APIRequest
taobao.item.permit.check

对淘宝商品发品、编辑前的预校验接口
*/
type TaobaoItemPermitCheckRequest struct {
    model.Params

    // 商品id
    itemId   int64 

    // 类目id
    cid   int64 

    // 发布类型。可选值:fixed(一口价),auction(拍卖)
    type   string 

}

func NewTaobaoItemPermitCheckRequest() *TaobaoItemPermitCheckRequest{
    return &TaobaoItemPermitCheckRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoItemPermitCheckRequest) GetApiMethodName() string {
    return "taobao.item.permit.check"
}

func (r TaobaoItemPermitCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoItemPermitCheckRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoItemPermitCheckRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TaobaoItemPermitCheckRequest) SetCid(cid int64) error {
    r.cid = cid
    r.Set("cid", cid)
    return nil
}

func (r TaobaoItemPermitCheckRequest) GetCid() int64 {
    return r.cid
}

func (r *TaobaoItemPermitCheckRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoItemPermitCheckRequest) GetType() string {
    return r.type
}

