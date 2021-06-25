package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
库存初始化 APIRequest
taobao.inventory.initial

建议使用新接口：taobao.inventory.merchant.adjust ，该接口会逐步停用。
商家仓库存初始化接口，直接按照商家指定的商品库存数进行填充，没有单据核对，不参与库存对账。
*/
type TaobaoInventoryInitialRequest struct {
    model.Params

    // 商家仓库编码
    storeCode   string 

    // 商品初始库存信息： [{"scItemId":"商品后端ID，如果有传scItemCode,参数可以为0","scItemCode":"商品商家编码","inventoryType":"库存类型  1：正常,2：损坏,3：冻结,10：质押,11-20:用户自定义","quantity":"数量"}]
    items   string 

}

func NewTaobaoInventoryInitialRequest() *TaobaoInventoryInitialRequest{
    return &TaobaoInventoryInitialRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoInventoryInitialRequest) GetApiMethodName() string {
    return "taobao.inventory.initial"
}

func (r TaobaoInventoryInitialRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoInventoryInitialRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

func (r TaobaoInventoryInitialRequest) GetStoreCode() string {
    return r.storeCode
}

func (r *TaobaoInventoryInitialRequest) SetItems(items string) error {
    r.items = items
    r.Set("items", items)
    return nil
}

func (r TaobaoInventoryInitialRequest) GetItems() string {
    return r.items
}

