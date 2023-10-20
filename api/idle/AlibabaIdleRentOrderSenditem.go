package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidlerentordersenditem 确认发货
// alibaba.idle.rent.order.senditem
//
// 确认发货
func Alibabaidlerentordersenditem(clt *core.SDKClient, req *idle.AlibabaidlerentordersenditemAPIRequest, session string) (*idle.AlibabaidlerentordersenditemAPIResponse, error) {
	var resp idle.AlibabaidlerentordersenditemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
