package vms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vms"
)

// CainiaoVmsServiceVehicleinfoUpload 新能源车--外部车辆信息回传
// cainiao.vms.service.vehicleinfo.upload
//
// 新能源车--外部车辆信息回传
func CainiaoVmsServiceVehicleinfoUpload(clt *core.SDKClient, req *vms.CainiaoVmsServiceVehicleinfoUploadAPIRequest, resp *vms.CainiaoVmsServiceVehicleinfoUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
