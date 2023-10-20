package mc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mc"
)

// TmallMcDeviceCircleCheck 云码设备圈选情况查询
// tmall.mc.device.circle.check
//
// 云码设备圈选情况查询
func TmallMcDeviceCircleCheck(clt *core.SDKClient, req *mc.TmallMcDeviceCircleCheckAPIRequest, session string) (*mc.TmallMcDeviceCircleCheckAPIResponse, error) {
	var resp mc.TmallMcDeviceCircleCheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
