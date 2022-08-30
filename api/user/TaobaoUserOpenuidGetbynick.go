package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoUserOpenuidGetbynick 根据买家nick获取买家openuid
// taobao.user.openuid.getbynick
//
// 根据买家nick获取买家openuid，最大查询30个
func TaobaoUserOpenuidGetbynick(clt *core.SDKClient, req *user.TaobaoUserOpenuidGetbynickAPIRequest, session string) (*user.TaobaoUserOpenuidGetbynickAPIResponse, error) {
	var resp user.TaobaoUserOpenuidGetbynickAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
