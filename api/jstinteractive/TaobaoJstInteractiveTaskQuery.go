package jstinteractive

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstinteractive"
)

// Taobaojstinteractivetaskquery 互动任务列表查询接口
// taobao.jst.interactive.task.query
//
// 查询用户的互动任务列表
func Taobaojstinteractivetaskquery(clt *core.SDKClient, req *jstinteractive.TaobaojstinteractivetaskqueryAPIRequest, session string) (*jstinteractive.TaobaojstinteractivetaskqueryAPIResponse, error) {
	var resp jstinteractive.TaobaojstinteractivetaskqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
