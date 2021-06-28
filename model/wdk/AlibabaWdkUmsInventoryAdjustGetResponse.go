package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
库调单-回流单 APIResponse
alibaba.wdk.ums.inventory.adjust.get

库调单-回流单
*/
type AlibabaWdkUmsInventoryAdjustGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkUmsInventoryAdjustGetResponse `json:"alibaba_wdk_ums_inventory_adjust_get_response,omitempty"` 
    AlibabaWdkUmsInventoryAdjustGetResponse
}

/* model for simplify = false
type AlibabaWdkUmsInventoryAdjustGetResponse struct {

    // result
    
    Result  *struct {
        UtmsResult  *UtmsResult `json:"utms_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkUmsInventoryAdjustGetResponse struct {

    // result
    Result   *UtmsResult `json:"result,omitempty"`

}
