package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUsceslBizStoreInsertAPIRequest
新增电子价签商家门店接口 API请求
taobao.uscesl.biz.store.insert

新增电子价签商家门店接口 */
type TaobaoUsceslBizStoreInsertAPIRequest struct {
	model.Params
	// 门店名称
	_storeName string
	// 门店外部ID，要保持同一商家下的唯一性
	_storeOutId string
	// 商家标识key
	_bizBrandKey string
}

// NewTaobaoUsceslBizStoreInsertRequest 初始化TaobaoUsceslBizStoreInsertAPIRequest对象
func NewTaobaoUsceslBizStoreInsertRequest() *TaobaoUsceslBizStoreInsertAPIRequest {
	return &TaobaoUsceslBizStoreInsertAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUsceslBizStoreInsertAPIRequest) GetApiMethodName() string {
	return "taobao.uscesl.biz.store.insert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUsceslBizStoreInsertAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StoreName Setter
// 门店名称
func (r *TaobaoUsceslBizStoreInsertAPIRequest) SetStoreName(_storeName string) error {
	r._storeName = _storeName
	r.Set("store_name", _storeName)
	return nil
}

// Get StoreName Getter
func (r TaobaoUsceslBizStoreInsertAPIRequest) GetStoreName() string {
	return r._storeName
}

// Set is StoreOutId Setter
// 门店外部ID，要保持同一商家下的唯一性
func (r *TaobaoUsceslBizStoreInsertAPIRequest) SetStoreOutId(_storeOutId string) error {
	r._storeOutId = _storeOutId
	r.Set("store_out_id", _storeOutId)
	return nil
}

// Get StoreOutId Getter
func (r TaobaoUsceslBizStoreInsertAPIRequest) GetStoreOutId() string {
	return r._storeOutId
}

// Set is BizBrandKey Setter
// 商家标识key
func (r *TaobaoUsceslBizStoreInsertAPIRequest) SetBizBrandKey(_bizBrandKey string) error {
	r._bizBrandKey = _bizBrandKey
	r.Set("biz_brand_key", _bizBrandKey)
	return nil
}

// Get BizBrandKey Getter
func (r TaobaoUsceslBizStoreInsertAPIRequest) GetBizBrandKey() string {
	return r._bizBrandKey
}
