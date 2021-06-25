package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证核销接口 APIRequest
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

func NewTaobaoEticketMerchantMaConsumeRequest() *TaobaoEticketMerchantMaConsumeRequest{
    return &TaobaoEticketMerchantMaConsumeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoEticketMerchantMaConsumeRequest) GetApiMethodName() string {
    return "taobao.eticket.merchant.ma.consume"
}

func (r TaobaoEticketMerchantMaConsumeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoEticketMerchantMaConsumeRequest) SetBizType(bizType int64) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

func (r TaobaoEticketMerchantMaConsumeRequest) GetBizType() int64 {
    return r.bizType
}

func (r *TaobaoEticketMerchantMaConsumeRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r TaobaoEticketMerchantMaConsumeRequest) GetCode() string {
    return r.code
}

func (r *TaobaoEticketMerchantMaConsumeRequest) SetConsumeNum(consumeNum int64) error {
    r.consumeNum = consumeNum
    r.Set("consume_num", consumeNum)
    return nil
}

func (r TaobaoEticketMerchantMaConsumeRequest) GetConsumeNum() int64 {
    return r.consumeNum
}

func (r *TaobaoEticketMerchantMaConsumeRequest) SetIsvMaList(isvMaList []IsvMa) error {
    r.isvMaList = isvMaList
    r.Set("isv_ma_list", isvMaList)
    return nil
}

func (r TaobaoEticketMerchantMaConsumeRequest) GetIsvMaList() []IsvMa {
    return r.isvMaList
}

func (r *TaobaoEticketMerchantMaConsumeRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r TaobaoEticketMerchantMaConsumeRequest) GetOuterId() string {
    return r.outerId
}

func (r *TaobaoEticketMerchantMaConsumeRequest) SetPosId(posId string) error {
    r.posId = posId
    r.Set("pos_id", posId)
    return nil
}

func (r TaobaoEticketMerchantMaConsumeRequest) GetPosId() string {
    return r.posId
}

func (r *TaobaoEticketMerchantMaConsumeRequest) SetSerialNum(serialNum string) error {
    r.serialNum = serialNum
    r.Set("serial_num", serialNum)
    return nil
}

func (r TaobaoEticketMerchantMaConsumeRequest) GetSerialNum() string {
    return r.serialNum
}

func (r *TaobaoEticketMerchantMaConsumeRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r TaobaoEticketMerchantMaConsumeRequest) GetToken() string {
    return r.token
}

