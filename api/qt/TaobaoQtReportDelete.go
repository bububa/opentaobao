package qt

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qt"
)

/* 
质检报告删除接口 
taobao.qt.report.delete

删除质检报告
*/
func TaobaoQtReportDelete(clt *core.SDKClient, req *qt.TaobaoQtReportDeleteRequest, session string) (*qt.TaobaoQtReportDeleteAPIResponse, error) {
    var resp qt.TaobaoQtReportDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
