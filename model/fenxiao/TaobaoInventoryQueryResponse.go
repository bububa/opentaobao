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
    TaobaoInventoryQueryResponse
}

type TaobaoInventoryQueryResponse struct {
    XMLName xml.Name `xml:"inventory_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品总体库存信息
    
    ItemInventorys   []InventorySum `json:"item_inventorys,omitempty" xml:"item_inventorys>inventory_sum,omitempty"`
    
    
    // 提示信息，提示不存在的后端商品
    
    TipInfos   []TipInfo `json:"tip_infos,omitempty" xml:"tip_infos>tip_info,omitempty"`
    
    
}
