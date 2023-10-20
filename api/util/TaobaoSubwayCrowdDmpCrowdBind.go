package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoSubwayCrowdDmpCrowdBind 直通车绑定达摩盘人群
// taobao.subway.crowd.dmp.crowd.bind
//
// 直通车绑定达摩盘人群
func TaobaoSubwayCrowdDmpCrowdBind(clt *core.SDKClient, req *util.TaobaoSubwayCrowdDmpCrowdBindAPIRequest, session string) (*util.TaobaoSubwayCrowdDmpCrowdBindAPIResponse, error) {
	var resp util.TaobaoSubwayCrowdDmpCrowdBindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
