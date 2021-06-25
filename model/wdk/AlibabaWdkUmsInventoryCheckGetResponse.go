package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
盘点结果单-回流单 APIResponse
alibaba.wdk.ums.inventory.check.get

盘点结果单-回流单
*/
type AlibabaWdkUmsInventoryCheckGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkUmsInventoryCheckGetResponse `json:"alibaba_wdk_ums_inventory_check_get_response,omitempty"`
}

type AlibabaWdkUmsInventoryCheckGetResponse struct {

    // result
    Result   *UtmsResult `json:"result,omitempty"`

}
