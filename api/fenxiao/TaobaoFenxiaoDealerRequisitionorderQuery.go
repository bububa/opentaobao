package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoDealerRequisitionorderQuery 按编号查询采购申请/经销采购单
// taobao.fenxiao.dealer.requisitionorder.query
//
// 按编号查询采购申请/经销采购单，目前支持供应商和分销商查询。
func TaobaoFenxiaoDealerRequisitionorderQuery(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoDealerRequisitionorderQueryAPIRequest, resp *fenxiao.TaobaoFenxiaoDealerRequisitionorderQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
