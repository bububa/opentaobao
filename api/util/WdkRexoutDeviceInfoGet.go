package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// WdkRexoutDeviceInfoGet 获取设备详情-外部对接
// wdk.rexout.device.info.get
//
// 获取设备详情-外部对接
func WdkRexoutDeviceInfoGet(clt *core.SDKClient, req *util.WdkRexoutDeviceInfoGetAPIRequest, session string) (*util.WdkRexoutDeviceInfoGetAPIResponse, error) {
	var resp util.WdkRexoutDeviceInfoGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
