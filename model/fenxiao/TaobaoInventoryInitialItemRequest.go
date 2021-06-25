package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品库存初始化 APIRequest
taobao.inventory.initial.item

建议使用新接口：taobao.inventory.merchant.adjust ，该接口会逐步停用。
商家仓商品初始化在各个仓中库存
*/
type TaobaoInventoryInitialItemRequest struct {
    model.Params

    // 后端商品id
    scItemId   int64 

    // 商品初始库存信息： [{"storeCode":"必选,商家仓库编号","inventoryType":"可选，库存类型 1：正常,2：损坏,3：冻结,10：质押,11-20:用户自定义,默认为1","quantity":"必选,数量"}]
    storeInventorys   string 

}

func NewTaobaoInventoryInitialItemRequest() *TaobaoInventoryInitialItemRequest{
    return &TaobaoInventoryInitialItemRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoInventoryInitialItemRequest) GetApiMethodName() string {
    return "taobao.inventory.initial.item"
}

func (r TaobaoInventoryInitialItemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoInventoryInitialItemRequest) SetScItemId(scItemId int64) error {
    r.scItemId = scItemId
    r.Set("sc_item_id", scItemId)
    return nil
}

func (r TaobaoInventoryInitialItemRequest) GetScItemId() int64 {
    return r.scItemId
}

func (r *TaobaoInventoryInitialItemRequest) SetStoreInventorys(storeInventorys string) error {
    r.storeInventorys = storeInventorys
    r.Set("store_inventorys", storeInventorys)
    return nil
}

func (r TaobaoInventoryInitialItemRequest) GetStoreInventorys() string {
    return r.storeInventorys
}

