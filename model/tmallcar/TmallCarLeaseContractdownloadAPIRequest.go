package tmallcar

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeaseContractdownloadAPIRequest 天猫开新车租后合同下载 API请求
// tmall.car.lease.contractdownload
//
// 天猫开新车租后合同下载
type TmallCarLeaseContractdownloadAPIRequest struct {
	model.Params
	// 续租协议： 1， 全款购车协议： 2，分期购买协议：3， 分期购买车辆资产验收协议：4，分期购买车辆抵押：5， 分期购买融资租赁合同：6
	_type string
	// 天猫开新车订单id
	_orderId int64
}

// NewTmallCarLeaseContractdownloadRequest 初始化TmallCarLeaseContractdownloadAPIRequest对象
func NewTmallCarLeaseContractdownloadRequest() *TmallCarLeaseContractdownloadAPIRequest {
	return &TmallCarLeaseContractdownloadAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallCarLeaseContractdownloadAPIRequest) Reset() {
	r._type = ""
	r._orderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarLeaseContractdownloadAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.contractdownload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarLeaseContractdownloadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallCarLeaseContractdownloadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetType is Type Setter
// 续租协议： 1， 全款购车协议： 2，分期购买协议：3， 分期购买车辆资产验收协议：4，分期购买车辆抵押：5， 分期购买融资租赁合同：6
func (r *TmallCarLeaseContractdownloadAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TmallCarLeaseContractdownloadAPIRequest) GetType() string {
	return r._type
}

// SetOrderId is OrderId Setter
// 天猫开新车订单id
func (r *TmallCarLeaseContractdownloadAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TmallCarLeaseContractdownloadAPIRequest) GetOrderId() int64 {
	return r._orderId
}

var poolTmallCarLeaseContractdownloadAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallCarLeaseContractdownloadRequest()
	},
}

// GetTmallCarLeaseContractdownloadRequest 从 sync.Pool 获取 TmallCarLeaseContractdownloadAPIRequest
func GetTmallCarLeaseContractdownloadAPIRequest() *TmallCarLeaseContractdownloadAPIRequest {
	return poolTmallCarLeaseContractdownloadAPIRequest.Get().(*TmallCarLeaseContractdownloadAPIRequest)
}

// ReleaseTmallCarLeaseContractdownloadAPIRequest 将 TmallCarLeaseContractdownloadAPIRequest 放入 sync.Pool
func ReleaseTmallCarLeaseContractdownloadAPIRequest(v *TmallCarLeaseContractdownloadAPIRequest) {
	v.Reset()
	poolTmallCarLeaseContractdownloadAPIRequest.Put(v)
}
