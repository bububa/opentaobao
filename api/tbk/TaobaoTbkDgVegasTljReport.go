package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkDgVegasTljReport 淘宝客-推广者-淘礼金效果数据
// taobao.tbk.dg.vegas.tlj.report
//
// 淘宝客推广者可查询淘礼金发放和使用等效果数据，只提供近150天的数据
func TaobaoTbkDgVegasTljReport(clt *core.SDKClient, req *tbk.TaobaoTbkDgVegasTljReportAPIRequest, resp *tbk.TaobaoTbkDgVegasTljReportAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
