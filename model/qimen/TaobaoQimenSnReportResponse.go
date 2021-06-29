package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发货单SN通知接口 API返回值 
taobao.qimen.sn.report

WMS调用奇门的接口,在仓库出库单后, 把SN信息回传给ERP
*/
type TaobaoQimenSnReportAPIResponse struct {
    model.CommonResponse
    TaobaoQimenSnReportResponse
}

// 发货单SN通知接口 成功返回结果
type TaobaoQimenSnReportResponse struct {
    XMLName xml.Name `xml:"qimen_sn_report_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 
    Response   *Response `json:"response,omitempty" xml:"response,omitempty"`
}
