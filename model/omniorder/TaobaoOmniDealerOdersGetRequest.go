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
type TaobaoOmniDealerOdersGetRequest struct {
    model.Params
    // 主订单ID
    mainOrderId   int64
}

// 初始化TaobaoOmniDealerOdersGetRequest对象
func NewTaobaoOmniDealerOdersGetRequest() *TaobaoOmniDealerOdersGetRequest{
    return &TaobaoOmniDealerOdersGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniDealerOdersGetRequest) GetApiMethodName() string {
    return "taobao.omni.dealer.oders.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniDealerOdersGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainOrderId Setter
// 主订单ID
func (r *TaobaoOmniDealerOdersGetRequest) SetMainOrderId(mainOrderId int64) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

// MainOrderId Getter
func (r TaobaoOmniDealerOdersGetRequest) GetMainOrderId() int64 {
    return r.mainOrderId
}
