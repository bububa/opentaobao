package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
计算补助金额 API请求
alibaba.alisports.efsp.countsubsidy

计算补助金额
*/
type AlibabaAlisportsEfspCountsubsidyRequest struct {
    model.Params
    // 订单金额
    _sumAmount   int64
    // 健身房ID
    _gymId   string
    // 企业ID
    _enterpriseId   string
    // alipayId
    _alipayId   string
    // 健身房所在省市
    _districtCode   string
}

// 初始化AlibabaAlisportsEfspCountsubsidyRequest对象
func NewAlibabaAlisportsEfspCountsubsidyRequest() *AlibabaAlisportsEfspCountsubsidyRequest{
    return &AlibabaAlisportsEfspCountsubsidyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsEfspCountsubsidyRequest) GetApiMethodName() string {
    return "alibaba.alisports.efsp.countsubsidy"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsEfspCountsubsidyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SumAmount Setter
// 订单金额
func (r *AlibabaAlisportsEfspCountsubsidyRequest) SetSumAmount(_sumAmount int64) error {
    r._sumAmount = _sumAmount
    r.Set("sum_amount", _sumAmount)
    return nil
}

// SumAmount Getter
func (r AlibabaAlisportsEfspCountsubsidyRequest) GetSumAmount() int64 {
    return r._sumAmount
}
// GymId Setter
// 健身房ID
func (r *AlibabaAlisportsEfspCountsubsidyRequest) SetGymId(_gymId string) error {
    r._gymId = _gymId
    r.Set("gym_id", _gymId)
    return nil
}

// GymId Getter
func (r AlibabaAlisportsEfspCountsubsidyRequest) GetGymId() string {
    return r._gymId
}
// EnterpriseId Setter
// 企业ID
func (r *AlibabaAlisportsEfspCountsubsidyRequest) SetEnterpriseId(_enterpriseId string) error {
    r._enterpriseId = _enterpriseId
    r.Set("enterprise_id", _enterpriseId)
    return nil
}

// EnterpriseId Getter
func (r AlibabaAlisportsEfspCountsubsidyRequest) GetEnterpriseId() string {
    return r._enterpriseId
}
// AlipayId Setter
// alipayId
func (r *AlibabaAlisportsEfspCountsubsidyRequest) SetAlipayId(_alipayId string) error {
    r._alipayId = _alipayId
    r.Set("alipay_id", _alipayId)
    return nil
}

// AlipayId Getter
func (r AlibabaAlisportsEfspCountsubsidyRequest) GetAlipayId() string {
    return r._alipayId
}
// DistrictCode Setter
// 健身房所在省市
func (r *AlibabaAlisportsEfspCountsubsidyRequest) SetDistrictCode(_districtCode string) error {
    r._districtCode = _districtCode
    r.Set("district_code", _districtCode)
    return nil
}

// DistrictCode Getter
func (r AlibabaAlisportsEfspCountsubsidyRequest) GetDistrictCode() string {
    return r._districtCode
}
