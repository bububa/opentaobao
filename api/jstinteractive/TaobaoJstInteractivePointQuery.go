package jstinteractive

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstinteractive"
)

// TaobaoJstInteractivePointQuery 互动积分查询接口
// taobao.jst.interactive.point.query
//
// 查询用户的互动积分
func TaobaoJstInteractivePointQuery(clt *core.SDKClient, req *jstinteractive.TaobaoJstInteractivePointQueryAPIRequest, session string) (*jstinteractive.TaobaoJstInteractivePointQueryAPIResponse, error) {
	var resp jstinteractive.TaobaoJstInteractivePointQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
