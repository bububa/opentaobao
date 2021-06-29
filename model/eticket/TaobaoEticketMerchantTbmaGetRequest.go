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
type TaobaoEticketMerchantTbmaGetRequest struct {
    model.Params
    // 查询淘宝码请求
    queryTbMaCallbackReq   *QueryTbMaCallbackReq
}

// 初始化TaobaoEticketMerchantTbmaGetRequest对象
func NewTaobaoEticketMerchantTbmaGetRequest() *TaobaoEticketMerchantTbmaGetRequest{
    return &TaobaoEticketMerchantTbmaGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoEticketMerchantTbmaGetRequest) GetApiMethodName() string {
    return "taobao.eticket.merchant.tbma.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoEticketMerchantTbmaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryTbMaCallbackReq Setter
// 查询淘宝码请求
func (r *TaobaoEticketMerchantTbmaGetRequest) SetQueryTbMaCallbackReq(queryTbMaCallbackReq *QueryTbMaCallbackReq) error {
    r.queryTbMaCallbackReq = queryTbMaCallbackReq
    r.Set("query_tb_ma_callback_req", queryTbMaCallbackReq)
    return nil
}

// QueryTbMaCallbackReq Getter
func (r TaobaoEticketMerchantTbmaGetRequest) GetQueryTbMaCallbackReq() *QueryTbMaCallbackReq {
    return r.queryTbMaCallbackReq
}
