package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
发货单缺货通知接口 APIResponse
taobao.qimen.itemlack.report

WMS调用奇门的接口,将商家在库某商品缺货的信息回传给ERP
*/
type TaobaoQimenItemlackReportAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenItemlackReportResponse `json:"taobao_qimen_itemlack_report_response,omitempty"`
}

type TaobaoQimenItemlackReportResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
