package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbacustomerssidget 查看功能权限
// taobao.simba.customers.sid.get
//
// 查询用户是否拥有某个功能权限
func Taobaosimbacustomerssidget(clt *core.SDKClient, req *simba.TaobaosimbacustomerssidgetAPIRequest, session string) (*simba.TaobaosimbacustomerssidgetAPIResponse, error) {
	var resp simba.TaobaosimbacustomerssidgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
