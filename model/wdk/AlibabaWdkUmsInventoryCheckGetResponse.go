package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
盘点结果单-回流单 API返回值 
alibaba.wdk.ums.inventory.check.get

盘点结果单-回流单
*/
type AlibabaWdkUmsInventoryCheckGetAPIResponse struct {
    model.CommonResponse
    AlibabaWdkUmsInventoryCheckGetResponse
}

// 盘点结果单-回流单 成功返回结果
type AlibabaWdkUmsInventoryCheckGetResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_ums_inventory_check_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *UtmsResult `json:"result,omitempty" xml:"result,omitempty"`
}
