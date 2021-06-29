package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发货单波次通知接口 API返回值 
taobao.qimen.wavenum.report

WMS调用奇门的接口,该接口用来给ERP或者OMS回传波次号及对应的发货单号，以支持商家货票同行、波次内包裹的处理等需求
*/
type TaobaoQimenWavenumReportAPIResponse struct {
    model.CommonResponse
    TaobaoQimenWavenumReportResponse
}

// 发货单波次通知接口 成功返回结果
type TaobaoQimenWavenumReportResponse struct {
    XMLName xml.Name `xml:"qimen_wavenum_report_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 
    Response   *Response `json:"response,omitempty" xml:"response,omitempty"`
}
