package jstinteractive

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstinteractive"
)

// Taobaojstinteractiveactivitycreate 互动任务活动创建接口
// taobao.jst.interactive.activity.create
//
// 调用活动创建接口为小程序创建互动任务活动，任务列表仅在活动期间内返回
func Taobaojstinteractiveactivitycreate(clt *core.SDKClient, req *jstinteractive.TaobaojstinteractiveactivitycreateAPIRequest, session string) (*jstinteractive.TaobaojstinteractiveactivitycreateAPIResponse, error) {
	var resp jstinteractive.TaobaojstinteractiveactivitycreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
