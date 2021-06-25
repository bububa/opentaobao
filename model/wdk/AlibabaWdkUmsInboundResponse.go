package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
入库-ERP下发单 APIResponse
alibaba.wdk.ums.inbound

入库-ERP下发单
*/
type AlibabaWdkUmsInboundAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkUmsInboundResponse `json:"alibaba_wdk_ums_inbound_response,omitempty"`
}

type AlibabaWdkUmsInboundResponse struct {

    // result
    Result   *UtmsResult `json:"result,omitempty"`

}
