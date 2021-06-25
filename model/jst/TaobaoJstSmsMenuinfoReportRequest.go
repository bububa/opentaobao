package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔菜单信息上报 APIRequest
taobao.jst.sms.menuinfo.report

聚石塔菜单信息上报
*/
type TaobaoJstSmsMenuinfoReportRequest struct {
    model.Params

    // 菜单信息上报接口的请求参数
    menuInfoReportRequest   *MenuInfoReportRequest 

}

func NewTaobaoJstSmsMenuinfoReportRequest() *TaobaoJstSmsMenuinfoReportRequest{
    return &TaobaoJstSmsMenuinfoReportRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstSmsMenuinfoReportRequest) GetApiMethodName() string {
    return "taobao.jst.sms.menuinfo.report"
}

func (r TaobaoJstSmsMenuinfoReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJstSmsMenuinfoReportRequest) SetMenuInfoReportRequest(menuInfoReportRequest *MenuInfoReportRequest) error {
    r.menuInfoReportRequest = menuInfoReportRequest
    r.Set("menu_info_report_request", menuInfoReportRequest)
    return nil
}

func (r TaobaoJstSmsMenuinfoReportRequest) GetMenuInfoReportRequest() *MenuInfoReportRequest {
    return r.menuInfoReportRequest
}

