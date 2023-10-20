package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticswarehouseoperationupdateAPIResponse 仓作业能力新建/更新 API返回值
// taobao.logistics.warehouse.operation.update
//
// 仓作业能力新建/更新
type TaobaologisticswarehouseoperationupdateAPIResponse struct {
	model.CommonResponse
	TaobaologisticswarehouseoperationupdateAPIResponseModel
}

// TaobaologisticswarehouseoperationupdateAPIResponseModel is 仓作业能力新建/更新 成功返回结果
type TaobaologisticswarehouseoperationupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_warehouse_operation_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 仓作业能力新建/更新出参
	WarehouseOperationUpdateResponse *WarehouseOperationUpdateResponse `json:"warehouse_operation_update_response,omitempty" xml:"warehouse_operation_update_response,omitempty"`
}
