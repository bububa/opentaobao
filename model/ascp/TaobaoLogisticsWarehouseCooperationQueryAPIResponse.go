package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticswarehousecooperationqueryAPIResponse 仓合作关系查询 API返回值
// taobao.logistics.warehouse.cooperation.query
//
// 仓合作关系查询
type TaobaologisticswarehousecooperationqueryAPIResponse struct {
	model.CommonResponse
	TaobaologisticswarehousecooperationqueryAPIResponseModel
}

// TaobaologisticswarehousecooperationqueryAPIResponseModel is 仓合作关系查询 成功返回结果
type TaobaologisticswarehousecooperationqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_warehouse_cooperation_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 仓合作关系查询出参
	WarehouseCooperationQueryResponse *WarehouseCooperationQueryResponse `json:"warehouse_cooperation_query_response,omitempty" xml:"warehouse_cooperation_query_response,omitempty"`
}
