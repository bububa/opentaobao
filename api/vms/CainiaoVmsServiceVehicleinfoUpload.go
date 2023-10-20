package vms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vms"
)

// Cainiaovmsservicevehicleinfoupload 新能源车--外部车辆信息回传
// cainiao.vms.service.vehicleinfo.upload
//
// 新能源车--外部车辆信息回传
func Cainiaovmsservicevehicleinfoupload(clt *core.SDKClient, req *vms.CainiaovmsservicevehicleinfouploadAPIRequest, session string) (*vms.CainiaovmsservicevehicleinfouploadAPIResponse, error) {
	var resp vms.CainiaovmsservicevehicleinfouploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
