package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
退货库位商品查询（退货出库辅助）-回流单 APIResponse
alibaba.wdk.ums.returnitems.get

退货库位商品查询（退货出库辅助）-回流单
*/
type AlibabaWdkUmsReturnitemsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_ums_returnitems_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // reslut
    
    Result   *UtmsResult `json:"result,omitempty" xml:"