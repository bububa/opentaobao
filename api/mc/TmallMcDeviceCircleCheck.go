package mc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mc"
)

// TmallMcDeviceCircleCheck 云码设备圈选情况查询
// tmall.mc.device.circle.check
//
// 云码设备圈选情况查询
func TmallMcDeviceCircleCheck(clt *core.SDKClient, req *mc.TmallMcDeviceCircleCheckAPIRequest, resp *mc.TmallMcDeviceCircleCheckAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
