package tbuser

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbuser"
)

// TaobaoUserOpenuidGetbynick 根据买家nick获取买家openuid
// taobao.user.openuid.getbynick
//
// 根据买家nick获取买家openuid，最大查询30个
func TaobaoUserOpenuidGetbynick(clt *core.SDKClient, req *tbuser.TaobaoUserOpenuidGetbynickAPIRequest, session string) (*tbuser.TaobaoUserOpenuidGetbynickAPIResponse, error) {
	var resp tbuser.TaobaoUserOpenuidGetbynickAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
