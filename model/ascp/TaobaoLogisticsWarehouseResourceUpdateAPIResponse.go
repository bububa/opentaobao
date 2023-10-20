package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticswarehouseresourceupdateAPIResponse 服务商新建/更新仓 API返回值
// taobao.logistics.warehouse.resource.update
//
// 服务商新建/更新仓
type TaobaologisticswarehouseresourceupdateAPIResponse struct {
	model.CommonResponse
	TaobaologisticswarehouseresourceupdateAPIResponseModel
}

// TaobaologisticswarehouseresourceupdateAPIResponseModel is 服务商新建/更新仓 成功返回结果
type TaobaologisticswarehouseresourceupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_warehouse_resource_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建/更新仓出参
	WarehouseUpdateResponse *WarehouseUpdateResponse `json:"warehouse_update_response,omitempty" xml:"warehouse_update_response,omitempty"`
}
