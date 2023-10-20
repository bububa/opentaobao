package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytdrgetproteminfo 疫苗，获取生产企业的存储和运输温度
// alibaba.alihealth.drug.kyt.dr.getproteminfo
//
// 疫苗，获取生产企业的存储和运输温度
func Alibabaalihealthdrugkytdrgetproteminfo(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytdrgetproteminfoAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytdrgetproteminfoAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytdrgetproteminfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
