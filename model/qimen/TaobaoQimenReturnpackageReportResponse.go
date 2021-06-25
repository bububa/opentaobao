package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
退货包裹状态通知接口 APIResponse
taobao.qimen.returnpackage.report

退货包裹状态通知接口
*/
type TaobaoQimenReturnpackageReportAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenReturnpackageReportResponse `json:"taobao_qimen_returnpackage_report_response,omitempty"`
}

type TaobaoQimenReturnpackageReportResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
