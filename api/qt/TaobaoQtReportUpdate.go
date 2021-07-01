package qt

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qt"
)

/* 
更新质检报告 
taobao.qt.report.update

更新质检报告
*/
func TaobaoQtReportUpdate(clt *core.SDKClient, req *qt.TaobaoQtReportUpdateAPIRequest, session string) (*qt.TaobaoQtReportUpdateAPIResponse, error) {
    var resp qt.TaobaoQtReportUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
