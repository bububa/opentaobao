package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
入库-ERP下发单 API返回值 
alibaba.wdk.ums.inbound

入库-ERP下发单
*/
type AlibabaWdkUmsInboundAPIResponse struct {
    model.CommonResponse
    AlibabaWdkUmsInboundResponse
}

// 入库-ERP下发单 成功返回结果
type AlibabaWdkUmsInboundResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_ums_inbound_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *UtmsResult `json:"result,omitempty" xml:"result,omitempty"`
}
