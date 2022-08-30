package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoUserOpenuidGetbyorder 根据订单获取买家openuid
// taobao.user.openuid.getbyorder
//
// 根据订单获取买家openuid，最大查询30个
func TaobaoUserOpenuidGetbyorder(clt *core.SDKClient, req *user.TaobaoUserOpenuidGetbyorderAPIRequest, session string) (*user.TaobaoUserOpenuidGetbyorderAPIResponse, error) {
	var resp user.TaobaoUserOpenuidGetbyorderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
