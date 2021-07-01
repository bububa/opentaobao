package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoInventoryWarehouseQueryAPIResponse
分页查询商家仓信息 API返回值
taobao.inventory.warehouse.query

分页查询商家仓信息 */
type TaobaoInventoryWarehouseQueryAPIResponse struct {
	model.CommonResponse
	TaobaoInventoryWarehouseQueryAPIResponseModel
}

// TaobaoInventoryWarehouseQueryAPIResponseModel is 分页查询商家仓信息 成功返回结果
type TaobaoInventoryWarehouseQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"inventory_warehouse_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PaginationResult `json:"result,omitempty" xml:"result,omitempty"`
}
