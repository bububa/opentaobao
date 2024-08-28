package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkDgVegasSendReport 淘宝客-推广者-查询红包发放个数
// taobao.tbk.dg.vegas.send.report
//
// 查询账号下的红包发放个数。
func TaobaoTbkDgVegasSendReport(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkDgVegasSendReportAPIRequest, resp *tbk.TaobaoTbkDgVegasSendReportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
