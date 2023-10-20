package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkDgVegasSendReport 淘宝客-推广者-查询红包发放个数
// taobao.tbk.dg.vegas.send.report
//
// 查询账号下的红包发放个数。
func TaobaoTbkDgVegasSendReport(clt *core.SDKClient, req *tbk.TaobaoTbkDgVegasSendReportAPIRequest, resp *tbk.TaobaoTbkDgVegasSendReportAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
