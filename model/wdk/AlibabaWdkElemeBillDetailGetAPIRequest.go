package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkelemebilldetailgetAPIRequest 饿了么对账单查询，带订单明细 API请求
// alibaba.wdk.eleme.bill.detail.get
//
// 查询饿了么对账单信息，带订单明细
type AlibabawdkelemebilldetailgetAPIRequest struct {
	model.Params
	// 对账单查询参数
	_eleBillRequest *EleBillRequest
}

// NewAlibabawdkelemebilldetailgetRequest 初始化AlibabawdkelemebilldetailgetAPIRequest对象
func NewAlibabawdkelemebilldetailgetRequest() *AlibabawdkelemebilldetailgetAPIRequest {
	return &AlibabawdkelemebilldetailgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkelemebilldetailgetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.eleme.bill.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkelemebilldetailgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkelemebilldetailgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEleBillRequest is EleBillRequest Setter
// 对账单查询参数
func (r *AlibabawdkelemebilldetailgetAPIRequest) SetEleBillRequest(_eleBillRequest *EleBillRequest) error {
	r._eleBillRequest = _eleBillRequest
	r.Set("ele_bill_request", _eleBillRequest)
	return nil
}

// GetEleBillRequest EleBillRequest Getter
func (r AlibabawdkelemebilldetailgetAPIRequest) GetEleBillRequest() *EleBillRequest {
	return r._eleBillRequest
}
