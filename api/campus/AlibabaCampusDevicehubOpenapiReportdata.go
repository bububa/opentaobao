package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusDevicehubOpenapiReportdata 设备数据上报
// alibaba.campus.devicehub.openapi.reportdata
//
// 设备数据上报
func AlibabaCampusDevicehubOpenapiReportdata(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusDevicehubOpenapiReportdataAPIRequest, resp *campus.AlibabaCampusDevicehubOpenapiReportdataAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
