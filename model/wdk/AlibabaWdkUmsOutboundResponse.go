package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
出库-ERP下发单(新接口，包含调拨出库单和退货出库单等) APIResponse
alibaba.wdk.ums.outbound

出库-ERP下发单(新接口，包含调拨出库单和退货出库单等)
*/
type AlibabaWdkUmsOutboundAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkUmsOutboundResponse `json:"alibaba_wdk_ums_outbound_response,omitempty"`
}

type AlibabaWdkUmsOutboundResponse struct {

    // result
    Result   *UtmsResult `json:"result,omitempty"`

}
