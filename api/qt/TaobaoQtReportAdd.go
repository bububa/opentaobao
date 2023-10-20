package qt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qt"
)

// TaobaoQtReportAdd 上传质检报告
// taobao.qt.report.add
//
// 上传质检报告
func TaobaoQtReportAdd(clt *core.SDKClient, req *qt.TaobaoQtReportAddAPIRequest, resp *qt.TaobaoQtReportAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
