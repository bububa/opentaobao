package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaorefundquery 批量查询采购退款
// taobao.fenxiao.refund.query
//
// 供应商按查询条件批量查询代销采购退款
func Taobaofenxiaorefundquery(clt *core.SDKClient, req *fenxiao.TaobaofenxiaorefundqueryAPIRequest, session string) (*fenxiao.TaobaofenxiaorefundqueryAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaorefundqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
