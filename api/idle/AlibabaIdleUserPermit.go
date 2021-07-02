package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleUserPermit 用户appkey授权
// alibaba.idle.user.permit
//
// 用于记录登录用户与服务商的绑定关系，用于业务数据分发和授权校验
func AlibabaIdleUserPermit(clt *core.SDKClient, req *idle.AlibabaIdleUserPermitAPIRequest, session string) (*idle.AlibabaIdleUserPermitAPIResponse, error) {
	var resp idle.AlibabaIdleUserPermitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
