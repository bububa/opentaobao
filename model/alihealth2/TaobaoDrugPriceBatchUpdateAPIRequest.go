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

// NewTaobaoDrugPriceBatchUpdateRequest 初始化TaobaoDrugPriceBatchUpdateAPIRequest对象
func NewTaobaoDrugPriceBatchUpdateRequest() *TaobaoDrugPriceBatchUpdateAPIRequest {
	return &TaobaoDrugPriceBatchUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDrugPriceBatchUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.drug.price.batch.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDrugPriceBatchUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OutStoreId Setter
// 外部店铺ID
func (r *TaobaoDrugPriceBatchUpdateAPIRequest) SetOutStoreId(_outStoreId string) error {
	r._outStoreId = _outStoreId
	r.Set("out_store_id", _outStoreId)
	return nil
}

// Get OutStoreId Getter
func (r TaobaoDrugPriceBatchUpdateAPIRequest) GetOutStoreId() string {
	return r._outStoreId
}

// Set is OutItemIdPriceMap Setter
// 商品ID和价格
func (r *TaobaoDrugPriceBatchUpdateAPIRequest) SetOutItemIdPriceMap(_outItemIdPriceMap string) error {
	r._outItemIdPriceMap = _outItemIdPriceMap
	r.Set("out_item_id_price_map", _outItemIdPriceMap)
	return nil
}

// Get OutItemIdPriceMap Getter
func (r TaobaoDrugPriceBatchUpdateAPIRequest) GetOutItemIdPriceMap() string {
	return r._outItemIdPriceMap
}
