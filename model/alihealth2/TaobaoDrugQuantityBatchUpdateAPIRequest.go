package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoDrugQuantityBatchUpdateAPIRequest
批量同步库存接口 API请求
taobao.drug.quantity.batch.update

商家通过top接口可以批量修改商品库存 */
type TaobaoDrugQuantityBatchUpdateAPIRequest struct {
	model.Params
	// 外部店铺ID
	_outStoreId string
	// 商品ID和库存
	_outItemIdQuantityMap string
}

// New
