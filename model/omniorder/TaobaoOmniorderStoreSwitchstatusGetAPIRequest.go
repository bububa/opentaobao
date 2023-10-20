package omniorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStoreSwitchstatusGetAPIRequest switchstatus.get API请求
// taobao.omniorder.store.switchstatus.get
//
// 查询门店发货、门店自提状态
type TaobaoOmniorderStoreSwitchstatusGetAPIRequest struct {
	model.Params
	// 门店ID
	_storeId int64
	// 卖家ID
	_sellerId int64
}

// NewTaobaoOmniorderStoreSwitchstatusGetRequest 初始化TaobaoOmniorderStoreSwitchstatusGetAPIRequest对象
func NewTaobaoOmniorderStoreSwitchstatusGetRequest() *TaobaoOmniorderStoreSwitchstatusGetAPIRequest {
	return &TaobaoOmniorderStoreSwitchstatusGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOmniorderStoreSwitchstatusGetAPIRequest) Reset() {
	r._storeId = 0
	r._sellerId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreSwitchstatusGetAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.store.switchstatus.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreSwitchstatusGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOmniorderStoreSwitchstatusGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *TaobaoOmniorderStoreSwitchstatusGetAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoOmniorderStoreSwitchstatusGetAPIRequest) GetStoreId() int64 {
	return r._storeId
}

// SetSellerId is SellerId Setter
// 卖家ID
func (r *TaobaoOmniorderStoreSwitchstatusGetAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// GetSellerId SellerId Getter
func (r TaobaoOmniorderStoreSwitchstatusGetAPIRequest) GetSellerId() int64 {
	return r._sellerId
}

var poolTaobaoOmniorderStoreSwitchstatusGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOmniorderStoreSwitchstatusGetRequest()
	},
}

// GetTaobaoOmniorderStoreSwitchstatusGetRequest 从 sync.Pool 获取 TaobaoOmniorderStoreSwitchstatusGetAPIRequest
func GetTaobaoOmniorderStoreSwitchstatusGetAPIRequest() *TaobaoOmniorderStoreSwitchstatusGetAPIRequest {
	return poolTaobaoOmniorderStoreSwitchstatusGetAPIRequest.Get().(*TaobaoOmniorderStoreSwitchstatusGetAPIRequest)
}

// ReleaseTaobaoOmniorderStoreSwitchstatusGetAPIRequest 将 TaobaoOmniorderStoreSwitchstatusGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoOmniorderStoreSwitchstatusGetAPIRequest(v *TaobaoOmniorderStoreSwitchstatusGetAPIRequest) {
	v.Reset()
	poolTaobaoOmniorderStoreSwitchstatusGetAPIRequest.Put(v)
}
