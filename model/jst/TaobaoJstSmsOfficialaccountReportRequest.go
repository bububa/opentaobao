package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔公众号信息上报 APIRequest
taobao.jst.sms.officialaccount.report

聚石塔公众号信息上报
*/
type TaobaoJstSmsOfficialaccountReportRequest struct {
    model.Params

    // 公众号信息上报接口入参
    officialAccountInfoReportRequest   *OfficialAccountInfoReportRequest 

}

func NewTaobaoJstSmsOfficialaccountReportRequest() *TaobaoJstSmsOfficialaccountReportRequest{
    return &TaobaoJstSmsOfficialaccountReportRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstSmsOfficialaccountReportRequest) GetApiMethodName() string {
    return "taobao.jst.sms.officialaccount.report"
}

func (r TaobaoJstSmsOfficialaccountReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJstSmsOfficialaccountReportRequest) SetOfficialAccountInfoReportRequest(officialAccountInfoReportRequest *OfficialAccountInfoReportRequest) error {
    r.officialAccountInfoReportRequest = officialAccountInfoReportRequest
    r.Set("official_account_info_report_request", officialAccountInfoReportRequest)
    return nil
}

func (r TaobaoJstSmsOfficialaccountReportRequest) GetOfficialAccountInfoReportRequest() *OfficialAccountInfoReportRequest {
    return r.officialAccountInfoReportRequest
}

