package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHealthVaccinVaccinateCompleteAPIRequest
接种完成反馈接口 API请求
alibaba.health.vaccin.vaccinate.complete

ISV 将用户完成接种的疫苗同步给免疫规划中心 */
type AlibabaHealthVaccinVaccinateCompleteAPIRequest struct {
	model.Params
	// 支付宝用户 ID
	_alipayUserId string
	// ISV 侧用户 ID
	_isvUserId string
	// 订单 ID
	_orderId string
	// 接种人姓名
	_name string
	// 联系电话
	_mobile string
	// 接种日期
	_vaccinateDate string
	// 接种时间
	_vaccinateTime string
	// 接种的疫苗信息
	_vaccineList []VaccineInfo
}

// NewAlibabaHealthVaccinVaccinateCompleteRequest 初始化AlibabaHealthVaccinVaccinateCompleteAPIRequest对象
func NewAlibabaHealthVaccinVaccinateCompleteRequest() *AlibabaHealthVaccinVaccinateCompleteAPIRequest {
	return &AlibabaHealthVaccinVaccinateCompleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinVaccinateCompleteAPIRequest) GetApiMethodName() string {
	return "alibaba.health.vaccin.vaccinate.complete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinVaccinateCompleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AlipayUserId Setter
// 支付宝用户 ID
func (r *AlibabaHealthVaccinVaccinateCompleteAPIRequest) SetAlipayUserId(_alipayUserId string) error {
	r._alipayUserId = _alipayUserId
	r.Set("alipay_user_id", _alipayUserId)
	return nil
}

// Get AlipayUserId Getter
func (r AlibabaHealthVaccinVaccinateCompleteAPIRequest) GetAlipayUserId() string {
	return r._alipayUserId
}

// Set is IsvUserId Setter
// ISV 侧用户 ID
func (r *AlibabaHealthVaccinVaccinateCompleteAPIRequest) SetIsvUserId(_isvUserId string) error {
	r._isvUserId = _isvUserId
	r.Set("isv_user_id", _isvUserId)
	return nil
}

// Get IsvUserId Getter
func (r AlibabaHealthVaccinVaccinateCompleteAPIRequest) GetIsvUserId() string {
	return r._isvUserId
}

// Set is OrderId Setter
// 订单 ID
func (r *AlibabaHealthVaccinVaccinateCompleteAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r AlibabaHealthVaccinVaccinateCompleteAPIRequest) GetOrderId() string {
	return r._orderId
}

// Set is Name Setter
// 接种人姓名
func (r *AlibabaHealthVaccinVaccinateCompleteAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// Get Name Getter
func (r AlibabaHealthVaccinVaccinateCompleteAPIRequest) GetName() string {
	return r._name
}

// Set is Mobile Setter
// 联系电话
func (r *AlibabaHealthVaccinVaccinateCompleteAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// Get Mobile Getter
func (r AlibabaHealthVaccinVaccinateCompleteAPIRequest) GetMobile() string {
	return r._mobile
}

// Set is VaccinateDate Setter
// 接种日期
func (r *AlibabaHealthVaccinVaccinateCompleteAPIRequest) SetVaccinateDate(_vaccinateDate string) error {
	r._vaccinateDate = _vaccinateDate
	r.Set("vaccinate_date", _vaccinateDate)
	return nil
}

// Get VaccinateDate Getter
func (r AlibabaHealthVaccinVaccinateCompleteAPIRequest) GetVaccinateDate() string {
	return r._vaccinateDate
}

// Set is VaccinateTime Setter
// 接种时间
func (r *AlibabaHealthVaccinVaccinateCompleteAPIRequest) SetVaccinateTime(_vaccinateTime string) error {
	r._vaccinateTime = _vaccinateTime
	r.Set("vaccinate_time", _vaccinateTime)
	return nil
}

// Get VaccinateTime Getter
func (r AlibabaHealthVaccinVaccinateCompleteAPIRequest) GetVaccinateTime() string {
	return r._vaccinateTime
}

// Set is VaccineList Setter
// 接种的疫苗信息
func (r *AlibabaHealthVaccinVaccinateCompleteAPIRequest) SetVaccineList(_vaccineList []VaccineInfo) error {
	r._vaccineList = _vaccineList
	r.Set("vaccine_list", _vaccineList)
	return nil
}

// Get VaccineList Getter
func (r AlibabaHealthVaccinVaccinateCompleteAPIRequest) GetVaccineList() []VaccineInfo {
	return r._vaccineList
}
