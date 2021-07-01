package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTradeDrugGetAPIRequest
查询商家未确认订单列表 API请求
taobao.trade.drug.get

可以按商家或是店铺维度的进行查询买家付款卖家未确认订单，一次返回不大于20条订单 */
type TaobaoTradeDrugGetAPIRequest struct {
	model.Params
	// 店铺id
	_storeId int64
	// true-商家下所有店铺的待确认订单, false—指定店铺的订单
	_isAll bool
	// 返回记录数，超过20按20条返回数据
	_maxSize int64
}

// NewTaobaoTradeDrugGetRequest 初始化TaobaoTradeDrugGetAPIRequest对象
func NewTaobaoTradeDrugGetRequest() *TaobaoTradeDrugGetAPIRequest {
	return &TaobaoTradeDrugGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTradeDrugGetAPIRequest) GetApiMethodName() string {
	return "taobao.trade.drug.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTradeDrugGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StoreId Setter
// 店铺id
func (r *TaobaoTradeDrugGetAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r TaobaoTradeDrugGetAPIRequest) GetStoreId() int64 {
	return r._storeId
}

// Set is IsAll Setter
// true-商家下所有店铺的待确认订单, false—指定店铺的订单
func (r *TaobaoTradeDrugGetAPIRequest) SetIsAll(_isAll bool) error {
	r._isAll = _isAll
	r.Set("is_all", _isAll)
	return nil
}

// Get IsAll Getter
func (r TaobaoTradeDrugGetAPIRequest) GetIsAll() bool {
	return r._isAll
}

// Set is MaxSize Setter
// 返回记录数，超过20按20条返回数据
func (r *TaobaoTradeDrugGetAPIRequest) SetMaxSize(_maxSize int64) error {
	r._maxSize = _maxSize
	r.Set("max_size", _maxSize)
	return nil
}

// Get MaxSize Getter
func (r TaobaoTradeDrugGetAPIRequest) GetMaxSize() int64 {
	return r._maxSize
}
