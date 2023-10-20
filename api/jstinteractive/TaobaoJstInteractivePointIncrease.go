package jstinteractive

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstinteractive"
)

// Taobaojstinteractivepointincrease 互动积分发放接口
// taobao.jst.interactive.point.increase
//
// 向用户发放互动积分
func Taobaojstinteractivepointincrease(clt *core.SDKClient, req *jstinteractive.TaobaojstinteractivepointincreaseAPIRequest, session string) (*jstinteractive.TaobaojstinteractivepointincreaseAPIResponse, error) {
	var resp jstinteractive.TaobaojstinteractivepointincreaseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
