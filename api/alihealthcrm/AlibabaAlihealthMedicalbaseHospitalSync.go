package alihealthcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcrm"
)

// Alibabaalihealthmedicalbasehospitalsync 互联网医院批量导入接口
// alibaba.alihealth.medicalbase.hospital.sync
//
// 互联网医院isv批量通过接口批量导入
func Alibabaalihealthmedicalbasehospitalsync(clt *core.SDKClient, req *alihealthcrm.AlibabaalihealthmedicalbasehospitalsyncAPIRequest, session string) (*alihealthcrm.AlibabaalihealthmedicalbasehospitalsyncAPIResponse, error) {
	var resp alihealthcrm.AlibabaalihealthmedicalbasehospitalsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
