package xiamitrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xiamitrace"
)

/* 
曲库开放平台内容行为上报接口 
xiami.content.resource.action.report

合作方对接入的曲库开放内容上报行为日志
*/
func XiamiContentResourceActionReport(clt *core.SDKClient, req *xiamitrace.XiamiContentResourceActionReportAPIRequest, session string) (*xiamitrace.XiamiContentResourceActionReportAPIResponse, error) {
    var resp xiamitrace.XiamiContentResourceActionReportAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
