package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// WdkRexoutDeviceIotRegisterid 通过设备ID获取三元组-外部
// wdk.rexout.device.iot.registerid
//
// 通过设备ID获取三元组-外部
func WdkRexoutDeviceIotRegisterid(clt *core.SDKClient, req *util.WdkRexoutDeviceIotRegisteridAPIRequest, session string) (*util.WdkRexoutDeviceIotRegisteridAPIResponse, error) {
	var resp util.WdkRexoutDeviceIotRegisteridAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
