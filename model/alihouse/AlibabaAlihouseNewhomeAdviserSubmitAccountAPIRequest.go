package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeadvisersubmitaccountAPIRequest 顾问入职离职 API请求
// alibaba.alihouse.newhome.adviser.submit.account
//
// 顾问入职离职
type AlibabaalihousenewhomeadvisersubmitaccountAPIRequest struct {
	model.Params
	// 外部顾问ID
	_outerConsultantId string
	// 外部门店id
	_outerShopId string
	// 1-入职 2-离职
	_status int64
	// 版本号，时间戳
	_version int64
}

// NewAlibabaalihousenewhomeadvisersubmitaccountRequest 初始化AlibabaalihousenewhomeadvisersubmitaccountAPIRequest对象
func NewAlibabaalihousenewhomeadvisersubmitaccountRequest() *AlibabaalihousenewhomeadvisersubmitaccountAPIRequest {
	return &AlibabaalihousenewhomeadvisersubmitaccountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeadvisersubmitaccountAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.adviser.submit.account"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeadvisersubmitaccountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeadvisersubmitaccountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterConsultantId is OuterConsultantId Setter
// 外部顾问ID
func (r *AlibabaalihousenewhomeadvisersubmitaccountAPIRequest) SetOuterConsultantId(_outerConsultantId string) error {
	r._outerConsultantId = _outerConsultantId
	r.Set("outer_consultant_id", _outerConsultantId)
	return nil
}

// GetOuterConsultantId OuterConsultantId Getter
func (r AlibabaalihousenewhomeadvisersubmitaccountAPIRequest) GetOuterConsultantId() string {
	return r._outerConsultantId
}

// SetOuterShopId is OuterShopId Setter
// 外部门店id
func (r *AlibabaalihousenewhomeadvisersubmitaccountAPIRequest) SetOuterShopId(_outerShopId string) error {
	r._outerShopId = _outerShopId
	r.Set("outer_shop_id", _outerShopId)
	return nil
}

// GetOuterShopId OuterShopId Getter
func (r AlibabaalihousenewhomeadvisersubmitaccountAPIRequest) GetOuterShopId() string {
	return r._outerShopId
}

// SetStatus is Status Setter
// 1-入职 2-离职
func (r *AlibabaalihousenewhomeadvisersubmitaccountAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaalihousenewhomeadvisersubmitaccountAPIRequest) GetStatus() int64 {
	return r._status
}

// SetVersion is Version Setter
// 版本号，时间戳
func (r *AlibabaalihousenewhomeadvisersubmitaccountAPIRequest) SetVersion(_version int64) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r AlibabaalihousenewhomeadvisersubmitaccountAPIRequest) GetVersion() int64 {
	return r._version
}
