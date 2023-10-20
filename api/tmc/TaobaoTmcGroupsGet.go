package tmc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// Taobaotmcgroupsget 获取自定义用户分组列表
// taobao.tmc.groups.get
//
// 获取自定义用户分组列表
func Taobaotmcgroupsget(clt *core.SDKClient, req *tmc.TaobaotmcgroupsgetAPIRequest, session string) (*tmc.TaobaotmcgroupsgetAPIResponse, error) {
	var resp tmc.TaobaotmcgroupsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
