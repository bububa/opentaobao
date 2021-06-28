package fenxiao

import (
    "encoding/xml"

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
	RequestId     string         `json:"request_id,omitempty" xml:"inventory_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品总体库存信息
    
    ItemInventorys   []InventorySum `json:"item_inventorys,omitempty" xml:"