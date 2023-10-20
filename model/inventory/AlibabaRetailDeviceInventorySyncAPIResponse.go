package inventory

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailDeviceInventorySyncAPIResponse 库存同步接口 API返回值
// alibaba.retail.device.inventory.sync
//
// 商库存同步接口
type AlibabaRetailDeviceInventorySyncAPIResponse struct {
	model.CommonResponse
	AlibabaRetailDeviceInventorySyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailDeviceInventorySyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailDeviceInventorySyncAPIResponseModel).Reset()
}

// AlibabaRetailDeviceInventorySyncAPIResponseModel is 库存同步接口 成功返回结果
type AlibabaRetailDeviceInventorySyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_device_inventory_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaRetailDeviceInventorySyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailDeviceInventorySyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaRetailDeviceInventorySyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailDeviceInventorySyncAPIResponse)
	},
}

// GetAlibabaRetailDeviceInventorySyncAPIResponse 从 sync.Pool 获取 AlibabaRetailDeviceInventorySyncAPIResponse
func GetAlibabaRetailDeviceInventorySyncAPIResponse() *AlibabaRetailDeviceInventorySyncAPIResponse {
	return poolAlibabaRetailDeviceInventorySyncAPIResponse.Get().(*AlibabaRetailDeviceInventorySyncAPIResponse)
}

// ReleaseAlibabaRetailDeviceInventorySyncAPIResponse 将 AlibabaRetailDeviceInventorySyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailDeviceInventorySyncAPIResponse(v *AlibabaRetailDeviceInventorySyncAPIResponse) {
	v.Reset()
	poolAlibabaRetailDeviceInventorySyncAPIResponse.Put(v)
}
