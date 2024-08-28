package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusDeviceGetdeviceforquery 下发设备的分页接口(无需AOP控制)
// alibaba.campus.device.getdeviceforquery
//
// 下发设备的分页接口(发布在TOP上，connect调用，无需AOP控制)
func AlibabaCampusDeviceGetdeviceforquery(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusDeviceGetdeviceforqueryAPIRequest, resp *campus.AlibabaCampusDeviceGetdeviceforqueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
