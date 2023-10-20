package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleUserPermitQuery 查询服务商与卖家之间的订单消息绑定关系
// alibaba.idle.user.permit.query
//
// 查询服务商与卖家之间的订单消息绑定关系
func AlibabaIdleUserPermitQuery(clt *core.SDKClient, req *idleisv.AlibabaIdleUserPermitQueryAPIRequest, session string) (*idleisv.AlibabaIdleUserPermitQueryAPIResponse, error) {
	var resp idleisv.AlibabaIdleUserPermitQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
