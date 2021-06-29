package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
码商发码成功回调接口 API请求
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

// 初始化TaobaoEticketMerchantMaSendRequest对象
func NewTaobaoEticketMerchantMaSendRequest() *TaobaoEticketMerchantMaSendRequest{
    return &TaobaoEticketMerchantMaSendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoEticketMerchantMaSendRequest) GetApiMethodName() string {
    return "taobao.eticket.merchant.ma.send"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoEticketMerchantMaSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizType Setter
// 业务类型
func (r *TaobaoEticketMerchantMaSendRequest) SetBizType(bizType int64) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

// BizType Getter
func (r TaobaoEticketMerchantMaSendRequest) GetBizType() int64 {
    return r.bizType
}
// IsvMaList Setter
// 需要发送的码列表
func (r *TaobaoEticketMerchantMaSendRequest) SetIsvMaList(isvMaList []IsvMa) error {
    r.isvMaList = isvMaList
    r.Set("isv_ma_list", isvMaList)
    return nil
}

// IsvMaList Getter
func (r TaobaoEticketMerchantMaSendRequest) GetIsvMaList() []IsvMa {
    return r.isvMaList
}
// OuterId Setter
// 业务id（订单号）
func (r *TaobaoEticketMerchantMaSendRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TaobaoEticketMerchantMaSendRequest) GetOuterId() string {
    return r.outerId
}
// Token Setter
// 需要跟发码通知获取到的参数一致
func (r *TaobaoEticketMerchantMaSendRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r TaobaoEticketMerchantMaSendRequest) GetToken() string {
    return r.token
}
