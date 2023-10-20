package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbacustomersauthorizedget 取得当前登录用户的授权账户列表
// taobao.simba.customers.authorized.get
//
// 取得当前登录用户的授权账户列表
func Taobaosimbacustomersauthorizedget(clt *core.SDKClient, req *simba.TaobaosimbacustomersauthorizedgetAPIRequest, session string) (*simba.TaobaosimbacustomersauthorizedgetAPIResponse, error) {
	var resp simba.TaobaosimbacustomersauthorizedgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
