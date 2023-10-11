package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsWarehouseCooperationBatchConfirmAPIResponse 仓合作关系确认 API返回值
// taobao.logistics.warehouse.cooperation.batch.confirm
//
// 仓合作关系确认
type TaobaoLogisticsWarehouseCooperationBatchConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsWarehouseCooperationBatchConfirmAPIResponseModel
}

// TaobaoLogisticsWarehouseCooperationBatchConfirmAPIResponseModel is 仓合作关系确认 成功返回结果
type TaobaoLogisticsWarehouseCooperationBatchConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_warehouse_cooperation_batch_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 仓合作关系确认出参
	WarehouseCooperationBatchConfirmResponse *WarehouseCooperationBatchConfirmResponse `json:"warehouse_cooperation_batch_confirm_response,omitempty" xml:"warehouse_cooperation_batch_confirm_response,omitempty"`
}
