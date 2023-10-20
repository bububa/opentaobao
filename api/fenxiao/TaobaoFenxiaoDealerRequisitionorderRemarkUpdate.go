package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoDealerRequisitionorderRemarkUpdate 修改经销采购单备注
// taobao.fenxiao.dealer.requisitionorder.remark.update
//
// 供应商修改经销采购单备注
func TaobaoFenxiaoDealerRequisitionorderRemarkUpdate(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIRequest, resp *fenxiao.TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
