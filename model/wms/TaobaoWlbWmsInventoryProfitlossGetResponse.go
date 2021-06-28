package wms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
通过订单列表批量获取库存损益单信息 APIResponse
taobao.wlb.wms.inventory.profitloss.get

通过订单列表批量获取库存损益单信息
*/
type TaobaoWlbWmsInventoryProfitlossGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_wms_inventory_profitloss_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 损益信息
    
    ProfitLossInfo   *CainiaoInventoryProfitlossProfitlossinfo `json:"profit_loss_info,omitempty" xml:"