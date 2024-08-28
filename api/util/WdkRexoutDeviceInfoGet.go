package util

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// WdkRexoutDeviceInfoGet 获取设备详情-外部对接
// wdk.rexout.device.info.get
//
// 获取设备详情-外部对接
func WdkRexoutDeviceInfoGet(ctx context.Context, clt *core.SDKClient, req *util.WdkRexoutDeviceInfoGetAPIRequest, resp *util.WdkRexoutDeviceInfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
