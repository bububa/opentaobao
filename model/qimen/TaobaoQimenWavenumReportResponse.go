package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发货单波次通知接口 APIResponse
taobao.qimen.wavenum.report

WMS调用奇门的接口,该接口用来给ERP或者OMS回传波次号及对应的发货单号，以支持商家货票同行、波次内包裹的处理等需求
*/
type TaobaoQimenWavenumReportAPIResponse struct {
    model.CommonResponse
    TaobaoQimenWavenumReportResponse
}

type TaobaoQimenWavenumReportResponse struct {
    XMLName xml.Name `xml:"qimen_wavenum_report_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *Response `json:"response,omitempty" xml:"response,omitempty"`

    
}
