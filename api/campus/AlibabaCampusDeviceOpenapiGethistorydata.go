package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusDeviceOpenapiGethistorydata 查询设备历史数据
// alibaba.campus.device.openapi.gethistorydata
//
// 查询历史数据的接口
func AlibabaCampusDeviceOpenapiGethistorydata(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusDeviceOpenapiGethistorydataAPIRequest, resp *campus.AlibabaCampusDeviceOpenapiGethistorydataAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
