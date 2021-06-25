package qt

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qt"
)

/* 
查询质检报告 
taobao.qt.report.get

质检报告查询
*/
func TaobaoQtReportGet(clt *core.SDKClient, req *qt.TaobaoQtReportGetRequest, session string) (*qt.TaobaoQtReportGetResponse, error) {
    var resp qt.TaobaoQtReportGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
