package wms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟商品库存查询 APIResponse
taobao.wlb.wms.inventory.query

支持按汇总（不分批次和渠道的总的库存数量）、渠道、批次三类方式查询商品实时库存
*/
type TaobaoWlbWmsInventoryQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_wms_inventory_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    WlSuccess   bool `json:"wl_success,omitempty" xml:"