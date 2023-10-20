package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomesyncbrokerisreceivecustomerAPIRequest 经纪人接待状态变更 API请求
// alibaba.alihouse.existinghome.sync.broker.is.receive.customer
//
// 经纪人接待状态变更
type AlibabaalihouseexistinghomesyncbrokerisreceivecustomerAPIRequest struct {
	model.Params
	// 1
	_outerId string
	// 1
	_outerStoreId string
	// 1
	_isReceiveCustomer int64
}

// NewAlibabaalihouseexistinghomesyncbrokerisreceivecustomerRequest 初始化AlibabaalihouseexistinghomesyncbrokerisreceivecustomerAPIRequest对象
func NewAlibabaalihouseexistinghomesyncbrokerisreceivecustomerRequest() *AlibabaalihouseexistinghomesyncbrokerisreceivecustomerAPIRequest {
	return &AlibabaalihouseexistinghomesyncbrokerisreceivecustomerAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomesyncbrokerisreceivecustomerAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.sync.broker.is.receive.customer"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomesyncbrokerisreceivecustomerAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomesyncbrokerisreceivecustomerAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterId is OuterId Setter
// 1
func (r *AlibabaalihouseexistinghomesyncbrokerisreceivecustomerAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaalihouseexistinghomesyncbrokerisreceivecustomerAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetOuterStoreId is OuterStoreId Setter
// 1
func (r *AlibabaalihouseexistinghomesyncbrokerisreceivecustomerAPIRequest) SetOuterStoreId(_outerStoreId string) error {
	r._outerStoreId = _outerStoreId
	r.Set("outer_store_id", _outerStoreId)
	return nil
}

// GetOuterStoreId OuterStoreId Getter
func (r AlibabaalihouseexistinghomesyncbrokerisreceivecustomerAPIRequest) GetOuterStoreId() string {
	return r._outerStoreId
}

// SetIsReceiveCustomer is IsReceiveCustomer Setter
// 1
func (r *AlibabaalihouseexistinghomesyncbrokerisreceivecustomerAPIRequest) SetIsReceiveCustomer(_isReceiveCustomer int64) error {
	r._isReceiveCustomer = _isReceiveCustomer
	r.Set("is_receive_customer", _isReceiveCustomer)
	return nil
}

// GetIsReceiveCustomer IsReceiveCustomer Getter
func (r AlibabaalihouseexistinghomesyncbrokerisreceivecustomerAPIRequest) GetIsReceiveCustomer() int64 {
	return r._isReceiveCustomer
}
