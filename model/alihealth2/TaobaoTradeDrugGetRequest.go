package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商家未确认订单列表 API请求
taobao.trade.drug.get

可以按商家或是店铺维度的进行查询买家付款卖家未确认订单，一次返回不大于20条订单
*/
type TaobaoTradeDrugGetRequest struct {
    model.Params
    // 店铺id
    _storeId   int64
    // true-商家下所有店铺的待确认订单, false—指定店铺的订单
    _isAll   bool
    // 返回记录数，超过20按20条返回数据
    _maxSize   int64
}

// 初始化TaobaoTradeDrugGetRequest对象
func NewTaobaoTradeDrugGetRequest() *TaobaoTradeDrugGetRequest{
    return &TaobaoTradeDrugGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTradeDrugGetRequest) GetApiMethodName() string {
    return "taobao.trade.drug.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTradeDrugGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 店铺id
func (r *TaobaoTradeDrugGetRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoTradeDrugGetRequest) GetStoreId() int64 {
    return r._storeId
}
// IsAll Setter
// true-商家下所有店铺的待确认订单, false—指定店铺的订单
func (r *TaobaoTradeDrugGetRequest) SetIsAll(_isAll bool) error {
    r._isAll = _isAll
    r.Set("is_all", _isAll)
    return nil
}

// IsAll Getter
func (r TaobaoTradeDrugGetRequest) GetIsAll() bool {
    return r._isAll
}
// MaxSize Setter
// 返回记录数，超过20按20条返回数据
func (r *TaobaoTradeDrugGetRequest) SetMaxSize(_maxSize int64) error {
    r._maxSize = _maxSize
    r.Set("max_size", _maxSize)
    return nil
}

// MaxSize Getter
func (r TaobaoTradeDrugGetRequest) GetMaxSize() int64 {
    return r._maxSize
}
