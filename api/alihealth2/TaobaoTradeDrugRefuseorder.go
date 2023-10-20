package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// TaobaoTradeDrugRefuseorder 阿里健康020拒单
// taobao.trade.drug.refuseorder
//
// 阿里健康020拒单
func TaobaoTradeDrugRefuseorder(clt *core.SDKClient, req *alihealth2.TaobaoTradeDrugRefuseorderAPIRequest, resp *alihealth2.TaobaoTradeDrugRefuseorderAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
