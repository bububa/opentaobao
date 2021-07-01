package qt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qt"
)

/* TaobaoQtReportAdd
上传质检报告
taobao.qt.report.add

上传质检报告 */
func TaobaoQtReportAdd(clt *core.SDKClient, req *qt.TaobaoQtReportAddAPIRequest, session string) (*qt.TaobaoQtReportAddAPIResponse, error) {
	var resp qt.TaobaoQtReportAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
