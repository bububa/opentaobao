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
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_transferorder_report_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *TaobaoQimenTransferorderReportStruct `json:"response,omitempty" xml:"