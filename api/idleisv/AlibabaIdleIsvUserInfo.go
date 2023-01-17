package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleIsvUserInfo 闲鱼用户信息查询接口
// alibaba.idle.isv.user.info
//
// 闲鱼用户信息查询接口
func AlibabaIdleIsvUserInfo(clt *core.SDKClient, req *idleisv.AlibabaIdleIsvUserInfoAPIRequest, session string) (*idleisv.AlibabaIdleIsvUserInfoAPIResponse, error) {
	var resp idleisv.AlibabaIdleIsvUserInfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
