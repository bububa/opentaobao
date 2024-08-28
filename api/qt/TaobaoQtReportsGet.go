package qt

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qt"
)

// TaobaoQtReportsGet 批量查询质检报告
// taobao.qt.reports.get
//
// 批量查询质检报告，目前只支持查询qtType=11（天猫真假鉴定）类型的报告
func TaobaoQtReportsGet(ctx context.Context, clt *core.SDKClient, req *qt.TaobaoQtReportsGetAPIRequest, resp *qt.TaobaoQtReportsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
