package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleUserPermitRevoke 删除服务商与卖家之间的订单消息绑定关系
// alibaba.idle.user.permit.revoke
//
// 删除服务商与卖家之间的订单消息绑定关系，删除后不再发送消息
func AlibabaIdleUserPermitRevoke(clt *core.SDKClient, req *idleisv.AlibabaIdleUserPermitRevokeAPIRequest, session string) (*idleisv.AlibabaIdleUserPermitRevokeAPIResponse, error) {
	var resp idleisv.AlibabaIdleUserPermitRevokeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
