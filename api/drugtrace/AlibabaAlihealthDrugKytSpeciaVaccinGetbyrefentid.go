package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytspeciavaccingetbyrefentid 根据企业唯一标识查看企业详情
// alibaba.alihealth.drug.kyt.specia.vaccin.getbyrefentid
//
// 根据企业唯一标识查看企业详细信息
func Alibabaalihealthdrugkytspeciavaccingetbyrefentid(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytspeciavaccingetbyrefentidAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytspeciavaccingetbyrefentidAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytspeciavaccingetbyrefentidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
