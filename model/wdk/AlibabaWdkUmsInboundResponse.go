package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
入库-ERP下发单 APIResponse
alibaba.wdk.ums.inbound

入库-ERP下发单
*/
type AlibabaWdkUmsInboundAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_ums_inbound_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *UtmsResult `json:"result,omitempty" xml:"