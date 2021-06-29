package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
凭证延期 API请求
taobao.eticket.merchant.ma.delay

订单延期
*/
type TaobaoEticketMerchantMaDelayRequest struct {
    model.Params
    // 业务类型
    _bizType   int64
    // 延期时间
    _endDate   string
    // 码
    _code   string
    // 订单号
    _outerId   string
    // 扩展
    _attributeMap   string
    // 请求ID，调用方保证惟一
    _requestId   string
}

// 初始化TaobaoEticketMerchantMaDelayRequest对象
func NewTaobaoEticketMerchantMaDelayRequest() *TaobaoEticketMerchantMaDelayRequest{
    return &TaobaoEticketMerchantMaDelayRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoEticketMerchantMaDelayRequest) GetApiMethodName() string {
    return "taobao.eticket.merchant.ma.delay"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoEticketMerchantMaDelayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizType Setter
// 业务类型
func (r *TaobaoEticketMerchantMaDelayRequest) SetBizType(_bizType int64) error {
    r._bizType = _bizType
    r.Set("biz_type", _bizType)
    return nil
}

// BizType Getter
func (r TaobaoEticketMerchantMaDelayRequest) GetBizType() int64 {
    return r._bizType
}
// EndDate Setter
// 延期时间
func (r *TaobaoEticketMerchantMaDelayRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r TaobaoEticketMerchantMaDelayRequest) GetEndDate() string {
    return r._endDate
}
// Code Setter
// 码
func (r *TaobaoEticketMerchantMaDelayRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r TaobaoEticketMerchantMaDelayRequest) GetCode() string {
    return r._code
}
// OuterId Setter
// 订单号
func (r *TaobaoEticketMerchantMaDelayRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r TaobaoEticketMerchantMaDelayRequest) GetOuterId() string {
    return r._outerId
}
// AttributeMap Setter
// 扩展
func (r *TaobaoEticketMerchantMaDelayRequest) SetAttributeMap(_attributeMap string) error {
    r._attributeMap = _attributeMap
    r.Set("attribute_map", _attributeMap)
    return nil
}

// AttributeMap Getter
func (r TaobaoEticketMerchantMaDelayRequest) GetAttributeMap() string {
    return r._attributeMap
}
// RequestId Setter
// 请求ID，调用方保证惟一
func (r *TaobaoEticketMerchantMaDelayRequest) SetRequestId(_requestId string) error {
    r._requestId = _requestId
    r.Set("request_id", _requestId)
    return nil
}

// RequestId Getter
func (r TaobaoEticketMerchantMaDelayRequest) GetRequestId() string {
    return r._requestId
}
