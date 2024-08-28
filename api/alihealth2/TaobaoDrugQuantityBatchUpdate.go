package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// TaobaoDrugQuantityBatchUpdate 批量同步库存接口
// taobao.drug.quantity.batch.update
//
// 商家通过top接口可以批量修改商品库存
func TaobaoDrugQuantityBatchUpdate(ctx context.Context, clt *core.SDKClient, req *alihealth2.TaobaoDrugQuantityBatchUpdateAPIRequest, resp *alihealth2.TaobaoDrugQuantityBatchUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
