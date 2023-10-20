package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoYphOrderGet 一盘货商家单个查询采购单信息
// taobao.fenxiao.yph.order.get
//
// 一盘货商家单个查询采购单信息
func TaobaoFenxiaoYphOrderGet(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoYphOrderGetAPIRequest, resp *fenxiao.TaobaoFenxiaoYphOrderGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
