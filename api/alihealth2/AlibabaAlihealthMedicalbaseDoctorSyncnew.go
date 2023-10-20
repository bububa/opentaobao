package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthmedicalbasedoctorsyncnew 直连医生上传
// alibaba.alihealth.medicalbase.doctor.syncnew
//
// 直连医生上传
func Alibabaalihealthmedicalbasedoctorsyncnew(clt *core.SDKClient, req *alihealth2.AlibabaalihealthmedicalbasedoctorsyncnewAPIRequest, session string) (*alihealth2.AlibabaalihealthmedicalbasedoctorsyncnewAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthmedicalbasedoctorsyncnewAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
