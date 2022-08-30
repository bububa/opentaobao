package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaAlihealthVaccineTradeSubscribeDetailSave 私立疫苗交易-预约详情更新或保存
// alibaba.alihealth.vaccine.trade.subscribe.detail.save
//
// 私立疫苗交易-预约详情更新或保存
func AlibabaAlihealthVaccineTradeSubscribeDetailSave(clt *core.SDKClient, req *vaccin.AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest, session string) (*vaccin.AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIResponse, error) {
	var resp vaccin.AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
