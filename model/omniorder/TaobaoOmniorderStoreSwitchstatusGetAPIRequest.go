package omniorder

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreSwitchstatusGetAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.store.switchstatus.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreSwitchstatusGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StoreId Setter
// 门店ID
func (r *TaobaoOmniorderStoreSwitchstatusGetAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r TaobaoOmniorderStoreSwitchstatusGetAPIRequest) GetStoreId() int64 {
	return r._storeId
}

// Set is SellerId Setter
// 卖家ID
func (r *TaobaoOmniorderStoreSwitchstatusGetAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// Get SellerId Getter
func (r TaobaoOmniorderStoreSwitchstatusGetAPIRequest) GetSellerId() int64 {
	return r._sellerId
}
