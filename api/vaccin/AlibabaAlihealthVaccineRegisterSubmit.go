package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// Alibabaalihealthvaccineregistersubmit cdc回传疫苗登记数据
// alibaba.alihealth.vaccine.register.submit
//
// cdc回传疫苗登记信息
func Alibabaalihealthvaccineregistersubmit(clt *core.SDKClient, req *vaccin.AlibabaalihealthvaccineregistersubmitAPIRequest, session string) (*vaccin.AlibabaalihealthvaccineregistersubmitAPIResponse, error) {
	var resp vaccin.AlibabaalihealthvaccineregistersubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
