package qt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qt"
)

// TaobaoQtReportGet 查询质检报告
// taobao.qt.report.get
//
// 质检报告查询
func TaobaoQtReportGet(clt *core.SDKClient, req *qt.TaobaoQtReportGetAPIRequest, resp *qt.TaobaoQtReportGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
