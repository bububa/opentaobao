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
type TaobaoEticketMerchantMaConsumeRequest struct {
    model.Params
    // 业务类型
    bizType   int64
    // 需要被核销的码
    code   string
    // 核销份数
    consumeNum   int64
    // 核销后换码的码列表
    isvMaList   []IsvMa
    // 业务id（订单号）
    outerId   string
    // 机具编号
    posId   string
    // 核销序列号，需要保证唯一
    serialNum   string
    // 需要跟发码通知获取到的参数一致
    token   string
}

// 初始化TaobaoEticketMerchantMaConsumeRequest对象
func NewTaobaoEticketMerchantMaConsumeRequest() *TaobaoEticketMerchantMaConsumeRequest{
    return &TaobaoEticketMerchantMaConsumeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoEticketMerchantMaConsumeRequest) GetApiMethodName() string {
    return "taobao.eticket.merchant.ma.consume"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoEticketMerchantMaConsumeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizType Setter
// 业务类型
func (r *TaobaoEticketMerchantMaConsumeRequest) SetBizType(bizType int64) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

// BizType Getter
func (r TaobaoEticketMerchantMaConsumeRequest) GetBizType() int64 {
    return r.bizType
}
// Code Setter
// 需要被核销的码
func (r *TaobaoEticketMerchantMaConsumeRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

// Code Getter
func (r TaobaoEticketMerchantMaConsumeRequest) GetCode() string {
    return r.code
}
// ConsumeNum Setter
// 核销份数
func (r *TaobaoEticketMerchantMaConsumeRequest) SetConsumeNum(consumeNum int64) error {
    r.consumeNum = consumeNum
    r.Set("consume_num", consumeNum)
    return nil
}

// ConsumeNum Getter
func (r TaobaoEticketMerchantMaConsumeRequest) GetConsumeNum() int64 {
    return r.consumeNum
}
// IsvMaList Setter
// 核销后换码的码列表
func (r *TaobaoEticketMerchantMaConsumeRequest) SetIsvMaList(isvMaList []IsvMa) error {
    r.isvMaList = isvMaList
    r.Set("isv_ma_list", isvMaList)
    return nil
}

// IsvMaList Getter
func (r TaobaoEticketMerchantMaConsumeRequest) GetIsvMaList() []IsvMa {
    return r.isvMaList
}
// OuterId Setter
// 业务id（订单号）
func (r *TaobaoEticketMerchantMaConsumeRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TaobaoEticketMerchantMaConsumeRequest) GetOuterId() string {
    return r.outerId
}
// PosId Setter
// 机具编号
func (r *TaobaoEticketMerchantMaConsumeRequest) SetPosId(posId string) error {
    r.posId = posId
    r.Set("pos_id", posId)
    return nil
}

// PosId Getter
func (r TaobaoEticketMerchantMaConsumeRequest) GetPosId() string {
    return r.posId
}
// SerialNum Setter
// 核销序列号，需要保证唯一
func (r *TaobaoEticketMerchantMaConsumeRequest) SetSerialNum(serialNum string) error {
    r.serialNum = serialNum
    r.Set("serial_num", serialNum)
    return nil
}

// SerialNum Getter
func (r TaobaoEticketMerchantMaConsumeRequest) GetSerialNum() string {
    return r.serialNum
}
// Token Setter
// 需要跟发码通知获取到的参数一致
func (r *TaobaoEticketMerchantMaConsumeRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r TaobaoEticketMerchantMaConsumeRequest) GetToken() string {
    return r.token
}
