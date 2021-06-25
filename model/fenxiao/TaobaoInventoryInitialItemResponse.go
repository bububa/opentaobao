package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品库存初始化 APIResponse
taobao.inventory.initial.item

建议使用新接口：taobao.inventory.merchant.adjust ，该接口会逐步停用。
商家仓商品初始化在各个仓中库存
*/
type TaobaoInventoryInitialItemAPIResponse struct {
    model.CommonResponse
    Response *TaobaoInventoryInitialItemResponse `json:"taobao_inventory_initial_item_response,omitempty"`
}

type TaobaoInventoryInitialItemResponse struct {

    // 提示信息
    TipInfos   []TipInfo `json:"tip_infos,omitempty"`

}
