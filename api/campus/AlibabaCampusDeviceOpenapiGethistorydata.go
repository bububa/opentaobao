package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusDeviceOpenapiGethistorydata 查询设备历史数据
// alibaba.campus.device.openapi.gethistorydata
//
// 查询历史数据的接口
func AlibabaCampusDeviceOpenapiGethistorydata(clt *core.SDKClient, req *campus.AlibabaCampusDeviceOpenapiGethistorydataAPIRequest, resp *campus.AlibabaCampusDeviceOpenapiGethistorydataAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
