package tbuser

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbuser"
)

// Taobaouseropenuidgetbyorder 根据订单获取买家openuid
// taobao.user.openuid.getbyorder
//
// 根据订单获取买家openuid，最大查询30个
func Taobaouseropenuidgetbyorder(clt *core.SDKClient, req *tbuser.TaobaouseropenuidgetbyorderAPIRequest, session string) (*tbuser.TaobaouseropenuidgetbyorderAPIResponse, error) {
	var resp tbuser.TaobaouseropenuidgetbyorderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
