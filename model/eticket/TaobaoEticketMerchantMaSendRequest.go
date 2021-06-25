package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
码商发码成功回调接口 APIRequest
taobao.eticket.merchant.ma.send

码商发码成功回调接口
*/
type TaobaoEticketMerchantMaSendRequest struct {
    model.Params

    // 业务类型
    bizType   int64 

    // 需要发送的码列表
    isvMaList   []IsvMa 

    // 业务id（订单号）
    outerId   string 

    // 需要跟发码通知获取到的参数一致
    token   string 

}

func NewTaobaoEticketMerchantMaSendRequest() *TaobaoEticketMerchantMaSendRequest{
    return &TaobaoEticketMerchantMaSendRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoEticketMerchantMaSendRequest) GetApiMethodName() string {
    return "taobao.eticket.merchant.ma.send"
}

func (r TaobaoEticketMerchantMaSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoEticketMerchantMaSendRequest) SetBizType(bizType int64) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

func (r TaobaoEticketMerchantMaSendRequest) GetBizType() int64 {
    return r.bizType
}

func (r *TaobaoEticketMerchantMaSendRequest) SetIsvMaList(isvMaList []IsvMa) error {
    r.isvMaList = isvMaList
    r.Set("isv_ma_list", isvMaList)
    return nil
}

func (r TaobaoEticketMerchantMaSendRequest) GetIsvMaList() []IsvMa {
    return r.isvMaList
}

func (r *TaobaoEticketMerchantMaSendRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r TaobaoEticketMerchantMaSendRequest) GetOuterId() string {
    return r.outerId
}

func (r *TaobaoEticketMerchantMaSendRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r TaobaoEticketMerchantMaSendRequest) GetToken() string {
    return r.token
}

