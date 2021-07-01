package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证核销前校验接口 API请求
taobao.eticket.merchant.ma.available

商家验码之前的调用接口，用来判断是否可以进行核销操作
*/
type TaobaoEticketMerchantMaAvailableAPIRequest struct {
    model.Params
    // 业务类型
    _bizType   int64
    // 需要被核销的码
    _code   string
    // 核销份数
    _consumeNum   int64
    // 业务id（订单号）
    _outerId   string
    // 机具编号
    _posId   string
    // 核销序列号，需要保证唯一
    _serialNum   string
    // 需要跟发码通知获取到的参数一致
    _token   string
}

// 初始化TaobaoEticketMerchantMaAvailableAPIRequest对象
func NewTaobaoEticketMerchantMaAvailableRequest() *TaobaoEticketMerchantMaAvailableAPIRequest{
    return &TaobaoEticketMerchantMaAvailableAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoEticketMerchantMaAvailableAPIRequest) GetApiMethodName() string {
    return "taobao.eticket.merchant.ma.available"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoEticketMerchantMaAvailableAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizType Setter
// 业务类型
func (r *TaobaoEticketMerchantMaAvailableAPIRequest) SetBizType(_bizType int64) error {
    r._bizType = _bizType
    r.Set("biz_type", _bizType)
    return nil
}

// BizType Getter
func (r TaobaoEticketMerchantMaAvailableAPIRequest) GetBizType() int64 {
    return r._bizType
}
// Code Setter
// 需要被核销的码
func (r *TaobaoEticketMerchantMaAvailableAPIRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r TaobaoEticketMerchantMaAvailableAPIRequest) GetCode() string {
    return r._code
}
// ConsumeNum Setter
// 核销份数
func (r *TaobaoEticketMerchantMaAvailableAPIRequest) SetConsumeNum(_consumeNum int64) error {
    r._consumeNum = _consumeNum
    r.Set("consume_num", _consumeNum)
    return nil
}

// ConsumeNum Getter
func (r TaobaoEticketMerchantMaAvailableAPIRequest) GetConsumeNum() int64 {
    return r._consumeNum
}
// OuterId Setter
// 业务id（订单号）
func (r *TaobaoEticketMerchantMaAvailableAPIRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r TaobaoEticketMerchantMaAvailableAPIRequest) GetOuterId() string {
    return r._outerId
}
// PosId Setter
// 机具编号
func (r *TaobaoEticketMerchantMaAvailableAPIRequest) SetPosId(_posId string) error {
    r._posId = _posId
    r.Set("pos_id", _posId)
    return nil
}

// PosId Getter
func (r TaobaoEticketMerchantMaAvailableAPIRequest) GetPosId() string {
    return r._posId
}
// SerialNum Setter
// 核销序列号，需要保证唯一
func (r *TaobaoEticketMerchantMaAvailableAPIRequest) SetSerialNum(_serialNum string) error {
    r._serialNum = _serialNum
    r.Set("serial_num", _serialNum)
    return nil
}

// SerialNum Getter
func (r TaobaoEticketMerchantMaAvailableAPIRequest) GetSerialNum() string {
    return r._serialNum
}
// Token Setter
// 需要跟发码通知获取到的参数一致
func (r *TaobaoEticketMerchantMaAvailableAPIRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r TaobaoEticketMerchantMaAvailableAPIRequest) GetToken() string {
    return r._token
}
