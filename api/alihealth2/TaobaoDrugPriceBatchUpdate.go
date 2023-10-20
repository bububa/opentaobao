package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// TaobaoDrugPriceBatchUpdate 商家批量更新宝贝价格
// taobao.drug.price.batch.update
//
// 商家批量更新宝贝价格
func TaobaoDrugPriceBatchUpdate(clt *core.SDKClient, req *alihealth2.TaobaoDrugPriceBatchUpdateAPIRequest, resp *alihealth2.TaobaoDrugPriceBatchUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
