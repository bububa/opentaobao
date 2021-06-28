package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发货单缺货通知接口 APIResponse
taobao.qimen.itemlack.report

WMS调用奇门的接口,将商家在库某商品缺货的信息回传给ERP
*/
type TaobaoQimenItemlackReportAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_itemlack_report_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *Response `json:"response,omitempty" xml:"