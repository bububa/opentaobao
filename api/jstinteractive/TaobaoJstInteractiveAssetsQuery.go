package jstinteractive

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstinteractive"
)

// Taobaojstinteractiveassetsquery 查询可配置任务素材接口
// taobao.jst.interactive.assets.query
//
// 查询可配置任务素材列表，用以配置任务素材
func Taobaojstinteractiveassetsquery(clt *core.SDKClient, req *jstinteractive.TaobaojstinteractiveassetsqueryAPIRequest, session string) (*jstinteractive.TaobaojstinteractiveassetsqueryAPIResponse, error) {
	var resp jstinteractive.TaobaojstinteractiveassetsqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
