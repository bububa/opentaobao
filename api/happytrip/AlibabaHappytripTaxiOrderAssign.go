package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// Alibabahappytriptaxiorderassign 订单指派
// alibaba.happytrip.taxi.order.assign
//
// 通知供应商订单指派成功
func Alibabahappytriptaxiorderassign(clt *core.SDKClient, req *happytrip.AlibabahappytriptaxiorderassignAPIRequest, session string) (*happytrip.AlibabahappytriptaxiorderassignAPIResponse, error) {
	var resp happytrip.AlibabahappytriptaxiorderassignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
