package jstinteractive

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstinteractive"
)

// Taobaojstinteractiveassetsconfiguredquery 查询已配置的任务素材列表接口
// taobao.jst.interactive.assets.configured.query
//
// 查询已配置任务素材列表
func Taobaojstinteractiveassetsconfiguredquery(clt *core.SDKClient, req *jstinteractive.TaobaojstinteractiveassetsconfiguredqueryAPIRequest, session string) (*jstinteractive.TaobaojstinteractiveassetsconfiguredqueryAPIResponse, error) {
	var resp jstinteractive.TaobaojstinteractiveassetsconfiguredqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
