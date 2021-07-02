package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardPushAPIRequest 推送服务工单信息 API请求
// tmall.servicecenter.workcard.push
//
// 服务商家推送工单信息到天猫。
type TmallServicecenterWorkcardPushAPIRequest struct {
	model.Params
	// 属性列表。使用半角分号隔开,字符串前后都需要有半角分号
	_attributes string
	// 描述
	_desc string
	// 淘宝交易订单号
	_bizOrderId int64
	// 服务预约安装时间
	_serviceReserveTime string
	// 服务预约安装地址。四级地址与街道地址用空格隔开
	_serviceReserveAddress string
	// 0=初始化, 3=授理， 10=拒绝 ，4=执行 ，5=成功，11=失败
	_status string
}

// NewTmallServicecenterWorkcardPushRequest 初始化TmallServicecenterWorkcardPushAPIRequest对象
func NewTmallServicecenterWorkcardPushRequest() *TmallServicecenterWorkcardPushAPIRequest {
	return &TmallServicecenterWorkcardPushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardPushAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardPushAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAttributes is Attributes Setter
// 属性列表。使用半角分号隔开,字符串前后都需要有半角分号
func (r *TmallServicecenterWorkcardPushAPIRequest) SetAttributes(_attributes string) error {
	r._attributes = _attributes
	r.Set("attributes", _attributes)
	return nil
}

// GetAttributes Attributes Getter
func (r TmallServicecenterWorkcardPushAPIRequest) GetAttributes() string {
	return r._attributes
}

// SetDesc is Desc Setter
// 描述
func (r *TmallServicecenterWorkcardPushAPIRequest) SetDesc(_desc string) error {
	r._desc = _desc
	r.Set("desc", _desc)
	return nil
}

// GetDesc Desc Getter
func (r TmallServicecenterWorkcardPushAPIRequest) GetDesc() string {
	return r._desc
}

// SetBizOrderId is BizOrderId Setter
// 淘宝交易订单号
func (r *TmallServicecenterWorkcardPushAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TmallServicecenterWorkcardPushAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

// SetServiceReserveTime is ServiceReserveTime Setter
// 服务预约安装时间
func (r *TmallServicecenterWorkcardPushAPIRequest) SetServiceReserveTime(_serviceReserveTime string) error {
	r._serviceReserveTime = _serviceReserveTime
	r.Set("service_reserve_time", _serviceReserveTime)
	return nil
}

// GetServiceReserveTime ServiceReserveTime Getter
func (r TmallServicecenterWorkcardPushAPIRequest) GetServiceReserveTime() string {
	return r._serviceReserveTime
}

// SetServiceReserveAddress is ServiceReserveAddress Setter
// 服务预约安装地址。四级地址与街道地址用空格隔开
func (r *TmallServicecenterWorkcardPushAPIRequest) SetServiceReserveAddress(_serviceReserveAddress string) error {
	r._serviceReserveAddress = _serviceReserveAddress
	r.Set("service_reserve_address", _serviceReserveAddress)
	return nil
}

// GetServiceReserveAddress ServiceReserveAddress Getter
func (r TmallServicecenterWorkcardPushAPIRequest) GetServiceReserveAddress() string {
	return r._serviceReserveAddress
}

// SetStatus is Status Setter
// 0=初始化, 3=授理， 10=拒绝 ，4=执行 ，5=成功，11=失败
func (r *TmallServicecenterWorkcardPushAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TmallServicecenterWorkcardPushAPIRequest) GetStatus() string {
	return r._status
}
