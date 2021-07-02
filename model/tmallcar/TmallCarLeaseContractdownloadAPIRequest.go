package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeaseContractdownloadAPIRequest 天猫开新车租后合同下载 API请求
// tmall.car.lease.contractdownload
//
// 天猫开新车租后合同下载
type TmallCarLeaseContractdownloadAPIRequest struct {
	model.Params
	// 天猫开新车订单id
	_orderId int64
	// 续租协议： 1， 全款购车协议： 2，分期购买协议：3， 分期购买车辆资产验收协议：4，分期购买车辆抵押：5， 分期购买融资租赁合同：6
	_type string
}

// NewTmallCarLeaseContractdownloadRequest 初始化TmallCarLeaseContractdownloadAPIRequest对象
func NewTmallCarLeaseContractdownloadRequest() *TmallCarLeaseContractdownloadAPIRequest {
	return &TmallCarLeaseContractdownloadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarLeaseContractdownloadAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.contractdownload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarLeaseContractdownloadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderId Setter
// 天猫开新车订单id
func (r *TmallCarLeaseContractdownloadAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r TmallCarLeaseContractdownloadAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// Set is Type Setter
// 续租协议： 1， 全款购车协议： 2，分期购买协议：3， 分期购买车辆资产验收协议：4，分期购买车辆抵押：5， 分期购买融资租赁合同：6
func (r *TmallCarLeaseContractdownloadAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// Get Type Getter
func (r TmallCarLeaseContractdownloadAPIRequest) GetType() string {
	return r._type
}
