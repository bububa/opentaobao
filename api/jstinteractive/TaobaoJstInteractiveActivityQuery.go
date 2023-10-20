package jstinteractive

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstinteractive"
)

// Taobaojstinteractiveactivityquery 互动任务活动查询接口
// taobao.jst.interactive.activity.query
//
// 互动任务活动查询接口
func Taobaojstinteractiveactivityquery(clt *core.SDKClient, req *jstinteractive.TaobaojstinteractiveactivityqueryAPIRequest, session string) (*jstinteractive.TaobaojstinteractiveactivityqueryAPIResponse, error) {
	var resp jstinteractive.TaobaojstinteractiveactivityqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
