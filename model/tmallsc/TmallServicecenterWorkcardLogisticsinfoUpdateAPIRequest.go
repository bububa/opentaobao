package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest 物流单信息回传接口 API请求
// tmall.servicecenter.workcard.logisticsinfo.update
//
// 物流单信息回传接口
type TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest struct {
	model.Params
	// create, pickup_finished 揽收完成， signed 签收
	_statusCode string
	// 外部单据id
	_outerId string
	// 物流单号
	_expressCode string
	// 物流公司
	_expressCompany string
	// 快递员手机号
	_courierMobile string
	// 快递员名称
	_courierName string
	// 扩展信息
	_extendInfo string
	// 揽收完成时间
	_pickupFinishTime string
	// 揽收上门确认时间
	_pickupDoorTime string
	// 签收时间
	_signTime string
	// 物流单id
	_logisticsOrderId int64
}

// NewTmallServicecenterWorkcardLogisticsinfoUpdateRequest 初始化TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest对象
func NewTmallServicecenterWorkcardLogisticsinfoUpdateRequest() *TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest {
	return &TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.logisticsinfo.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStatusCode is StatusCode Setter
// create, pickup_finished 揽收完成， signed 签收
func (r *TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest) SetStatusCode(_statusCode string) error {
	r._statusCode = _statusCode
	r.Set("status_code", _statusCode)
	return nil
}

// GetStatusCode StatusCode Getter
func (r TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest) GetStatusCode() string {
	return r._statusCode
}

// SetOuterId is OuterId Setter
// 外部单据id
func (r *TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetExpressCode is ExpressCode Setter
// 物流单号
func (r *TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest) SetExpressCode(_expressCode string) error {
	r._expressCode = _expressCode
	r.Set("express_code", _expressCode)
	return nil
}

// GetExpressCode ExpressCode Getter
func (r TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest) GetExpressCode() string {
	return r._expressCode
}

// SetExpressCompany is ExpressCompany Setter
// 物流公司
func (r *TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest) SetExpressCompany(_expressCompany string) error {
	r._expressCompany = _expressCompany
	r.Set("express_company", _expressCompany)
	return nil
}

// GetExpressCompany ExpressCompany Getter
func (r TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest) GetExpressCompany() string {
	return r._expressCompany
}

// SetCourierMobile is CourierMobile Setter
// 快递员手机号
func (r *TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest) SetCourierMobile(_courierMobile string) error {
	r._courierMobile = _courierMobile
	r.Set("courier_mobile", _courierMobile)
	return nil
}

// GetCourierMobile CourierMobile Getter
func (r TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest) GetCourierMobile() string {
	return r._courierMobile
}

// SetCourierName is CourierName Setter
// 快递员名称
func (r *TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest) SetCourierName(_courierName string) error {
	r._courierName = _courierName
	r.Set("courier_name", _courierName)
	return nil
}

// GetCourierName CourierName Getter
func (r TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest) GetCourierName() string {
	return r._courierName
}

// SetExtendInfo is ExtendInfo Setter
// 扩展信息
func (r *TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest) SetExtendInfo(_extendInfo string) error {
	r._extendInfo = _extendInfo
	r.Set("extend_info", _extendInfo)
	return nil
}

// GetExtendInfo ExtendInfo Getter
func (r TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest) GetExtendInfo() string {
	return r._extendInfo
}

// SetPickupFinishTime is PickupFinishTime Setter
// 揽收完成时间
func (r *TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest) SetPickupFinishTime(_pickupFinishTime string) error {
	r._pickupFinishTime = _pickupFinishTime
	r.Set("pickup_finish_time", _pickupFinishTime)
	return nil
}

// GetPickupFinishTime PickupFinishTime Getter
func (r TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest) GetPickupFinishTime() string {
	return r._pickupFinishTime
}

// SetPickupDoorTime is PickupDoorTime Setter
// 揽收上门确认时间
func (r *TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest) SetPickupDoorTime(_pickupDoorTime string) error {
	r._pickupDoorTime = _pickupDoorTime
	r.Set("pickup_door_time", _pickupDoorTime)
	return nil
}

// GetPickupDoorTime PickupDoorTime Getter
func (r TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest) GetPickupDoorTime() string {
	return r._pickupDoorTime
}

// SetSignTime is SignTime Setter
// 签收时间
func (r *TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest) SetSignTime(_signTime string) error {
	r._signTime = _signTime
	r.Set("sign_time", _signTime)
	return nil
}

// GetSignTime SignTime Getter
func (r TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest) GetSignTime() string {
	return r._signTime
}

// SetLogisticsOrderId is LogisticsOrderId Setter
// 物流单id
func (r *TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest) SetLogisticsOrderId(_logisticsOrderId int64) error {
	r._logisticsOrderId = _logisticsOrderId
	r.Set("logistics_order_id", _logisticsOrderId)
	return nil
}

// GetLogisticsOrderId LogisticsOrderId Getter
func (r TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest) GetLogisticsOrderId() int64 {
	return r._logisticsOrderId
}
