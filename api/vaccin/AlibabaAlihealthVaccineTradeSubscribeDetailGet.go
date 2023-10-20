package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaAlihealthVaccineTradeSubscribeDetailGet 私立疫苗交易-预约详情获取
// alibaba.alihealth.vaccine.trade.subscribe.detail.get
//
// 私立疫苗交易-预约详情获取
func AlibabaAlihealthVaccineTradeSubscribeDetailGet(clt *core.SDKClient, req *vaccin.AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIRequest, session string) (*vaccin.AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIResponse, error) {
	var resp vaccin.AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
