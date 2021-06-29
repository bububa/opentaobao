package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
计算补助金额 APIRequest
alibaba.alisports.efsp.countsubsidy

计算补助金额
*/
type AlibabaAlisportsEfspCountsubsidyRequest struct {
    model.Params

    // 订单金额
    sumAmount   int64 

    // 健身房ID
    gymId   string 

    // 企业ID
    enterpriseId   string 

    // alipayId
    alipayId   string 

    // 健身房所在省市
    districtCode   string 

}

func NewAlibabaAlisportsEfspCountsubsidyRequest() *AlibabaAlisportsEfspCountsubsidyRequest{
    return &AlibabaAlisportsEfspCountsubsidyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlisportsEfspCountsubsidyRequest) GetApiMethodName() string {
    return "alibaba.alisports.efsp.countsubsidy"
}

func (r AlibabaAlisportsEfspCountsubsidyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlisportsEfspCountsubsidyRequest) SetSumAmount(sumAmount int64) error {
    r.sumAmount = sumAmount
    r.Set("sum_amount", sumAmount)
    return nil
}

func (r AlibabaAlisportsEfspCountsubsidyRequest) GetSumAmount() int64 {
    return r.sumAmount
}

func (r *AlibabaAlisportsEfspCountsubsidyRequest) SetGymId(gymId string) error {
    r.gymId = gymId
    r.Set("gym_id", gymId)
    return nil
}

func (r AlibabaAlisportsEfspCountsubsidyRequest) GetGymId() string {
    return r.gymId
}

func (r *AlibabaAlisportsEfspCountsubsidyRequest) SetEnterpriseId(enterpriseId string) error {
    r.enterpriseId = enterpriseId
    r.Set("enterprise_id", enterpriseId)
    return nil
}

func (r AlibabaAlisportsEfspCountsubsidyRequest) GetEnterpriseId() string {
    return r.enterpriseId
}

func (r *AlibabaAlisportsEfspCountsubsidyRequest) SetAlipayId(alipayId string) error {
    r.alipayId = alipayId
    r.Set("alipay_id", alipayId)
    return nil
}

func (r AlibabaAlisportsEfspCountsubsidyRequest) GetAlipayId() string {
    return r.alipayId
}

func (r *AlibabaAlisportsEfspCountsubsidyRequest) SetDistrictCode(districtCode string) error {
    r.districtCode = districtCode
    r.Set("district_code", districtCode)
    return nil
}

func (r AlibabaAlisportsEfspCountsubsidyRequest) GetDistrictCode() string {
    return r.districtCode
}

