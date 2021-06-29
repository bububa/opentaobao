package inventory

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
库存同步接口 APIResponse
alibaba.retail.device.inventory.sync

商库存同步接口
*/
type AlibabaRetailDeviceInventorySyncAPIResponse struct {
    model.CommonResponse
    AlibabaRetailDeviceInventorySyncResponse
}

type AlibabaRetailDeviceInventorySyncResponse struct {
    XMLName xml.Name `xml:"alibaba_retail_device_inventory_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaRetailDeviceInventorySyncResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
