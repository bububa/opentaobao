package tbk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-淘礼金发放及使用报表 APIResponse
taobao.tbk.dg.vegas.tlj.instance.report

淘礼金实例维度相关报表数据查询
*/
type TaobaoTbkDgVegasTljInstanceReportAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTbkDgVegasTljInstanceReportResponse `json:"taobao_tbk_dg_vegas_tlj_instance_report_response,omitempty"`
}

type TaobaoTbkDgVegasTljInstanceReportResponse struct {

    // 接口返回model
    Result   *TaobaoTbkDgVegasTljInstanceReportResult `json:"result,omitempty"`

}
