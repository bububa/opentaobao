package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
退货包裹状态通知接口 APIResponse
taobao.qimen.returnpackage.report

退货包裹状态通知接口
*/
type TaobaoQimenReturnpackageReportAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_returnpackage_report_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *Response `json:"response,omitempty" xml:"