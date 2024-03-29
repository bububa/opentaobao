package uscesl

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsceslBizStoreInsertAPIRequest 新增电子价签商家门店接口 API请求
// taobao.uscesl.biz.store.insert
//
// 新增电子价签商家门店接口
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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUsceslBizStoreInsertAPIRequest) Reset() {
	r._storeName = ""
	r._storeOutId = ""
	r._bizBrandKey = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUsceslBizStoreInsertAPIRequest) GetApiMethodName() string {
	return "taobao.uscesl.biz.store.insert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUsceslBizStoreInsertAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUsceslBizStoreInsertAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreName is StoreName Setter
// 门店名称
func (r *TaobaoUsceslBizStoreInsertAPIRequest) SetStoreName(_storeName string) error {
	r._storeName = _storeName
	r.Set("store_name", _storeName)
	return nil
}

// GetStoreName StoreName Getter
func (r TaobaoUsceslBizStoreInsertAPIRequest) GetStoreName() string {
	return r._storeName
}

// SetStoreOutId is StoreOutId Setter
// 门店外部ID，要保持同一商家下的唯一性
func (r *TaobaoUsceslBizStoreInsertAPIRequest) SetStoreOutId(_storeOutId string) error {
	r._storeOutId = _storeOutId
	r.Set("store_out_id", _storeOutId)
	return nil
}

// GetStoreOutId StoreOutId Getter
func (r TaobaoUsceslBizStoreInsertAPIRequest) GetStoreOutId() string {
	return r._storeOutId
}

// SetBizBrandKey is BizBrandKey Setter
// 商家标识key
func (r *TaobaoUsceslBizStoreInsertAPIRequest) SetBizBrandKey(_bizBrandKey string) error {
	r._bizBrandKey = _bizBrandKey
	r.Set("biz_brand_key", _bizBrandKey)
	return nil
}

// GetBizBrandKey BizBrandKey Getter
func (r TaobaoUsceslBizStoreInsertAPIRequest) GetBizBrandKey() string {
	return r._bizBrandKey
}

var poolTaobaoUsceslBizStoreInsertAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUsceslBizStoreInsertRequest()
	},
}

// GetTaobaoUsceslBizStoreInsertRequest 从 sync.Pool 获取 TaobaoUsceslBizStoreInsertAPIRequest
func GetTaobaoUsceslBizStoreInsertAPIRequest() *TaobaoUsceslBizStoreInsertAPIRequest {
	return poolTaobaoUsceslBizStoreInsertAPIRequest.Get().(*TaobaoUsceslBizStoreInsertAPIRequest)
}

// ReleaseTaobaoUsceslBizStoreInsertAPIRequest 将 TaobaoUsceslBizStoreInsertAPIRequest 放入 sync.Pool
func ReleaseTaobaoUsceslBizStoreInsertAPIRequest(v *TaobaoUsceslBizStoreInsertAPIRequest) {
	v.Reset()
	poolTaobaoUsceslBizStoreInsertAPIRequest.Put(v)
}
