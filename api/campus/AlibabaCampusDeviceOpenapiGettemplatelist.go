package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusDeviceOpenapiGettemplatelist 查询设备模板
// alibaba.campus.device.openapi.gettemplatelist
//
// 查询设备模板信息
func AlibabaCampusDeviceOpenapiGettemplatelist(clt *core.SDKClient, req *campus.AlibabaCampusDeviceOpenapiGettemplatelistAPIRequest, session string) (*campus.AlibabaCampusDeviceOpenapiGettemplatelistAPIResponse, error) {
	var resp campus.AlibabaCampusDeviceOpenapiGettemplatelistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
