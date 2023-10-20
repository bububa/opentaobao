package inventory

import (
	"encoding/xml"

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

// AlibabaRetailDeviceInventorySyncAPIResponseModel is 库存同步接口 成功返回结果
type AlibabaRetailDeviceInventorySyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_device_inventory_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaRetailDeviceInventorySyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
