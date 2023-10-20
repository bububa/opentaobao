package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpGlobalSupplierItemListInfoQueryAPIResponse 供应商供品信息查询 API返回值
// alibaba.ascp.global.supplier.item.list.info.query
//
// 供应商供品信息查询
type AlibabaAscpGlobalSupplierItemListInfoQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAscpGlobalSupplierItemListInfoQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpGlobalSupplierItemListInfoQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpGlobalSupplierItemListInfoQueryAPIResponseModel).Reset()
}

// AlibabaAscpGlobalSupplierItemListInfoQueryAPIResponseModel is 供应商供品信息查询 成功返回结果
type AlibabaAscpGlobalSupplierItemListInfoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_global_supplier_item_list_info_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpGlobalSupplierItemListInfoQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpGlobalSupplierItemListInfoQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpGlobalSupplierItemListInfoQueryAPIResponse)
	},
}

// GetAlibabaAscpGlobalSupplierItemListInfoQueryAPIResponse 从 sync.Pool 获取 AlibabaAscpGlobalSupplierItemListInfoQueryAPIResponse
func GetAlibabaAscpGlobalSupplierItemListInfoQueryAPIResponse() *AlibabaAscpGlobalSupplierItemListInfoQueryAPIResponse {
	return poolAlibabaAscpGlobalSupplierItemListInfoQueryAPIResponse.Get().(*AlibabaAscpGlobalSupplierItemListInfoQueryAPIResponse)
}

// ReleaseAlibabaAscpGlobalSupplierItemListInfoQueryAPIResponse 将 AlibabaAscpGlobalSupplierItemListInfoQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpGlobalSupplierItemListInfoQueryAPIResponse(v *AlibabaAscpGlobalSupplierItemListInfoQueryAPIResponse) {
	v.Reset()
	poolAlibabaAscpGlobalSupplierItemListInfoQueryAPIResponse.Put(v)
}
