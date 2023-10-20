package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaodrugpricebatchupdateAPIRequest 商家批量更新宝贝价格 API请求
// taobao.drug.price.batch.update
//
// 商家批量更新宝贝价格
type TaobaodrugpricebatchupdateAPIRequest struct {
	model.Params
	// 外部店铺ID
	_outStoreId string
	// 商品ID和价格
	_outItemIdPriceMap string
}

// NewTaobaodrugpricebatchupdateRequest 初始化TaobaodrugpricebatchupdateAPIRequest对象
func NewTaobaodrugpricebatchupdateRequest() *TaobaodrugpricebatchupdateAPIRequest {
	return &TaobaodrugpricebatchupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaodrugpricebatchupdateAPIRequest) GetApiMethodName() string {
	return "taobao.drug.price.batch.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaodrugpricebatchupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaodrugpricebatchupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutStoreId is OutStoreId Setter
// 外部店铺ID
func (r *TaobaodrugpricebatchupdateAPIRequest) SetOutStoreId(_outStoreId string) error {
	r._outStoreId = _outStoreId
	r.Set("out_store_id", _outStoreId)
	return nil
}

// GetOutStoreId OutStoreId Getter
func (r TaobaodrugpricebatchupdateAPIRequest) GetOutStoreId() string {
	return r._outStoreId
}

// SetOutItemIdPriceMap is OutItemIdPriceMap Setter
// 商品ID和价格
func (r *TaobaodrugpricebatchupdateAPIRequest) SetOutItemIdPriceMap(_outItemIdPriceMap string) error {
	r._outItemIdPriceMap = _outItemIdPriceMap
	r.Set("out_item_id_price_map", _outItemIdPriceMap)
	return nil
}

// GetOutItemIdPriceMap OutItemIdPriceMap Getter
func (r TaobaodrugpricebatchupdateAPIRequest) GetOutItemIdPriceMap() string {
	return r._outItemIdPriceMap
}
