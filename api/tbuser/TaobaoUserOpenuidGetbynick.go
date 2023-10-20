package tbuser

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbuser"
)

// Taobaouseropenuidgetbynick 根据买家nick获取买家openuid
// taobao.user.openuid.getbynick
//
// 根据买家nick获取买家openuid，最大查询30个
func Taobaouseropenuidgetbynick(clt *core.SDKClient, req *tbuser.TaobaouseropenuidgetbynickAPIRequest, session string) (*tbuser.TaobaouseropenuidgetbynickAPIResponse, error) {
	var resp tbuser.TaobaouseropenuidgetbynickAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
