package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// TaobaoDrugQuantityBatchUpdate 批量同步库存接口
// taobao.drug.quantity.batch.update
//
// 商家通过top接口可以批量修改商品库存
func TaobaoDrugQuantityBatchUpdate(clt *core.SDKClient, req *alihealth2.TaobaoDrugQuantityBatchUpdateAPIRequest, session string) (*alihealth2.TaobaoDrugQuantityBatchUpdateAPIResponse, error) {
	var resp alihealth2.TaobaoDrugQuantityBatchUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
