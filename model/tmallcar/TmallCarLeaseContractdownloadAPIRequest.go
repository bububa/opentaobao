package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcarleasecontractdownloadAPIRequest 天猫开新车租后合同下载 API请求
// tmall.car.lease.contractdownload
//
// 天猫开新车租后合同下载
type TmallcarleasecontractdownloadAPIRequest struct {
	model.Params
	// 续租协议： 1， 全款购车协议： 2，分期购买协议：3， 分期购买车辆资产验收协议：4，分期购买车辆抵押：5， 分期购买融资租赁合同：6
	_type string
	// 天猫开新车订单id
	_orderId int64
}

// NewTmallcarleasecontractdownloadRequest 初始化TmallcarleasecontractdownloadAPIRequest对象
func NewTmallcarleasecontractdownloadRequest() *TmallcarleasecontractdownloadAPIRequest {
	return &TmallcarleasecontractdownloadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcarleasecontractdownloadAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.contractdownload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcarleasecontractdownloadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcarleasecontractdownloadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetType is Type Setter
// 续租协议： 1， 全款购车协议： 2，分期购买协议：3， 分期购买车辆资产验收协议：4，分期购买车辆抵押：5， 分期购买融资租赁合同：6
func (r *TmallcarleasecontractdownloadAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TmallcarleasecontractdownloadAPIRequest) GetType() string {
	return r._type
}

// SetOrderId is OrderId Setter
// 天猫开新车订单id
func (r *TmallcarleasecontractdownloadAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TmallcarleasecontractdownloadAPIRequest) GetOrderId() int64 {
	return r._orderId
}
