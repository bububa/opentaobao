package alilabs

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// AlibabaAilabsTvsDeviceList 获取TVS设备列表
// alibaba.ailabs.tvs.device.list
//
// 获取用户所绑定的TVS设备列表
func AlibabaAilabsTvsDeviceList(ctx context.Context, clt *core.SDKClient, req *alilabs.AlibabaAilabsTvsDeviceListAPIRequest, resp *alilabs.AlibabaAilabsTvsDeviceListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
