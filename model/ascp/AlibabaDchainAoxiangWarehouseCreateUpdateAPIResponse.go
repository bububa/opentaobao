package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangWarehouseCreateUpdateAPIResponse 仓库信息同步 API返回值
// alibaba.dchain.aoxiang.warehouse.create.update
//
// 仓库信息同步
type AlibabaDchainAoxiangWarehouseCreateUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangWarehouseCreateUpdateAPIResponseModel
}

// AlibabaDchainAoxiangWarehouseCreateUpdateAPIResponseModel is 仓库信息同步 成功返回结果
type AlibabaDchainAoxiangWarehouseCreateUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_warehouse_create_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 处理结果
	WarehouseUpsertResponse *WarehouseUpsertResponse `json:"warehouse_upsert_response,omitempty" xml:"warehouse_upsert_response,omitempty"`
}
