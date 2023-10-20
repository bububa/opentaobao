package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthmedicalbasedoctorstatussync 挂号医生上下架
// alibaba.alihealth.medicalbase.doctor.status.sync
//
// 挂号医院上下线
func Alibabaalihealthmedicalbasedoctorstatussync(clt *core.SDKClient, req *alihealth2.AlibabaalihealthmedicalbasedoctorstatussyncAPIRequest, session string) (*alihealth2.AlibabaalihealthmedicalbasedoctorstatussyncAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthmedicalbasedoctorstatussyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
