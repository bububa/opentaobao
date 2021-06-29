package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证重发回调接口 API请求
taobao.eticket.merchant.ma.resend

码商重发电子凭证回调接口
*/
type TaobaoEticketMerchantMaResendRequest struct {
    model.Params
    // 业务类型
    bizType   int64
    // 待重发的码列表
    isvMaList   []IsvMa
    // 业务id（订单号）
    outerId   string
    // 需要跟发码通知获取到的参数一致
    token   string
}

// 初始化TaobaoEticketMerchantMaResendRequest对象
func NewTaobaoEticketMerchantMaResendRequest() *TaobaoEticketMerchantMaResendRequest{
    return &TaobaoEticketMerchantMaResendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoEticketMerchantMaResendRequest) GetApiMethodName() string {
    return "taobao.eticket.merchant.ma.resend"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoEticketMerchantMaResendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizType Setter
// 业务类型
func (r *TaobaoEticketMerchantMaResendRequest) SetBizType(bizType int64) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

// BizType Getter
func (r TaobaoEticketMerchantMaResendRequest) GetBizType() int64 {
    return r.bizType
}
// IsvMaList Setter
// 待重发的码列表
func (r *TaobaoEticketMerchantMaResendRequest) SetIsvMaList(isvMaList []IsvMa) error {
    r.isvMaList = isvMaList
    r.Set("isv_ma_list", isvMaList)
    return nil
}

// IsvMaList Getter
func (r TaobaoEticketMerchantMaResendRequest) GetIsvMaList() []IsvMa {
    return r.isvMaList
}
// OuterId Setter
// 业务id（订单号）
func (r *TaobaoEticketMerchantMaResendRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TaobaoEticketMerchantMaResendRequest) GetOuterId() string {
    return r.outerId
}
// Token Setter
// 需要跟发码通知获取到的参数一致
func (r *TaobaoEticketMerchantMaResendRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r TaobaoEticketMerchantMaResendRequest) GetToken() string {
    return r.token
}
