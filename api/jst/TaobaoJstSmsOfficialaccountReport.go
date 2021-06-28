package jst

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jst"
)

/* 
聚石塔公众号信息上报 
taobao.jst.sms.officialaccount.report

聚石塔公众号信息上报
*/
func TaobaoJstSmsOfficialaccountReport(clt *core.SDKClient, req *jst.TaobaoJstSmsOfficialaccountReportRequest, session string) (*jst.TaobaoJstSmsOfficialaccountReportAPIResponse, error) {
    var resp jst.TaobaoJstSmsOfficialaccountReportAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
