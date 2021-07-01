package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证冲正接口 API请求
taobao.eticket.merchant.ma.reverse

电子凭证平台冲正接口
*/
type TaobaoEticketMerchantMaReverseAPIRequest struct {
    model.Params
    // 业务类型
    _bizType   int64
    // 码值
    _code   string
    // 业务id（订单号）
    _outerId   string
    // 机具编号，如果核销时有则必传
    _posId   string
    // 冲正份数，需要与核销份数一致
    _reverseNum   int64
    // 需要冲正的核销序列号
    _serialNum   string
    // 需要跟发码通知获取到的参数一致
    _token   string
}

// 初始化TaobaoEticketMerchantMaReverseAPIRequest对象
func NewTaobaoEticketMerchantMaReverseRequest() *TaobaoEticketMerchantMaReverseAPIRequest{
    return &TaobaoEticketMerchantMaReverseAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoEticketMerchantMaReverseAPIRequest) GetApiMethodName() string {
    return "taobao.eticket.merchant.ma.reverse"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoEticketMerchantMaReverseAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizType Setter
// 业务类型
func (r *TaobaoEticketMerchantMaReverseAPIRequest) SetBizType(_bizType int64) error {
    r._bizType = _bizType
    r.Set("biz_type", _bizType)
    return nil
}

// BizType Getter
func (r TaobaoEticketMerchantMaReverseAPIRequest) GetBizType() int64 {
    return r._bizType
}
// Code Setter
// 码值
func (r *TaobaoEticketMerchantMaReverseAPIRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r TaobaoEticketMerchantMaReverseAPIRequest) GetCode() string {
    return r._code
}
// OuterId Setter
// 业务id（订单号）
func (r *TaobaoEticketMerchantMaReverseAPIRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r TaobaoEticketMerchantMaReverseAPIRequest) GetOuterId() string {
    return r._outerId
}
// PosId Setter
// 机具编号，如果核销时有则必传
func (r *TaobaoEticketMerchantMaReverseAPIRequest) SetPosId(_posId string) error {
    r._posId = _posId
    r.Set("pos_id", _posId)
    return nil
}

// PosId Getter
func (r TaobaoEticketMerchantMaReverseAPIRequest) GetPosId() string {
    return r._posId
}
// ReverseNum Setter
// 冲正份数，需要与核销份数一致
func (r *TaobaoEticketMerchantMaReverseAPIRequest) SetReverseNum(_reverseNum int64) error {
    r._reverseNum = _reverseNum
    r.Set("reverse_num", _reverseNum)
    return nil
}

// ReverseNum Getter
func (r TaobaoEticketMerchantMaReverseAPIRequest) GetReverseNum() int64 {
    return r._reverseNum
}
// SerialNum Setter
// 需要冲正的核销序列号
func (r *TaobaoEticketMerchantMaReverseAPIRequest) SetSerialNum(_serialNum string) error {
    r._serialNum = _serialNum
    r.Set("serial_num", _serialNum)
    return nil
}

// SerialNum Getter
func (r TaobaoEticketMerchantMaReverseAPIRequest) GetSerialNum() string {
    return r._serialNum
}
// Token Setter
// 需要跟发码通知获取到的参数一致
func (r *TaobaoEticketMerchantMaReverseAPIRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r TaobaoEticketMerchantMaReverseAPIRequest) GetToken() string {
    return r._token
}
