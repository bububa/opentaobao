package util

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// WdkRexoutDeviceIotRegisterid 通过设备ID获取三元组-外部
// wdk.rexout.device.iot.registerid
//
// 通过设备ID获取三元组-外部
func WdkRexoutDeviceIotRegisterid(ctx context.Context, clt *core.SDKClient, req *util.WdkRexoutDeviceIotRegisteridAPIRequest, resp *util.WdkRexoutDeviceIotRegisteridAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
