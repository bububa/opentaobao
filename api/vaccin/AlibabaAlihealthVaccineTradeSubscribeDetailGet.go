package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// Alibabaalihealthvaccinetradesubscribedetailget 私立疫苗交易-预约详情获取
// alibaba.alihealth.vaccine.trade.subscribe.detail.get
//
// 私立疫苗交易-预约详情获取
func Alibabaalihealthvaccinetradesubscribedetailget(clt *core.SDKClient, req *vaccin.AlibabaalihealthvaccinetradesubscribedetailgetAPIRequest, session string) (*vaccin.AlibabaalihealthvaccinetradesubscribedetailgetAPIResponse, error) {
	var resp vaccin.AlibabaalihealthvaccinetradesubscribedetailgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
