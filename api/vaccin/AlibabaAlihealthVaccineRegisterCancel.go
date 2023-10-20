package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// Alibabaalihealthvaccineregistercancel 取消登记
// alibaba.alihealth.vaccine.register.cancel
//
// 取消登记
func Alibabaalihealthvaccineregistercancel(clt *core.SDKClient, req *vaccin.AlibabaalihealthvaccineregistercancelAPIRequest, session string) (*vaccin.AlibabaalihealthvaccineregistercancelAPIResponse, error) {
	var resp vaccin.AlibabaalihealthvaccineregistercancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
