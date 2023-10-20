package tmc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// TaobaoTmcGroupAdd 为已开通用户添加用户分组
// taobao.tmc.group.add
//
// 为已开通用户添加用户分组，授权消息使用
func TaobaoTmcGroupAdd(clt *core.SDKClient, req *tmc.TaobaoTmcGroupAddAPIRequest, session string) (*tmc.TaobaoTmcGroupAddAPIResponse, error) {
	var resp tmc.TaobaoTmcGroupAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
