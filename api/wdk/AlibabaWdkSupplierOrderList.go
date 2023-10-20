package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdksupplierorderlist 五道口供应商维度正向订单拉取
// alibaba.wdk.supplier.order.list
//
// 五道口供应商维度正向订单拉取
func Alibabawdksupplierorderlist(clt *core.SDKClient, req *wdk.AlibabawdksupplierorderlistAPIRequest, session string) (*wdk.AlibabawdksupplierorderlistAPIResponse, error) {
	var resp wdk.AlibabawdksupplierorderlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
