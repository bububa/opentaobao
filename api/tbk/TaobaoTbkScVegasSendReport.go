package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkScVegasSendReport 淘宝客-服务商-查询红包发放个数
// taobao.tbk.sc.vegas.send.report
//
// 服务商使用。可通过此接口查询对应推广者账号下的红包发放个数。
func TaobaoTbkScVegasSendReport(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkScVegasSendReportAPIRequest, resp *tbk.TaobaoTbkScVegasSendReportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
