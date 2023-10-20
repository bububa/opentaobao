package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabatclsaelophyorderreceiptquery 订单小票查询
// alibaba.tcls.aelophy.order.receipt.query
//
// 订单小票查询
func Alibabatclsaelophyorderreceiptquery(clt *core.SDKClient, req *wdk.AlibabatclsaelophyorderreceiptqueryAPIRequest, session string) (*wdk.AlibabatclsaelophyorderreceiptqueryAPIResponse, error) {
	var resp wdk.AlibabatclsaelophyorderreceiptqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
