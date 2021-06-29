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
    bizType   int64
    // 延期时间
    endDate   string
    // 码
    code   string
    // 订单号
    outerId   string
    // 扩展
    attributeMap   string
    // 请求ID，调用方保证惟一
    requestId   string
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
func (r *TaobaoEticketMerchantMaDelayRequest) SetBizType(bizType int64) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

// BizType Getter
func (r TaobaoEticketMerchantMaDelayRequest) GetBizType() int64 {
    return r.bizType
}
// EndDate Setter
// 延期时间
func (r *TaobaoEticketMerchantMaDelayRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r TaobaoEticketMerchantMaDelayRequest) GetEndDate() string {
    return r.endDate
}
// Code Setter
// 码
func (r *TaobaoEticketMerchantMaDelayRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

// Code Getter
func (r TaobaoEticketMerchantMaDelayRequest) GetCode() string {
    return r.code
}
// OuterId Setter
// 订单号
func (r *TaobaoEticketMerchantMaDelayRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TaobaoEticketMerchantMaDelayRequest) GetOuterId() string {
    return r.outerId
}
// AttributeMap Setter
// 扩展
func (r *TaobaoEticketMerchantMaDelayRequest) SetAttributeMap(attributeMap string) error {
    r.attributeMap = attributeMap
    r.Set("attribute_map", attributeMap)
    return nil
}

// AttributeMap Getter
func (r TaobaoEticketMerchantMaDelayRequest) GetAttributeMap() string {
    return r.attributeMap
}
// RequestId Setter
// 请求ID，调用方保证惟一
func (r *TaobaoEticketMerchantMaDelayRequest) SetRequestId(requestId string) error {
    r.requestId = requestId
    r.Set("request_id", requestId)
    return nil
}

// RequestId Getter
func (r TaobaoEticketMerchantMaDelayRequest) GetRequestId() string {
    return r.requestId
}
