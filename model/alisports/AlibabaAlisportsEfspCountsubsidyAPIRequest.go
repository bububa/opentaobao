package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsEfspCountsubsidyAPIRequest 计算补助金额 API请求
// alibaba.alisports.efsp.countsubsidy
//
// 计算补助金额
type AlibabaAlisportsEfspCountsubsidyAPIRequest struct {
	model.Params
	// 订单金额
	_sumAmount int64
	// 健身房ID
	_gymId string
	// 企业ID
	_enterpriseId string
	// alipayId
	_alipayId string
	// 健身房所在省市
	_districtCode string
}

// NewAlibabaAlisportsEfspCountsubsidyRequest 初始化AlibabaAlisportsEfspCountsubsidyAPIRequest对象
func NewAlibabaAlisportsEfspCountsubsidyRequest() *AlibabaAlisportsEfspCountsubsidyAPIRequest {
	return &AlibabaAlisportsEfspCountsubsidyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsEfspCountsubsidyAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.efsp.countsubsidy"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsEfspCountsubsidyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SumAmount Setter
// 订单金额
func (r *AlibabaAlisportsEfspCountsubsidyAPIRequest) SetSumAmount(_sumAmount int64) error {
	r._sumAmount = _sumAmount
	r.Set("sum_amount", _sumAmount)
	return nil
}

// Get SumAmount Getter
func (r AlibabaAlisportsEfspCountsubsidyAPIRequest) GetSumAmount() int64 {
	return r._sumAmount
}

// Set is GymId Setter
// 健身房ID
func (r *AlibabaAlisportsEfspCountsubsidyAPIRequest) SetGymId(_gymId string) error {
	r._gymId = _gymId
	r.Set("gym_id", _gymId)
	return nil
}

// Get GymId Getter
func (r AlibabaAlisportsEfspCountsubsidyAPIRequest) GetGymId() string {
	return r._gymId
}

// Set is EnterpriseId Setter
// 企业ID
func (r *AlibabaAlisportsEfspCountsubsidyAPIRequest) SetEnterpriseId(_enterpriseId string) error {
	r._enterpriseId = _enterpriseId
	r.Set("enterprise_id", _enterpriseId)
	return nil
}

// Get EnterpriseId Getter
func (r AlibabaAlisportsEfspCountsubsidyAPIRequest) GetEnterpriseId() string {
	return r._enterpriseId
}

// Set is AlipayId Setter
// alipayId
func (r *AlibabaAlisportsEfspCountsubsidyAPIRequest) SetAlipayId(_alipayId string) error {
	r._alipayId = _alipayId
	r.Set("alipay_id", _alipayId)
	return nil
}

// Get AlipayId Getter
func (r AlibabaAlisportsEfspCountsubsidyAPIRequest) GetAlipayId() string {
	return r._alipayId
}

// Set is DistrictCode Setter
// 健身房所在省市
func (r *AlibabaAlisportsEfspCountsubsidyAPIRequest) SetDistrictCode(_districtCode string) error {
	r._districtCode = _districtCode
	r.Set("district_code", _districtCode)
	return nil
}

// Get DistrictCode Getter
func (r AlibabaAlisportsEfspCountsubsidyAPIRequest) GetDistrictCode() string {
	return r._districtCode
}
