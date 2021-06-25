package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品信息查询 APIRequest
taobao.wlb.wms.sku.get

商品信息查询
*/
type TaobaoWlbWmsSkuGetRequest struct {
    model.Params

    // 菜鸟商品ID,与itemcode必须有一个值不为空
    itemCode   string 

    // 商家商品编码,与itemid必须有一个值不为空
    itemId   string 

    // 货主ID
    ownerUserId   string 

}

func NewTaobaoWlbWmsSkuGetRequest() *TaobaoWlbWmsSkuGetRequest{
    return &TaobaoWlbWmsSkuGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbWmsSkuGetRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.sku.get"
}

func (r TaobaoWlbWmsSkuGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbWmsSkuGetRequest) SetItemCode(itemCode string) error {
    r.itemCode = itemCode
    r.Set("item_code", itemCode)
    return nil
}

func (r TaobaoWlbWmsSkuGetRequest) GetItemCode() string {
    return r.itemCode
}

func (r *TaobaoWlbWmsSkuGetRequest) SetItemId(itemId string) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoWlbWmsSkuGetRequest) GetItemId() string {
    return r.itemId
}

func (r *TaobaoWlbWmsSkuGetRequest) SetOwnerUserId(ownerUserId string) error {
    r.ownerUserId = ownerUserId
    r.Set("owner_user_id", ownerUserId)
    return nil
}

func (r TaobaoWlbWmsSkuGetRequest) GetOwnerUserId() string {
    return r.ownerUserId
}

