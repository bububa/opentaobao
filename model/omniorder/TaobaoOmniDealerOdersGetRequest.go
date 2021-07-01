package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单笔全渠道经销商订单的详细信息 API请求
taobao.omni.dealer.oders.get

全渠道经销商获取单笔订单的详细信息
*/
type TaobaoOmniDealerOdersGetAPIRequest struct {
    model.Params
    // 主订单ID
    _mainOrderId   int64
}

// 初始化TaobaoOmniDealerOdersGetAPIRequest对象
func NewTaobaoOmniDealerOdersGetRequest() *TaobaoOmniDealerOdersGetAPIRequest{
    return &TaobaoOmniDealerOdersGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniDealerOdersGetAPIRequest) GetApiMethodName() string {
    return "taobao.omni.dealer.oders.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniDealerOdersGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainOrderId Setter
// 主订单ID
func (r *TaobaoOmniDealerOdersGetAPIRequest) SetMainOrderId(_mainOrderId int64) error {
    r._mainOrderId = _mainOrderId
    r.Set("main_order_id", _mainOrderId)
    return nil
}

// MainOrderId Getter
func (r TaobaoOmniDealerOdersGetAPIRequest) GetMainOrderId() int64 {
    return r._mainOrderId
}
