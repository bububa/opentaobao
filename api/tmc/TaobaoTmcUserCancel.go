package tmc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// Taobaotmcusercancel 取消用户的消息服务
// taobao.tmc.user.cancel
//
// 取消用户的消息服务
func Taobaotmcusercancel(clt *core.SDKClient, req *tmc.TaobaotmcusercancelAPIRequest, session string) (*tmc.TaobaotmcusercancelAPIResponse, error) {
	var resp tmc.TaobaotmcusercancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
