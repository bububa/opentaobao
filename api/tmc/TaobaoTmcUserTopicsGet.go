package tmc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// TaobaoTmcUserTopicsGet 获取用户开通的topic列表
// taobao.tmc.user.topics.get
//
// 获取用户开通的topic列表，授权消息使用
func TaobaoTmcUserTopicsGet(clt *core.SDKClient, req *tmc.TaobaoTmcUserTopicsGetAPIRequest, session string) (*tmc.TaobaoTmcUserTopicsGetAPIResponse, error) {
	var resp tmc.TaobaoTmcUserTopicsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
