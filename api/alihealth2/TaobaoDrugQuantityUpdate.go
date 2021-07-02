package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// TaobaoDrugQuantityUpdate 商家更新库存
// taobao.drug.quantity.update
//
// 商家通过top接口可以直接修改商品库存
func TaobaoDrugQuantityUpdate(clt *core.SDKClient, req *alihealth2.TaobaoDrugQuantityUpdateAPIRequest, session string) (*alihealth2.TaobaoDrugQuantityUpdateAPIResponse, error) {
	var resp alihealth2.TaobaoDrugQuantityUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
