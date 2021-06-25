package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询商品库存信息 APIResponse
taobao.inventory.query

建议使用新接口：tmall.inventory.query.forstore ，新ISV不推荐使用。
商家查询商品总体库存信息
*/
type TaobaoInventoryQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoInventoryQueryResponse `json:"taobao_inventory_query_response,omitempty"`
}

type TaobaoInventoryQueryResponse struct {

    // 商品总体库存信息
    ItemInventorys   []InventorySum `json:"item_inventorys,omitempty"`

    // 提示信息，提示不存在的后端商品
    TipInfos   []TipInfo `json:"tip_infos,omitempty"`

}
