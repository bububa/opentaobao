package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractsensorcalendar 天猫互动游戏开放平台需要授权的传感器类接口(日历提醒)
// alibaba.interact.sensor.calendar
//
// 天猫互动游戏开放平台需要授权的传感器类接口(日历提醒)
func Alibabainteractsensorcalendar(clt *core.SDKClient, req *interact.AlibabainteractsensorcalendarAPIRequest, session string) (*interact.AlibabainteractsensorcalendarAPIResponse, error) {
	var resp interact.AlibabainteractsensorcalendarAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
