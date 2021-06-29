package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
出库-ERP下发单(新接口，包含调拨出库单和退货出库单等) API返回值 
alibaba.wdk.ums.outbound

出库-ERP下发单(新接口，包含调拨出库单和退货出库单等)
*/
type AlibabaWdkUmsOutboundAPIResponse struct {
    model.CommonResponse
    AlibabaWdkUmsOutboundResponse
}

// 出库-ERP下发单(新接口，包含调拨出库单和退货出库单等) 成功返回结果
type AlibabaWdkUmsOutboundResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_ums_outbound_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *UtmsResult `json:"result,omitempty" xml:"result,omitempty"`
}
