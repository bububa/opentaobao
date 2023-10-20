package tbuser

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbuser"
)

// TaobaoUserOpenuidGetbynick 根据买家nick获取买家openuid
// taobao.user.openuid.getbynick
//
// 根据买家nick获取买家openuid，最大查询30个
func TaobaoUserOpenuidGetbynick(clt *core.SDKClient, req *tbuser.TaobaoUserOpenuidGetbynickAPIRequest, resp *tbuser.TaobaoUserOpenuidGetbynickAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
