package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdksupplierrefundget 五道口按订单号批量查询供应商退款单
// alibaba.wdk.supplier.refund.get
//
// 五道口按订单号批量查询供应商退款单
func Alibabawdksupplierrefundget(clt *core.SDKClient, req *wdk.AlibabawdksupplierrefundgetAPIRequest, session string) (*wdk.AlibabawdksupplierrefundgetAPIResponse, error) {
	var resp wdk.AlibabawdksupplierrefundgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
