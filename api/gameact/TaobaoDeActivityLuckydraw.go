package gameact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/gameact"
)

// Taobaodeactivityluckydraw 抽奖
// taobao.de.activity.luckydraw
//
// 用于激励平台对外提供抽奖功能，包括但不限于集分宝、红包、宝点、淘金币、淘彩票等
func Taobaodeactivityluckydraw(clt *core.SDKClient, req *gameact.TaobaodeactivityluckydrawAPIRequest, session string) (*gameact.TaobaodeactivityluckydrawAPIResponse, error) {
	var resp gameact.TaobaodeactivityluckydrawAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
