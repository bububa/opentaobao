package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔菜单信息上报 API请求
taobao.jst.sms.menuinfo.report

聚石塔菜单信息上报
*/
type TaobaoJstSmsMenuinfoReportRequest struct {
    model.Params
    // 菜单信息上报接口的请求参数
    menuInfoReportRequest   *MenuInfoReportRequest
}

// 初始化TaobaoJstSmsMenuinfoReportRequest对象
func NewTaobaoJstSmsMenuinfoReportRequest() *TaobaoJstSmsMenuinfoReportRequest{
    return &TaobaoJstSmsMenuinfoReportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsMenuinfoReportRequest) GetApiMethodName() string {
    return "taobao.jst.sms.menuinfo.report"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsMenuinfoReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MenuInfoReportRequest Setter
// 菜单信息上报接口的请求参数
func (r *TaobaoJstSmsMenuinfoReportRequest) SetMenuInfoReportRequest(menuInfoReportRequest *MenuInfoReportRequest) error {
    r.menuInfoReportRequest = menuInfoReportRequest
    r.Set("menu_info_report_request", menuInfoReportRequest)
    return nil
}

// MenuInfoReportRequest Getter
func (r TaobaoJstSmsMenuinfoReportRequest) GetMenuInfoReportRequest() *MenuInfoReportRequest {
    return r.menuInfoReportRequest
}
