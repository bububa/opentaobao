package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleIsvUserQuery 服务商ISV闲鱼用户信息查询
// alibaba.idle.isv.user.query
//
// 服务商ISV闲鱼用户信息查询
func AlibabaIdleIsvUserQuery(clt *core.SDKClient, req *idleisv.AlibabaIdleIsvUserQueryAPIRequest, session string) (*idleisv.AlibabaIdleIsvUserQueryAPIResponse, error) {
	var resp idleisv.AlibabaIdleIsvUserQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
