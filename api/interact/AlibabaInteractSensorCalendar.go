package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorCalendar 天猫互动游戏开放平台需要授权的传感器类接口(日历提醒)
// alibaba.interact.sensor.calendar
//
// 天猫互动游戏开放平台需要授权的传感器类接口(日历提醒)
func AlibabaInteractSensorCalendar(clt *core.SDKClient, req *interact.AlibabaInteractSensorCalendarAPIRequest, resp *interact.AlibabaInteractSensorCalendarAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
