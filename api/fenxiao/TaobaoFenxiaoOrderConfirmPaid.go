package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaoorderconfirmpaid 确认收款
// taobao.fenxiao.order.confirm.paid
//
// 供应商确认收款（非支付宝交易）。
func Taobaofenxiaoorderconfirmpaid(clt *core.SDKClient, req *fenxiao.TaobaofenxiaoorderconfirmpaidAPIRequest, session string) (*fenxiao.TaobaofenxiaoorderconfirmpaidAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaoorderconfirmpaidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
