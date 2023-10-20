package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkUmsInventoryCheckGetAPIResponse 盘点结果单-回流单 API返回值
// alibaba.wdk.ums.inventory.check.get
//
// 盘点结果单-回流单
type AlibabaWdkUmsInventoryCheckGetAPIResponse struct {
	model.CommonResponse
	AlibabaWdkUmsInventoryCheckGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkUmsInventoryCheckGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkUmsInventoryCheckGetAPIResponseModel).Reset()
}

// AlibabaWdkUmsInventoryCheckGetAPIResponseModel is 盘点结果单-回流单 成功返回结果
type AlibabaWdkUmsInventoryCheckGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_ums_inventory_check_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *UtmsResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkUmsInventoryCheckGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkUmsInventoryCheckGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkUmsInventoryCheckGetAPIResponse)
	},
}

// GetAlibabaWdkUmsInventoryCheckGetAPIResponse 从 sync.Pool 获取 AlibabaWdkUmsInventoryCheckGetAPIResponse
func GetAlibabaWdkUmsInventoryCheckGetAPIResponse() *AlibabaWdkUmsInventoryCheckGetAPIResponse {
	return poolAlibabaWdkUmsInventoryCheckGetAPIResponse.Get().(*AlibabaWdkUmsInventoryCheckGetAPIResponse)
}

// ReleaseAlibabaWdkUmsInventoryCheckGetAPIResponse 将 AlibabaWdkUmsInventoryCheckGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkUmsInventoryCheckGetAPIResponse(v *AlibabaWdkUmsInventoryCheckGetAPIResponse) {
	v.Reset()
	poolAlibabaWdkUmsInventoryCheckGetAPIResponse.Put(v)
}
