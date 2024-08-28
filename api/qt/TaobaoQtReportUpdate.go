package qt

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qt"
)

// TaobaoQtReportUpdate 更新质检报告
// taobao.qt.report.update
//
// 更新质检报告
func TaobaoQtReportUpdate(ctx context.Context, clt *core.SDKClient, req *qt.TaobaoQtReportUpdateAPIRequest, resp *qt.TaobaoQtReportUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
