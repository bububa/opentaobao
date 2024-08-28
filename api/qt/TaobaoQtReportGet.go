package qt

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qt"
)

// TaobaoQtReportGet 查询质检报告
// taobao.qt.report.get
//
// 质检报告查询
func TaobaoQtReportGet(ctx context.Context, clt *core.SDKClient, req *qt.TaobaoQtReportGetAPIRequest, resp *qt.TaobaoQtReportGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
