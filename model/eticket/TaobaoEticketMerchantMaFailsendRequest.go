package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
码商发码失败回调接口 API请求
taobao.eticket.merchant.ma.failsend

针对一次发码通知，码商无法完成发码，则可以通过此接口告知电子凭证
*/
type TaobaoEticketMerchantMaFailsendRequest struct {
    model.Params
    // 业务id（订单号）
    outerId   string
    // 错误原因码
    subErrCode   string
    // 错误码描述
    subErrMsg   string
    // 需要与发码通知获取的值一致
    token   string
    // 业务类型
    bizType   int64
}

// 初始化TaobaoEticketMerchantMaFailsendRequest对象
func NewTaobaoEticketMerchantMaFailsendRequest() *TaobaoEticketMerchantMaFailsendRequest{
    return &TaobaoEticketMerchantMaFailsendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoEticketMerchantMaFailsendRequest) GetApiMethodName() string {
    return "taobao.eticket.merchant.ma.failsend"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoEticketMerchantMaFailsendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterId Setter
// 业务id（订单号）
func (r *TaobaoEticketMerchantMaFailsendRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TaobaoEticketMerchantMaFailsendRequest) GetOuterId() string {
    return r.outerId
}
// SubErrCode Setter
// 错误原因码
func (r *TaobaoEticketMerchantMaFailsendRequest) SetSubErrCode(subErrCode string) error {
    r.subErrCode = subErrCode
    r.Set("sub_err_code", subErrCode)
    return nil
}

// SubErrCode Getter
func (r TaobaoEticketMerchantMaFailsendRequest) GetSubErrCode() string {
    return r.subErrCode
}
// SubErrMsg Setter
// 错误码描述
func (r *TaobaoEticketMerchantMaFailsendRequest) SetSubErrMsg(subErrMsg string) error {
    r.subErrMsg = subErrMsg
    r.Set("sub_err_msg", subErrMsg)
    return nil
}

// SubErrMsg Getter
func (r TaobaoEticketMerchantMaFailsendRequest) GetSubErrMsg() string {
    return r.subErrMsg
}
// Token Setter
// 需要与发码通知获取的值一致
func (r *TaobaoEticketMerchantMaFailsendRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r TaobaoEticketMerchantMaFailsendRequest) GetToken() string {
    return r.token
}
// BizType Setter
// 业务类型
func (r *TaobaoEticketMerchantMaFailsendRequest) SetBizType(bizType int64) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

// BizType Getter
func (r TaobaoEticketMerchantMaFailsendRequest) GetBizType() int64 {
    return r.bizType
}
