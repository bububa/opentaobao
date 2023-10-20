package tmc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// Taobaotmcusertopicsget 获取用户开通的topic列表
// taobao.tmc.user.topics.get
//
// 获取用户开通的topic列表，授权消息使用
func Taobaotmcusertopicsget(clt *core.SDKClient, req *tmc.TaobaotmcusertopicsgetAPIRequest, session string) (*tmc.TaobaotmcusertopicsgetAPIResponse, error) {
	var resp tmc.TaobaotmcusertopicsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
