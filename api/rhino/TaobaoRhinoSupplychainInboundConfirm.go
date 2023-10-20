package rhino

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/rhino"
)

// Taobaorhinosupplychaininboundconfirm WMS003成衣入库确认
// taobao.rhino.supplychain.inbound.confirm
//
// 【WMS003】【同步成衣入库完成信息】
func Taobaorhinosupplychaininboundconfirm(clt *core.SDKClient, req *rhino.TaobaorhinosupplychaininboundconfirmAPIRequest, session string) (*rhino.TaobaorhinosupplychaininboundconfirmAPIResponse, error) {
	var resp rhino.TaobaorhinosupplychaininboundconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
