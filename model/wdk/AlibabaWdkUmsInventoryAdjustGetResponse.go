package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
库调单-回流单 APIResponse
alibaba.wdk.ums.inventory.adjust.get

库调单-回流单
*/
type AlibabaWdkUmsInventoryAdjustGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_ums_inventory_adjust_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *UtmsResult `json:"result,omitempty" xml:"