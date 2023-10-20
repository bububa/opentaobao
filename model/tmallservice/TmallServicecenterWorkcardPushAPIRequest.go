package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkcardpushAPIRequest 推送服务工单信息 API请求
// tmall.servicecenter.workcard.push
//
// 服务商家推送工单信息到天猫。
type TmallservicecenterworkcardpushAPIRequest struct {
	model.Params
	// 属性列表。使用半角分号隔开,字符串前后都需要有半角分号
	_attributes string
	// 描述
	_desc string
	// 服务预约安装时间
	_serviceReserveTime string
	// 服务预约安装地址。四级地址与街道地址用空格隔开
	_serviceReserveAddress string
	// 0=初始化, 3=授理， 10=拒绝 ，4=执行 ，5=成功，11=失败
	_status string
	// 淘宝交易订单号
	_bizOrderId int64
}

// NewTmallservicecenterworkcardpushRequest 初始化TmallservicecenterworkcardpushAPIRequest对象
func NewTmallservicecenterworkcardpushRequest() *TmallservicecenterworkcardpushAPIRequest {
	return &TmallservicecenterworkcardpushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterworkcardpushAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterworkcardpushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterworkcardpushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAttributes is Attributes Setter
// 属性列表。使用半角分号隔开,字符串前后都需要有半角分号
func (r *TmallservicecenterworkcardpushAPIRequest) SetAttributes(_attributes string) error {
	r._attributes = _attributes
	r.Set("attributes", _attributes)
	return nil
}

// GetAttributes Attributes Getter
func (r TmallservicecenterworkcardpushAPIRequest) GetAttributes() string {
	return r._attributes
}

// SetDesc is Desc Setter
// 描述
func (r *TmallservicecenterworkcardpushAPIRequest) SetDesc(_desc string) error {
	r._desc = _desc
	r.Set("desc", _desc)
	return nil
}

// GetDesc Desc Getter
func (r TmallservicecenterworkcardpushAPIRequest) GetDesc() string {
	return r._desc
}

// SetServiceReserveTime is ServiceReserveTime Setter
// 服务预约安装时间
func (r *TmallservicecenterworkcardpushAPIRequest) SetServiceReserveTime(_serviceReserveTime string) error {
	r._serviceReserveTime = _serviceReserveTime
	r.Set("service_reserve_time", _serviceReserveTime)
	return nil
}

// GetServiceReserveTime ServiceReserveTime Getter
func (r TmallservicecenterworkcardpushAPIRequest) GetServiceReserveTime() string {
	return r._serviceReserveTime
}

// SetServiceReserveAddress is ServiceReserveAddress Setter
// 服务预约安装地址。四级地址与街道地址用空格隔开
func (r *TmallservicecenterworkcardpushAPIRequest) SetServiceReserveAddress(_serviceReserveAddress string) error {
	r._serviceReserveAddress = _serviceReserveAddress
	r.Set("service_reserve_address", _serviceReserveAddress)
	return nil
}

// GetServiceReserveAddress ServiceReserveAddress Getter
func (r TmallservicecenterworkcardpushAPIRequest) GetServiceReserveAddress() string {
	return r._serviceReserveAddress
}

// SetStatus is Status Setter
// 0=初始化, 3=授理， 10=拒绝 ，4=执行 ，5=成功，11=失败
func (r *TmallservicecenterworkcardpushAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TmallservicecenterworkcardpushAPIRequest) GetStatus() string {
	return r._status
}

// SetBizOrderId is BizOrderId Setter
// 淘宝交易订单号
func (r *TmallservicecenterworkcardpushAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TmallservicecenterworkcardpushAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}
