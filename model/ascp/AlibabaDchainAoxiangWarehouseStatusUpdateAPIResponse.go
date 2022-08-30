package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangWarehouseStatusUpdateAPIResponse 启用/停用仓资源 API返回值
// alibaba.dchain.aoxiang.warehouse.status.update
//
// 启用/停用仓资源
type AlibabaDchainAoxiangWarehouseStatusUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangWarehouseStatusUpdateAPIResponseModel
}

// AlibabaDchainAoxiangWarehouseStatusUpdateAPIResponseModel is 启用/停用仓资源 成功返回结果
type AlibabaDchainAoxiangWarehouseStatusUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_warehouse_status_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 启用/停用仓资源出参
	WarehouseStatusUpdateResponse *WarehouseStatusUpdateResponse `json:"warehouse_status_update_response,omitempty" xml:"warehouse_status_update_response,omitempty"`
}
