package tmc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// Taobaotmcuserget 获取用户已开通消息
// taobao.tmc.user.get
//
// 查询指定用户开通的消息通道和组
func Taobaotmcuserget(clt *core.SDKClient, req *tmc.TaobaotmcusergetAPIRequest, session string) (*tmc.TaobaotmcusergetAPIResponse, error) {
	var resp tmc.TaobaotmcusergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
