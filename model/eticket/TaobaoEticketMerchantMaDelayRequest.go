package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
凭证延期 APIRequest
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

func NewTaobaoEticketMerchantMaDelayRequest() *TaobaoEticketMerchantMaDelayRequest{
    return &TaobaoEticketMerchantMaDelayRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoEticketMerchantMaDelayRequest) GetApiMethodName() string {
    return "taobao.eticket.merchant.ma.delay"
}

func (r TaobaoEticketMerchantMaDelayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoEticketMerchantMaDelayRequest) SetBizType(bizType int64) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

func (r TaobaoEticketMerchantMaDelayRequest) GetBizType() int64 {
    return r.bizType
}

func (r *TaobaoEticketMerchantMaDelayRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r TaobaoEticketMerchantMaDelayRequest) GetEndDate() string {
    return r.endDate
}

func (r *TaobaoEticketMerchantMaDelayRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r TaobaoEticketMerchantMaDelayRequest) GetCode() string {
    return r.code
}

func (r *TaobaoEticketMerchantMaDelayRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r TaobaoEticketMerchantMaDelayRequest) GetOuterId() string {
    return r.outerId
}

func (r *TaobaoEticketMerchantMaDelayRequest) SetAttributeMap(attributeMap string) error {
    r.attributeMap = attributeMap
    r.Set("attribute_map", attributeMap)
    return nil
}

func (r TaobaoEticketMerchantMaDelayRequest) GetAttributeMap() string {
    return r.attributeMap
}

func (r *TaobaoEticketMerchantMaDelayRequest) SetRequestId(requestId string) error {
    r.requestId = requestId
    r.Set("request_id", requestId)
    return nil
}

func (r TaobaoEticketMerchantMaDelayRequest) GetRequestId() string {
    return r.requestId
}

