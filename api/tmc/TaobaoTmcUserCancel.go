package tmc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// TaobaoTmcUserCancel 取消用户的消息服务
// taobao.tmc.user.cancel
//
// 取消用户的消息服务
func TaobaoTmcUserCancel(clt *core.SDKClient, req *tmc.TaobaoTmcUserCancelAPIRequest, session string) (*tmc.TaobaoTmcUserCancelAPIResponse, error) {
	var resp tmc.TaobaoTmcUserCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
