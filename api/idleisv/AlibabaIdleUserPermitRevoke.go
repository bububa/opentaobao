package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// Alibabaidleuserpermitrevoke 删除服务商与卖家之间的订单消息绑定关系
// alibaba.idle.user.permit.revoke
//
// 删除服务商与卖家之间的订单消息绑定关系，删除后不再发送消息
func Alibabaidleuserpermitrevoke(clt *core.SDKClient, req *idleisv.AlibabaidleuserpermitrevokeAPIRequest, session string) (*idleisv.AlibabaidleuserpermitrevokeAPIResponse, error) {
	var resp idleisv.AlibabaidleuserpermitrevokeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
