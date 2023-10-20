package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidlerentorderreceiveitem 确认签收
// alibaba.idle.rent.order.receiveitem
//
// 确认揽收/签收
func Alibabaidlerentorderreceiveitem(clt *core.SDKClient, req *idle.AlibabaidlerentorderreceiveitemAPIRequest, session string) (*idle.AlibabaidlerentorderreceiveitemAPIResponse, error) {
	var resp idle.AlibabaidlerentorderreceiveitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
