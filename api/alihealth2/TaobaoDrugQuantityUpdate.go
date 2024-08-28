package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// TaobaoDrugQuantityUpdate 商家更新库存
// taobao.drug.quantity.update
//
// 商家通过top接口可以直接修改商品库存
func TaobaoDrugQuantityUpdate(ctx context.Context, clt *core.SDKClient, req *alihealth2.TaobaoDrugQuantityUpdateAPIRequest, resp *alihealth2.TaobaoDrugQuantityUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
