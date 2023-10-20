package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// TaobaoTradeDrugConfirmorder 阿里健康020接单
// taobao.trade.drug.confirmorder
//
// 阿里健康020接单
func TaobaoTradeDrugConfirmorder(clt *core.SDKClient, req *alihealth2.TaobaoTradeDrugConfirmorderAPIRequest, session string) (*alihealth2.TaobaoTradeDrugConfirmorderAPIResponse, error) {
	var resp alihealth2.TaobaoTradeDrugConfirmorderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
