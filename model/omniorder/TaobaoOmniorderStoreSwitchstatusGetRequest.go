package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
switchstatus.get API请求
taobao.omniorder.store.switchstatus.get

查询门店发货、门店自提状态
*/
type TaobaoOmniorderStoreSwitchstatusGetRequest struct {
    model.Params
    // 门店ID
    _storeId   int64
    // 卖家ID
    _sellerId   int64
}

// 初始化TaobaoOmniorderStoreSwitchstatusGetRequest对象
func NewTaobaoOmniorderStoreSwitchstatusGetRequest() *TaobaoOmniorderStoreSwitchstatusGetRequest{
    return &TaobaoOmniorderStoreSwitchstatusGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreSwitchstatusGetRequest) GetApiMethodName() string {
    return "taobao.omniorder.store.switchstatus.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreSwitchstatusGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 门店ID
func (r *TaobaoOmniorderStoreSwitchstatusGetRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoOmniorderStoreSwitchstatusGetRequest) GetStoreId() int64 {
    return r._storeId
}
// SellerId Setter
// 卖家ID
func (r *TaobaoOmniorderStoreSwitchstatusGetRequest) SetSellerId(_sellerId int64) error {
    r._sellerId = _sellerId
    r.Set("seller_id", _sellerId)
    return nil
}

// SellerId Getter
func (r TaobaoOmniorderStoreSwitchstatusGetRequest) GetSellerId() int64 {
    return r._sellerId
}
