package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// TaobaoTradeDrugConfirmorder 阿里健康020接单
// taobao.trade.drug.confirmorder
//
// 阿里健康020接单
func TaobaoTradeDrugConfirmorder(clt *core.SDKClient, req *alihealth2.TaobaoTradeDrugConfirmorderAPIRequest, resp *alihealth2.TaobaoTradeDrugConfirmorderAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
