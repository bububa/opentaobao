package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证核销接口 API请求
taobao.eticket.merchant.ma.consume

电子凭证核销接口
*/
type TaobaoEticketMerchantMaConsumeAPIRequest struct {
    model.Params
    // 业务类型
    _bizType   int64
    // 需要被核销的码
    _code   string
    // 核销份数
    _consumeNum   int64
    // 核销后换码的码列表
    _isvMaList   []IsvMa
    // 业务id（订单号）
    _outerId   string
    // 机具编号
    _posId   string
    // 核销序列号，需要保证唯一
    _serialNum   string
    // 需要跟发码通知获取到的参数一致
    _token   string
}

// 初始化TaobaoEticketMerchantMaConsumeAPIRequest对象
func NewTaobaoEticketMerchantMaConsumeRequest() *TaobaoEticketMerchantMaConsumeAPIRequest{
    return &TaobaoEticketMerchantMaConsumeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoEticketMerchantMaConsumeAPIRequest) GetApiMethodName() string {
    return "taobao.eticket.merchant.ma.consume"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoEticketMerchantMaConsumeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizType Setter
// 业务类型
func (r *TaobaoEticketMerchantMaConsumeAPIRequest) SetBizType(_bizType int64) error {
    r._bizType = _bizType
    r.Set("biz_type", _bizType)
    return nil
}

// BizType Getter
func (r TaobaoEticketMerchantMaConsumeAPIRequest) GetBizType() int64 {
    return r._bizType
}
// Code Setter
// 需要被核销的码
func (r *TaobaoEticketMerchantMaConsumeAPIRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r TaobaoEticketMerchantMaConsumeAPIRequest) GetCode() string {
    return r._code
}
// ConsumeNum Setter
// 核销份数
func (r *TaobaoEticketMerchantMaConsumeAPIRequest) SetConsumeNum(_consumeNum int64) error {
    r._consumeNum = _consumeNum
    r.Set("consume_num", _consumeNum)
    return nil
}

// ConsumeNum Getter
func (r TaobaoEticketMerchantMaConsumeAPIRequest) GetConsumeNum() int64 {
    return r._consumeNum
}
// IsvMaList Setter
// 核销后换码的码列表
func (r *TaobaoEticketMerchantMaConsumeAPIRequest) SetIsvMaList(_isvMaList []IsvMa) error {
    r._isvMaList = _isvMaList
    r.Set("isv_ma_list", _isvMaList)
    return nil
}

// IsvMaList Getter
func (r TaobaoEticketMerchantMaConsumeAPIRequest) GetIsvMaList() []IsvMa {
    return r._isvMaList
}
// OuterId Setter
// 业务id（订单号）
func (r *TaobaoEticketMerchantMaConsumeAPIRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r TaobaoEticketMerchantMaConsumeAPIRequest) GetOuterId() string {
    return r._outerId
}
// PosId Setter
// 机具编号
func (r *TaobaoEticketMerchantMaConsumeAPIRequest) SetPosId(_posId string) error {
    r._posId = _posId
    r.Set("pos_id", _posId)
    return nil
}

// PosId Getter
func (r TaobaoEticketMerchantMaConsumeAPIRequest) GetPosId() string {
    return r._posId
}
// SerialNum Setter
// 核销序列号，需要保证唯一
func (r *TaobaoEticketMerchantMaConsumeAPIRequest) SetSerialNum(_serialNum string) error {
    r._serialNum = _serialNum
    r.Set("serial_num", _serialNum)
    return nil
}

// SerialNum Getter
func (r TaobaoEticketMerchantMaConsumeAPIRequest) GetSerialNum() string {
    return r._serialNum
}
// Token Setter
// 需要跟发码通知获取到的参数一致
func (r *TaobaoEticketMerchantMaConsumeAPIRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r TaobaoEticketMerchantMaConsumeAPIRequest) GetToken() string {
    return r._token
}
