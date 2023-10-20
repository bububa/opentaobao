package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoOrderRemarkUpdate 修改采购单备注
// taobao.fenxiao.order.remark.update
//
// 供应商修改采购单备注
func TaobaoFenxiaoOrderRemarkUpdate(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoOrderRemarkUpdateAPIRequest, resp *fenxiao.TaobaoFenxiaoOrderRemarkUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
