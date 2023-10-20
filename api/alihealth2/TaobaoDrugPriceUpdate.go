package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// TaobaoDrugPriceUpdate 商家更新宝贝价格
// taobao.drug.price.update
//
// 商家更新价格
func TaobaoDrugPriceUpdate(clt *core.SDKClient, req *alihealth2.TaobaoDrugPriceUpdateAPIRequest, resp *alihealth2.TaobaoDrugPriceUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
