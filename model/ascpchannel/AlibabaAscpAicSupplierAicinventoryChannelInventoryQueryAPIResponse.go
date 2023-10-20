package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIResponse 供应商直发-商家仓库存查询服务 API返回值
// alibaba.ascp.aic.supplier.aicinventory.channel.inventory.query
//
// 提供商家基于货品、供应商、仓，查询ascp 实时商家仓库存查询数据。
type AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIResponseModel).Reset()
}

// AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIResponseModel is 供应商直发-商家仓库存查询服务 成功返回结果
type AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_aic_supplier_aicinventory_channel_inventory_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	MerchantInventoryQueryResponse *ResultWrapper `json:"merchant_inventory_query_response,omitempty" xml:"merchant_inventory_query_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MerchantInventoryQueryResponse = nil
}

var poolAlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIResponse)
	},
}

// GetAlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIResponse 从 sync.Pool 获取 AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIResponse
func GetAlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIResponse() *AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIResponse {
	return poolAlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIResponse.Get().(*AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIResponse)
}

// ReleaseAlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIResponse 将 AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIResponse(v *AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIResponse) {
	v.Reset()
	poolAlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIResponse.Put(v)
}
