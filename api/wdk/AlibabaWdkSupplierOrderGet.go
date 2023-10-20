package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdksupplierorderget 五道口按订单号批量查询供应商正向订单
// alibaba.wdk.supplier.order.get
//
// 五道口按订单号批量查询供应商正向订单
func Alibabawdksupplierorderget(clt *core.SDKClient, req *wdk.AlibabawdksupplierordergetAPIRequest, session string) (*wdk.AlibabawdksupplierordergetAPIResponse, error) {
	var resp wdk.AlibabawdksupplierordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
