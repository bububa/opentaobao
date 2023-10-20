package jstinteractive

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstinteractive"
)

// Taobaojstinteractivepointquery 互动积分查询接口
// taobao.jst.interactive.point.query
//
// 查询用户的互动积分
func Taobaojstinteractivepointquery(clt *core.SDKClient, req *jstinteractive.TaobaojstinteractivepointqueryAPIRequest, session string) (*jstinteractive.TaobaojstinteractivepointqueryAPIResponse, error) {
	var resp jstinteractive.TaobaojstinteractivepointqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
