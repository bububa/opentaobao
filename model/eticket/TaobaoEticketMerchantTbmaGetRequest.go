package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
码商查询淘宝码接口 API请求
taobao.eticket.merchant.tbma.get

码商查询淘宝码接口
*/
type TaobaoEticketMerchantTbmaGetAPIRequest struct {
    model.Params
    // 查询淘宝码请求
    _queryTbMaCallbackReq   *QueryTbMaCallbackReq
}

// 初始化TaobaoEticketMerchantTbmaGetAPIRequest对象
func NewTaobaoEticketMerchantTbmaGetRequest() *TaobaoEticketMerchantTbmaGetAPIRequest{
    return &TaobaoEticketMerchantTbmaGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoEticketMerchantTbmaGetAPIRequest) GetApiMethodName() string {
    return "taobao.eticket.merchant.tbma.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoEticketMerchantTbmaGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryTbMaCallbackReq Setter
// 查询淘宝码请求
func (r *TaobaoEticketMerchantTbmaGetAPIRequest) SetQueryTbMaCallbackReq(_queryTbMaCallbackReq *QueryTbMaCallbackReq) error {
    r._queryTbMaCallbackReq = _queryTbMaCallbackReq
    r.Set("query_tb_ma_callback_req", _queryTbMaCallbackReq)
    return nil
}

// QueryTbMaCallbackReq Getter
func (r TaobaoEticketMerchantTbmaGetAPIRequest) GetQueryTbMaCallbackReq() *QueryTbMaCallbackReq {
    return r._queryTbMaCallbackReq
}
