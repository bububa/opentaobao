package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytspeciavaccinuploadretail 零售单据上传接口（疫苗）
// alibaba.alihealth.drug.kyt.specia.vaccin.uploadretail
//
// 零售上传单据信息接口
func Alibabaalihealthdrugkytspeciavaccinuploadretail(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytspeciavaccinuploadretailAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytspeciavaccinuploadretailAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytspeciavaccinuploadretailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
