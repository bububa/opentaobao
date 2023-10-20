package tbuser

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbuser"
)

// TaobaoUserOpenuidGetbyorder 根据订单获取买家openuid
// taobao.user.openuid.getbyorder
//
// 根据订单获取买家openuid，最大查询30个
func TaobaoUserOpenuidGetbyorder(clt *core.SDKClient, req *tbuser.TaobaoUserOpenuidGetbyorderAPIRequest, resp *tbuser.TaobaoUserOpenuidGetbyorderAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
