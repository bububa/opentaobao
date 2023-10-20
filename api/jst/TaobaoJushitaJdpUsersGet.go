package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaojushitajdpusersget 获取开通的订单同步服务的用户
// taobao.jushita.jdp.users.get
//
// 获取开通的订单同步服务的用户，含有rds的路由关系
func Taobaojushitajdpusersget(clt *core.SDKClient, req *jst.TaobaojushitajdpusersgetAPIRequest, session string) (*jst.TaobaojushitajdpusersgetAPIResponse, error) {
	var resp jst.TaobaojushitajdpusersgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
