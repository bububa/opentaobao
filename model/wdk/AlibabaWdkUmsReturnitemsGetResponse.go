package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
退货库位商品查询（退货出库辅助）-回流单 APIResponse
alibaba.wdk.ums.returnitems.get

退货库位商品查询（退货出库辅助）-回流单
*/
type AlibabaWdkUmsReturnitemsGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkUmsReturnitemsGetResponse `json:"alibaba_wdk_ums_returnitems_get_response,omitempty"` 
    AlibabaWdkUmsReturnitemsGetResponse
}

/* model for simplify = false
type AlibabaWdkUmsReturnitemsGetResponse struct {

    // reslut
    
    Result  *struct {
        UtmsResult  *UtmsResult `json:"utms_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkUmsReturnitemsGetResponse struct {

    // reslut
    Result   *UtmsResult `json:"result,omitempty"`

}
