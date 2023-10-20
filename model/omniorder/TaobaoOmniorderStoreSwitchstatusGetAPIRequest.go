package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniorderstoreswitchstatusgetAPIRequest switchstatus.get API请求
// taobao.omniorder.store.switchstatus.get
//
// 查询门店发货、门店自提状态
type TaobaoomniorderstoreswitchstatusgetAPIRequest struct {
	model.Params
	// 门店ID
	_storeId int64
	// 卖家ID
	_sellerId int64
}

// NewTaobaoomniorderstoreswitchstatusgetRequest 初始化TaobaoomniorderstoreswitchstatusgetAPIRequest对象
func NewTaobaoomniorderstoreswitchstatusgetRequest() *TaobaoomniorderstoreswitchstatusgetAPIRequest {
	return &TaobaoomniorderstoreswitchstatusgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoomniorderstoreswitchstatusgetAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.store.switchstatus.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoomniorderstoreswitchstatusgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoomniorderstoreswitchstatusgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *TaobaoomniorderstoreswitchstatusgetAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoomniorderstoreswitchstatusgetAPIRequest) GetStoreId() int64 {
	return r._storeId
}

// SetSellerId is SellerId Setter
// 卖家ID
func (r *TaobaoomniorderstoreswitchstatusgetAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// GetSellerId SellerId Getter
func (r TaobaoomniorderstoreswitchstatusgetAPIRequest) GetSellerId() int64 {
	return r._sellerId
}
