package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaousceslbizstoreinsertAPIRequest 新增电子价签商家门店接口 API请求
// taobao.uscesl.biz.store.insert
//
// 新增电子价签商家门店接口
type TaobaousceslbizstoreinsertAPIRequest struct {
	model.Params
	// 门店名称
	_storeName string
	// 门店外部ID，要保持同一商家下的唯一性
	_storeOutId string
	// 商家标识key
	_bizBrandKey string
}

// NewTaobaousceslbizstoreinsertRequest 初始化TaobaousceslbizstoreinsertAPIRequest对象
func NewTaobaousceslbizstoreinsertRequest() *TaobaousceslbizstoreinsertAPIRequest {
	return &TaobaousceslbizstoreinsertAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaousceslbizstoreinsertAPIRequest) GetApiMethodName() string {
	return "taobao.uscesl.biz.store.insert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaousceslbizstoreinsertAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaousceslbizstoreinsertAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreName is StoreName Setter
// 门店名称
func (r *TaobaousceslbizstoreinsertAPIRequest) SetStoreName(_storeName string) error {
	r._storeName = _storeName
	r.Set("store_name", _storeName)
	return nil
}

// GetStoreName StoreName Getter
func (r TaobaousceslbizstoreinsertAPIRequest) GetStoreName() string {
	return r._storeName
}

// SetStoreOutId is StoreOutId Setter
// 门店外部ID，要保持同一商家下的唯一性
func (r *TaobaousceslbizstoreinsertAPIRequest) SetStoreOutId(_storeOutId string) error {
	r._storeOutId = _storeOutId
	r.Set("store_out_id", _storeOutId)
	return nil
}

// GetStoreOutId StoreOutId Getter
func (r TaobaousceslbizstoreinsertAPIRequest) GetStoreOutId() string {
	return r._storeOutId
}

// SetBizBrandKey is BizBrandKey Setter
// 商家标识key
func (r *TaobaousceslbizstoreinsertAPIRequest) SetBizBrandKey(_bizBrandKey string) error {
	r._bizBrandKey = _bizBrandKey
	r.Set("biz_brand_key", _bizBrandKey)
	return nil
}

// GetBizBrandKey BizBrandKey Getter
func (r TaobaousceslbizstoreinsertAPIRequest) GetBizBrandKey() string {
	return r._bizBrandKey
}
