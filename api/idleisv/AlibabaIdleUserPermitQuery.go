package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// Alibabaidleuserpermitquery 查询服务商与卖家之间的订单消息绑定关系
// alibaba.idle.user.permit.query
//
// 查询服务商与卖家之间的订单消息绑定关系
func Alibabaidleuserpermitquery(clt *core.SDKClient, req *idleisv.AlibabaidleuserpermitqueryAPIRequest, session string) (*idleisv.AlibabaidleuserpermitqueryAPIResponse, error) {
	var resp idleisv.AlibabaidleuserpermitqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
