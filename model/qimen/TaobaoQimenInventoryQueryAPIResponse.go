package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenInventoryQueryAPIResponse 库存查询接口（多商品） API返回值
// taobao.qimen.inventory.query
//
// taobao.qimen.inventory.query
type TaobaoQimenInventoryQueryAPIResponse struct {
	model.CommonResponse
	TaobaoQimenInventoryQueryAPIResponseModel
}

// TaobaoQimenInventoryQueryAPIResponseModel is 库存查询接口（多商品） 成功返回结果
type TaobaoQimenInventoryQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_inventory_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *InventoryQueryResponse `json:"response,omitempty" xml:"response,omitempty"`
}
