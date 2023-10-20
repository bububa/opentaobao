package tmc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// Taobaotmcgroupadd 为已开通用户添加用户分组
// taobao.tmc.group.add
//
// 为已开通用户添加用户分组，授权消息使用
func Taobaotmcgroupadd(clt *core.SDKClient, req *tmc.TaobaotmcgroupaddAPIRequest, session string) (*tmc.TaobaotmcgroupaddAPIResponse, error) {
	var resp tmc.TaobaotmcgroupaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
