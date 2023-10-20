package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangInventoryBatchQueryAPIResponse 批量查询库存 API返回值
// alibaba.dchain.aoxiang.inventory.batch.query
//
// 批量查询库存
type AlibabaDchainAoxiangInventoryBatchQueryAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangInventoryBatchQueryAPIResponseModel
}

// AlibabaDchainAoxiangInventoryBatchQueryAPIResponseModel is 批量查询库存 成功返回结果
type AlibabaDchainAoxiangInventoryBatchQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_inventory_batch_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	BatchQueryInventoryResponse *BatchQueryInventoryResponse `json:"batch_query_inventory_response,omitempty" xml:"batch_query_inventory_response,omitempty"`
}
