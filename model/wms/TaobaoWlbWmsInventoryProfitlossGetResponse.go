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
    TaobaoWlbWmsInventoryProfitlossGetResponse
}

type TaobaoWlbWmsInventoryProfitlossGetResponse struct {
    XMLName xml.Name `xml:"wlb_wms_inventory_profitloss_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 损益信息
    
    ProfitLossInfo   *CainiaoInventoryProfitlossProfitlossinfo `json:"profit_loss_info,omitempty" xml:"profit_loss_info,omitempty"`

    
}
