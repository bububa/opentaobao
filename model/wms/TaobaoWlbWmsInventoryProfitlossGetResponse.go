package wms

import (
    "github.com/bububa/opentaobao/model"
)

/* 
通过订单列表批量获取库存损益单信息 APIResponse
taobao.wlb.wms.inventory.profitloss.get

通过订单列表批量获取库存损益单信息
*/
type TaobaoWlbWmsInventoryProfitlossGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbWmsInventoryProfitlossGetResponse `json:"taobao_wlb_wms_inventory_profitloss_get_response,omitempty"`
}

type TaobaoWlbWmsInventoryProfitlossGetResponse struct {

    // 损益信息
    ProfitLossInfo   *CainiaoInventoryProfitlossProfitlossinfo `json:"profit_loss_info,omitempty"`

}
