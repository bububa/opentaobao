package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店自送发货 API请求
taobao.omniorder.dtd.consign

该接口触发门店自送发货，推进淘系订单状态为发货，为消费者发送核销码短信，并将物流信息写入订单
*/
type TaobaoOmniorderDtdConsignRequest struct {
    model.Params
    // 淘宝订单主订单号
    _mainOrderId   int64
    // 发货对应的商户中心门店ID
    _storeId   int64
}

// 初始化TaobaoOmniorderDtdConsignRequest对象
func NewTaobaoOmniorderDtdConsignRequest() *TaobaoOmniorderDtdConsignRequest{
    return &TaobaoOmniorderDtdConsignRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderDtdConsignRequest) GetApiMethodName() string {
    return "taobao.omniorder.dtd.consign"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderDtdConsignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainOrderId Setter
// 淘宝订单主订单号
func (r *TaobaoOmniorderDtdConsignRequest) SetMainOrderId(_mainOrderId int64) error {
    r._mainOrderId = _mainOrderId
    r.Set("main_order_id", _mainOrderId)
    return nil
}

// MainOrderId Getter
func (r TaobaoOmniorderDtdConsignRequest) GetMainOrderId() int64 {
    return r._mainOrderId
}
// StoreId Setter
// 发货对应的商户中心门店ID
func (r *TaobaoOmniorderDtdConsignRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoOmniorderDtdConsignRequest) GetStoreId() int64 {
    return r._storeId
}
