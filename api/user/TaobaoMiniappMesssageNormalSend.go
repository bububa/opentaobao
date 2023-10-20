package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// Taobaominiappmesssagenormalsend 轻店铺下行普通消息给用户
// taobao.miniapp.messsage.normal.send
//
// 小程序下行单个普通消息
func Taobaominiappmesssagenormalsend(clt *core.SDKClient, req *user.TaobaominiappmesssagenormalsendAPIRequest, session string) (*user.TaobaominiappmesssagenormalsendAPIResponse, error) {
	var resp user.TaobaominiappmesssagenormalsendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
