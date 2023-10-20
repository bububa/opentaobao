package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytdrgetupteminfo 获取上游温度信息（疫苗）
// alibaba.alihealth.drug.kyt.dr.getupteminfo
//
// 根据追溯码及企业ID获取上游运输及存储温度信息（疫苗）
func Alibabaalihealthdrugkytdrgetupteminfo(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytdrgetupteminfoAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytdrgetupteminfoAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytdrgetupteminfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
