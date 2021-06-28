package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
调拨单通知 APIResponse
taobao.qimen.transferorder.report

调拨单通知
*/
type TaobaoQimenTransferorderReportAPIResponse struct {
    model.CommonResponse
    TaobaoQimenTransferorderReportResponse
}

type TaobaoQimenTransferorderReportResponse struct {
    XMLName xml.Name `xml:"qimen_transferorder_report_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *TaobaoQimenTransferorderReportStruct `json:"response,omitempty" xml:"response,omitempty"`

    
}
