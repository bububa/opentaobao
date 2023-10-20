package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdksupplierrefundlist 五道口按供应商拉取退款单
// alibaba.wdk.supplier.refund.list
//
// 五道口按供应商拉取退款单
func Alibabawdksupplierrefundlist(clt *core.SDKClient, req *wdk.AlibabawdksupplierrefundlistAPIRequest, session string) (*wdk.AlibabawdksupplierrefundlistAPIResponse, error) {
	var resp wdk.AlibabawdksupplierrefundlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
