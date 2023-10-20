package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaAlihealthVaccineTradeSubscribeDetailSave 私立疫苗交易-预约详情更新或保存
// alibaba.alihealth.vaccine.trade.subscribe.detail.save
//
// 私立疫苗交易-预约详情更新或保存
func AlibabaAlihealthVaccineTradeSubscribeDetailSave(clt *core.SDKClient, req *vaccin.AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest, resp *vaccin.AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
