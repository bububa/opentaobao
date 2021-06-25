package qt

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qt"
)

/* 
上传质检报告 
taobao.qt.report.add

上传质检报告
*/
func TaobaoQtReportAdd(clt *core.SDKClient, req *qt.TaobaoQtReportAddRequest, session string) (*qt.TaobaoQtReportAddResponse, error) {
    var resp qt.TaobaoQtReportAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
