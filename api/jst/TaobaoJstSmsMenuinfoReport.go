package jst

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jst"
)

/* 
聚石塔菜单信息上报 
taobao.jst.sms.menuinfo.report

聚石塔菜单信息上报
*/
func TaobaoJstSmsMenuinfoReport(clt *core.SDKClient, req *jst.TaobaoJstSmsMenuinfoReportAPIRequest, session string) (*jst.TaobaoJstSmsMenuinfoReportAPIResponse, error) {
    var resp jst.TaobaoJstSmsMenuinfoReportAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
