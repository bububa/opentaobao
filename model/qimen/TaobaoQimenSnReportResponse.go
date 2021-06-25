package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
发货单SN通知接口 APIResponse
taobao.qimen.sn.report

WMS调用奇门的接口,在仓库出库单后, 把SN信息回传给ERP
*/
type TaobaoQimenSnReportAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenSnReportResponse `json:"taobao_qimen_sn_report_response,omitempty"`
}

type TaobaoQimenSnReportResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
