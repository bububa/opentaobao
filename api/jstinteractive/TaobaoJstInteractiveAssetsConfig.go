package jstinteractive

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstinteractive"
)

// Taobaojstinteractiveassetsconfig 任务素材配置接口
// taobao.jst.interactive.assets.config
//
// 任务素材配置接口
func Taobaojstinteractiveassetsconfig(clt *core.SDKClient, req *jstinteractive.TaobaojstinteractiveassetsconfigAPIRequest, session string) (*jstinteractive.TaobaojstinteractiveassetsconfigAPIResponse, error) {
	var resp jstinteractive.TaobaojstinteractiveassetsconfigAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
