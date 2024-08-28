package qt

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qt"
)

// TaobaoQtReportDelete 质检报告删除接口
// taobao.qt.report.delete
//
// 删除质检报告
func TaobaoQtReportDelete(ctx context.Context, clt *core.SDKClient, req *qt.TaobaoQtReportDeleteAPIRequest, resp *qt.TaobaoQtReportDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
