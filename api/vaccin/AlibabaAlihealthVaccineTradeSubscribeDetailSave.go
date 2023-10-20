package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// Alibabaalihealthvaccinetradesubscribedetailsave 私立疫苗交易-预约详情更新或保存
// alibaba.alihealth.vaccine.trade.subscribe.detail.save
//
// 私立疫苗交易-预约详情更新或保存
func Alibabaalihealthvaccinetradesubscribedetailsave(clt *core.SDKClient, req *vaccin.AlibabaalihealthvaccinetradesubscribedetailsaveAPIRequest, session string) (*vaccin.AlibabaalihealthvaccinetradesubscribedetailsaveAPIResponse, error) {
	var resp vaccin.AlibabaalihealthvaccinetradesubscribedetailsaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
