package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytspeciavaccinsearchstatus 疫苗企业上传单据后处理状态查询
// alibaba.alihealth.drug.kyt.specia.vaccin.searchstatus
//
// 单据处理状态查询
func Alibabaalihealthdrugkytspeciavaccinsearchstatus(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytspeciavaccinsearchstatusAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytspeciavaccinsearchstatusAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytspeciavaccinsearchstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
