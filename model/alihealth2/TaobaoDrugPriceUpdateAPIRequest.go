package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoDrugPriceUpdateAPIRequest
商家更新宝贝价格 API请求
taobao.drug.price.update

商家更新价格 */
type TaobaoDrugPriceUpdateAPIRequest struct {
	model.Params
	// 对应的外部店铺ID
	_outStoreId string
	// 对应的外部商品编码
	_outItemId string
	// 商品价格
	_price float64
}

// New
