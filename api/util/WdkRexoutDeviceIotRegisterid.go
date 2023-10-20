package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Wdkrexoutdeviceiotregisterid 通过设备ID获取三元组-外部
// wdk.rexout.device.iot.registerid
//
// 通过设备ID获取三元组-外部
func Wdkrexoutdeviceiotregisterid(clt *core.SDKClient, req *util.WdkrexoutdeviceiotregisteridAPIRequest, session string) (*util.WdkrexoutdeviceiotregisteridAPIResponse, error) {
	var resp util.WdkrexoutdeviceiotregisteridAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
