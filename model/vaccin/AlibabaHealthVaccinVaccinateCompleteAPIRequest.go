package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahealthvaccinvaccinatecompleteAPIRequest 接种完成反馈接口 API请求
// alibaba.health.vaccin.vaccinate.complete
//
// ISV 将用户完成接种的疫苗同步给免疫规划中心
type AlibabahealthvaccinvaccinatecompleteAPIRequest struct {
	model.Params
	// 接种的疫苗信息
	_vaccineList []VaccineInfo
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
}

// NewAlibabahealthvaccinvaccinatecompleteRequest 初始化AlibabahealthvaccinvaccinatecompleteAPIRequest对象
func NewAlibabahealthvaccinvaccinatecompleteRequest() *AlibabahealthvaccinvaccinatecompleteAPIRequest {
	return &AlibabahealthvaccinvaccinatecompleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahealthvaccinvaccinatecompleteAPIRequest) GetApiMethodName() string {
	return "alibaba.health.vaccin.vaccinate.complete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahealthvaccinvaccinatecompleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahealthvaccinvaccinatecompleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVaccineList is VaccineList Setter
// 接种的疫苗信息
func (r *AlibabahealthvaccinvaccinatecompleteAPIRequest) SetVaccineList(_vaccineList []VaccineInfo) error {
	r._vaccineList = _vaccineList
	r.Set("vaccine_list", _vaccineList)
	return nil
}

// GetVaccineList VaccineList Getter
func (r AlibabahealthvaccinvaccinatecompleteAPIRequest) GetVaccineList() []VaccineInfo {
	return r._vaccineList
}

// SetAlipayUserId is AlipayUserId Setter
// 支付宝用户 ID
func (r *AlibabahealthvaccinvaccinatecompleteAPIRequest) SetAlipayUserId(_alipayUserId string) error {
	r._alipayUserId = _alipayUserId
	r.Set("alipay_user_id", _alipayUserId)
	return nil
}

// GetAlipayUserId AlipayUserId Getter
func (r AlibabahealthvaccinvaccinatecompleteAPIRequest) GetAlipayUserId() string {
	return r._alipayUserId
}

// SetIsvUserId is IsvUserId Setter
// ISV 侧用户 ID
func (r *AlibabahealthvaccinvaccinatecompleteAPIRequest) SetIsvUserId(_isvUserId string) error {
	r._isvUserId = _isvUserId
	r.Set("isv_user_id", _isvUserId)
	return nil
}

// GetIsvUserId IsvUserId Getter
func (r AlibabahealthvaccinvaccinatecompleteAPIRequest) GetIsvUserId() string {
	return r._isvUserId
}

// SetOrderId is OrderId Setter
// 订单 ID
func (r *AlibabahealthvaccinvaccinatecompleteAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabahealthvaccinvaccinatecompleteAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetName is Name Setter
// 接种人姓名
func (r *AlibabahealthvaccinvaccinatecompleteAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r AlibabahealthvaccinvaccinatecompleteAPIRequest) GetName() string {
	return r._name
}

// SetMobile is Mobile Setter
// 联系电话
func (r *AlibabahealthvaccinvaccinatecompleteAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r AlibabahealthvaccinvaccinatecompleteAPIRequest) GetMobile() string {
	return r._mobile
}

// SetVaccinateDate is VaccinateDate Setter
// 接种日期
func (r *AlibabahealthvaccinvaccinatecompleteAPIRequest) SetVaccinateDate(_vaccinateDate string) error {
	r._vaccinateDate = _vaccinateDate
	r.Set("vaccinate_date", _vaccinateDate)
	return nil
}

// GetVaccinateDate VaccinateDate Getter
func (r AlibabahealthvaccinvaccinatecompleteAPIRequest) GetVaccinateDate() string {
	return r._vaccinateDate
}

// SetVaccinateTime is VaccinateTime Setter
// 接种时间
func (r *AlibabahealthvaccinvaccinatecompleteAPIRequest) SetVaccinateTime(_vaccinateTime string) error {
	r._vaccinateTime = _vaccinateTime
	r.Set("vaccinate_time", _vaccinateTime)
	return nil
}

// GetVaccinateTime VaccinateTime Getter
func (r AlibabahealthvaccinvaccinatecompleteAPIRequest) GetVaccinateTime() string {
	return r._vaccinateTime
}
