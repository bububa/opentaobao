package jstinteractive

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstinteractive"
)

// Taobaojstinteractiveactivityupdate 互动任务活动修改接口
// taobao.jst.interactive.activity.update
//
// 互动任务活动修改接口
func Taobaojstinteractiveactivityupdate(clt *core.SDKClient, req *jstinteractive.TaobaojstinteractiveactivityupdateAPIRequest, session string) (*jstinteractive.TaobaojstinteractiveactivityupdateAPIResponse, error) {
	var resp jstinteractive.TaobaojstinteractiveactivityupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
