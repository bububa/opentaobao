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
    TaobaoQimenReturnpackageReportResponse
}

type TaobaoQimenReturnpackageReportResponse struct {
    XMLName xml.Name `xml:"qimen_returnpackage_report_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *Response `json:"response,omitempty" xml:"response,omitempty"`

    
}
