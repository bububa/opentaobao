package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotradedruggetAPIRequest 查询商家未确认订单列表 API请求
// taobao.trade.drug.get
//
// 可以按商家或是店铺维度的进行查询买家付款卖家未确认订单，一次返回不大于20条订单
type TaobaotradedruggetAPIRequest struct {
	model.Params
	// 店铺id
	_storeId int64
	// 返回记录数，超过20按20条返回数据
	_maxSize int64
	// true-商家下所有店铺的待确认订单, false—指定店铺的订单
	_isAll bool
}

// NewTaobaotradedruggetRequest 初始化TaobaotradedruggetAPIRequest对象
func NewTaobaotradedruggetRequest() *TaobaotradedruggetAPIRequest {
	return &TaobaotradedruggetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotradedruggetAPIRequest) GetApiMethodName() string {
	return "taobao.trade.drug.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotradedruggetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotradedruggetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 店铺id
func (r *TaobaotradedruggetAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaotradedruggetAPIRequest) GetStoreId() int64 {
	return r._storeId
}

// SetMaxSize is MaxSize Setter
// 返回记录数，超过20按20条返回数据
func (r *TaobaotradedruggetAPIRequest) SetMaxSize(_maxSize int64) error {
	r._maxSize = _maxSize
	r.Set("max_size", _maxSize)
	return nil
}

// GetMaxSize MaxSize Getter
func (r TaobaotradedruggetAPIRequest) GetMaxSize() int64 {
	return r._maxSize
}

// SetIsAll is IsAll Setter
// true-商家下所有店铺的待确认订单, false—指定店铺的订单
func (r *TaobaotradedruggetAPIRequest) SetIsAll(_isAll bool) error {
	r._isAll = _isAll
	r.Set("is_all", _isAll)
	return nil
}

// GetIsAll IsAll Getter
func (r TaobaotradedruggetAPIRequest) GetIsAll() bool {
	return r._isAll
}
