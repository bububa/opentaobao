package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
体检机构对接_体检状态查询 APIRequest
alibaba.alihealth.examination.reserve.state

体检机构对接_体检状态查询
*/
type AlibabaAlihealthExaminationReserveStateRequest struct {
    model.Params

    // 商户唯一码
    merchantCode   string 

    // 阿里健康预约唯一标识
    reserveNumber   string 

    // 体检机构预约唯一标识码
    uniqReserveCode   string 

}

func NewAlibabaAlihealthExaminationReserveStateRequest() *AlibabaAlihealthExaminationReserveStateRequest{
    return &AlibabaAlihealthExaminationReserveStateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthExaminationReserveStateRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.reserve.state"
}

func (r AlibabaAlihealthExaminationReserveStateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthExaminationReserveStateRequest) SetMerchantCode(merchantCode string) error {
    r.merchantCode = merchantCode
    r.Set("merchant_code", merchantCode)
    return nil
}

func (r AlibabaAlihealthExaminationReserveStateRequest) GetMerchantCode() string {
    return r.merchantCode
}

func (r *AlibabaAlihealthExaminationReserveStateRequest) SetReserveNumber(reserveNumber string) error {
    r.reserveNumber = reserveNumber
    r.Set("reserve_number", reserveNumber)
    return nil
}

func (r AlibabaAlihealthExaminationReserveStateRequest) GetReserveNumber() string {
    return r.reserveNumber
}

func (r *AlibabaAlihealthExaminationReserveStateRequest) SetUniqReserveCode(uniqReserveCode string) error {
    r.uniqReserveCode = uniqReserveCode
    r.Set("uniq_reserve_code", uniqReserveCode)
    return nil
}

func (r AlibabaAlihealthExaminationReserveStateRequest) GetUniqReserveCode() string {
    return r.uniqReserveCode
}

