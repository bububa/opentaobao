package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIRequest 经纪人接待状态变更 API请求
// alibaba.alihouse.existinghome.sync.broker.is.receive.customer
//
// 经纪人接待状态变更
type AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIRequest struct {
	model.Params
	// 1
	_outerId string
	// 1
	_outerStoreId string
	// 1
	_isReceiveCustomer int64
}

// NewAlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerRequest 初始化AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIRequest对象
func NewAlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerRequest() *AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIRequest {
	return &AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIRequest) Reset() {
	r._outerId = ""
	r._outerStoreId = ""
	r._isReceiveCustomer = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.sync.broker.is.receive.customer"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterId is OuterId Setter
// 1
func (r *AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetOuterStoreId is OuterStoreId Setter
// 1
func (r *AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIRequest) SetOuterStoreId(_outerStoreId string) error {
	r._outerStoreId = _outerStoreId
	r.Set("outer_store_id", _outerStoreId)
	return nil
}

// GetOuterStoreId OuterStoreId Getter
func (r AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIRequest) GetOuterStoreId() string {
	return r._outerStoreId
}

// SetIsReceiveCustomer is IsReceiveCustomer Setter
// 1
func (r *AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIRequest) SetIsReceiveCustomer(_isReceiveCustomer int64) error {
	r._isReceiveCustomer = _isReceiveCustomer
	r.Set("is_receive_customer", _isReceiveCustomer)
	return nil
}

// GetIsReceiveCustomer IsReceiveCustomer Getter
func (r AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIRequest) GetIsReceiveCustomer() int64 {
	return r._isReceiveCustomer
}

var poolAlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerRequest()
	},
}

// GetAlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerRequest 从 sync.Pool 获取 AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIRequest
func GetAlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIRequest() *AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIRequest {
	return poolAlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIRequest.Get().(*AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIRequest)
}

// ReleaseAlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIRequest 将 AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIRequest(v *AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIRequest.Put(v)
}
