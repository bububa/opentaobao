package wlb

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbInventoryDetailGetAPIResponse 查询库存详情 API返回值
// taobao.wlb.inventory.detail.get
//
// 查询库存详情，通过商品ID获取发送请求的卖家的库存详情
type TaobaoWlbInventoryDetailGetAPIResponse struct {
	model.CommonResponse
	TaobaoWlbInventoryDetailGetAPIResponseModel
}

// TaobaoWlbInventoryDetailGetAPIResponseModel is 查询库存详情 成功返回结果
type TaobaoWlbInventoryDetailGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_inventory_detail_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 库存详情对象。其中包括货主ID，仓库编码，库存，库存类型等属性
	InventoryList []WlbInventory `json:"inventory_list,omitempty" xml:"inventory_list>wlb_inventory,omitempty"`
	// 入参的item_id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}
