package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoDrugQuantityUpdateAPIRequest
商家更新库存 API请求
taobao.drug.quantity.update

商家通过top接口可以直接修改商品库存 */
type TaobaoDrugQuantityUpdateAPIRequest struct {
	model.Params
	// 外部店铺ID
	_outStoreId string
	// 外部商品ID
	_outItemId string
	// 库存数量
	_quantity int64
}

// New
