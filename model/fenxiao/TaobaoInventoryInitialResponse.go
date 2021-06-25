package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
库存初始化 APIResponse
taobao.inventory.initial

建议使用新接口：taobao.inventory.merchant.adjust ，该接口会逐步停用。
商家仓库存初始化接口，直接按照商家指定的商品库存数进行填充，没有单据核对，不参与库存对账。
*/
type TaobaoInventoryInitialAPIResponse struct {
    model.CommonResponse
    Response *TaobaoInventoryInitialResponse `json:"taobao_inventory_initial_response,omitempty"`
}

type TaobaoInventoryInitialResponse struct {

    // 提示信息
    TipInfos   []TipInfo `json:"tip_infos,omitempty"`

}
