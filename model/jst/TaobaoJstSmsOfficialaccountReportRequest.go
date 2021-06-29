package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔公众号信息上报 API请求
taobao.jst.sms.officialaccount.report

聚石塔公众号信息上报
*/
type TaobaoJstSmsOfficialaccountReportRequest struct {
    model.Params
    // 公众号信息上报接口入参
    _officialAccountInfoReportRequest   *OfficialAccountInfoReportRequest
}

// 初始化TaobaoJstSmsOfficialaccountReportRequest对象
func NewTaobaoJstSmsOfficialaccountReportRequest() *TaobaoJstSmsOfficialaccountReportRequest{
    return &TaobaoJstSmsOfficialaccountReportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsOfficialaccountReportRequest) GetApiMethodName() string {
    return "taobao.jst.sms.officialaccount.report"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsOfficialaccountReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OfficialAccountInfoReportRequest Setter
// 公众号信息上报接口入参
func (r *TaobaoJstSmsOfficialaccountReportRequest) SetOfficialAccountInfoReportRequest(_officialAccountInfoReportRequest *OfficialAccountInfoReportRequest) error {
    r._officialAccountInfoReportRequest = _officialAccountInfoReportRequest
    r.Set("official_account_info_report_request", _officialAccountInfoReportRequest)
    return nil
}

// OfficialAccountInfoReportRequest Getter
func (r TaobaoJstSmsOfficialaccountReportRequest) GetOfficialAccountInfoReportRequest() *OfficialAccountInfoReportRequest {
    return r._officialAccountInfoReportRequest
}
