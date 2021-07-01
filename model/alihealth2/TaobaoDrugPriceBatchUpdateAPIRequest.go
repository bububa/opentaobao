package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoDrugPriceBatchUpdateAPIRequest
商家批量更新宝贝价格 API请求
taobao.drug.price.batch.update

商家批量更新宝贝价格 */
type TaobaoDrugPriceBatchUpdateAPIRequest struct {
	model.Params
	// 外部店铺ID
	_outStoreId string
	// 商品ID和价格
	_outItemIdPriceMap string
}

// New
