package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Taobaoqimentradeusersget 获取奇门用户列表
// taobao.qimen.trade.users.get
//
// 获取已开通奇门订单服务的用户列表
func Taobaoqimentradeusersget(clt *core.SDKClient, req *util.TaobaoqimentradeusersgetAPIRequest, session string) (*util.TaobaoqimentradeusersgetAPIResponse, error) {
	var resp util.TaobaoqimentradeusersgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
