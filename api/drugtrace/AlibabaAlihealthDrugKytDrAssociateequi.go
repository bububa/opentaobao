package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytdrassociateequi 疫苗单据与设备绑定
// alibaba.alihealth.drug.kyt.dr.associateequi
//
// 疫苗单据与设备绑定
func Alibabaalihealthdrugkytdrassociateequi(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytdrassociateequiAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytdrassociateequiAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytdrassociateequiAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
