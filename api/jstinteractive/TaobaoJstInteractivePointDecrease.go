package jstinteractive

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstinteractive"
)

// Taobaojstinteractivepointdecrease 互动积分扣减接口
// taobao.jst.interactive.point.decrease
//
// 扣减用户的互动积分
func Taobaojstinteractivepointdecrease(clt *core.SDKClient, req *jstinteractive.TaobaojstinteractivepointdecreaseAPIRequest, session string) (*jstinteractive.TaobaojstinteractivepointdecreaseAPIResponse, error) {
	var resp jstinteractive.TaobaojstinteractivepointdecreaseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
