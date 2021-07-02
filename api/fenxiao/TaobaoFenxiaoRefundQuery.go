package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoRefundQuery 批量查询采购退款
// taobao.fenxiao.refund.query
//
// 供应商按查询条件批量查询代销采购退款
func TaobaoFenxiaoRefundQuery(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoRefundQueryAPIRequest, session string) (*fenxiao.TaobaoFenxiaoRefundQueryAPIResponse, error) {
	var resp fenxiao.TaobaoFenxiaoRefundQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
