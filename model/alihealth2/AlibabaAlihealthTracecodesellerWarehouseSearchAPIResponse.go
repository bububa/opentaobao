package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTracecodesellerWarehouseSearchAPIResponse 查询仓库api API返回值
// alibaba.alihealth.tracecodeseller.warehouse.search
//
// 查询仓库api
type AlibabaAlihealthTracecodesellerWarehouseSearchAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthTracecodesellerWarehouseSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthTracecodesellerWarehouseSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthTracecodesellerWarehouseSearchAPIResponseModel).Reset()
}

// AlibabaAlihealthTracecodesellerWarehouseSearchAPIResponseModel is 查询仓库api 成功返回结果
type AlibabaAlihealthTracecodesellerWarehouseSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_tracecodeseller_warehouse_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthTracecodesellerWarehouseSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthTracecodesellerWarehouseSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthTracecodesellerWarehouseSearchAPIResponse)
	},
}

// GetAlibabaAlihealthTracecodesellerWarehouseSearchAPIResponse 从 sync.Pool 获取 AlibabaAlihealthTracecodesellerWarehouseSearchAPIResponse
func GetAlibabaAlihealthTracecodesellerWarehouseSearchAPIResponse() *AlibabaAlihealthTracecodesellerWarehouseSearchAPIResponse {
	return poolAlibabaAlihealthTracecodesellerWarehouseSearchAPIResponse.Get().(*AlibabaAlihealthTracecodesellerWarehouseSearchAPIResponse)
}

// ReleaseAlibabaAlihealthTracecodesellerWarehouseSearchAPIResponse 将 AlibabaAlihealthTracecodesellerWarehouseSearchAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthTracecodesellerWarehouseSearchAPIResponse(v *AlibabaAlihealthTracecodesellerWarehouseSearchAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthTracecodesellerWarehouseSearchAPIResponse.Put(v)
}
