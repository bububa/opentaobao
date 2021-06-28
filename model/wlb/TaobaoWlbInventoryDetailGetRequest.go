package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询库存详情 APIRequest
taobao.wlb.inventory.detail.get

查询库存详情，通过商品ID获取发送请求的卖家的库存详情
*/
type TaobaoWlbInventoryDetailGetRequest struct {
    model.Params

    // 库存类型列表，值包括：<br/>VENDIBLE--可销售库存<br/>FREEZE--冻结库存<br/>ONWAY--在途库存<br/>DEFECT--残次品库存<br/>ENGINE_DAMAGE--机损<br/>BOX_DAMAGE--箱损<br/>EXPIRATION--过保
    inventoryTypeList   []string 

    // 仓库编码
    storeCode   string 

    // 商品ID
    itemId   int64 

}

func NewTaobaoWlbInventoryDetailGetRequest() *TaobaoWlbInventoryDetailGetRequest{
    return &TaobaoWlbInventoryDetailGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbInventoryDetailGetRequest) GetApiMethodName() string {
    return "taobao.wlb.inventory.detail.get"
}

func (r TaobaoWlbInventoryDetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbInventoryDetailGetRequest) SetInventoryTypeList(inventoryTypeList []string) error {
    r.inventoryTypeList = inventoryTypeList
    r.Set("inventory_type_list", inventoryTypeList)
    return nil
}

func (r TaobaoWlbInventoryDetailGetRequest) GetInventoryTypeList() []string {
    return r.inventoryTypeList
}

func (r *TaobaoWlbInventoryDetailGetRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

func (r TaobaoWlbInventoryDetailGetRequest) GetStoreCode() string {
    return r.storeCode
}

func (r *TaobaoWlbInventoryDetailGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoWlbInventoryDetailGetRequest) GetItemId() int64 {
    return r.itemId
}

