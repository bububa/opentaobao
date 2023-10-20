package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaAlihealthVaccineTradeSubscribeDetailGet 私立疫苗交易-预约详情获取
// alibaba.alihealth.vaccine.trade.subscribe.detail.get
//
// 私立疫苗交易-预约详情获取
func AlibabaAlihealthVaccineTradeSubscribeDetailGet(clt *core.SDKClient, req *vaccin.AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIRequest, resp *vaccin.AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
