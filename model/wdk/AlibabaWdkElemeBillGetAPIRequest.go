package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkelemebillgetAPIRequest 饿了么日维度对账单查询 API请求
// alibaba.wdk.eleme.bill.get
//
// 查询饿了么日维度对账单信息
type AlibabawdkelemebillgetAPIRequest struct {
	model.Params
	// 对账单查询参数
	_eleBillRequest *EleBillRequest
}

// NewAlibabawdkelemebillgetRequest 初始化AlibabawdkelemebillgetAPIRequest对象
func NewAlibabawdkelemebillgetRequest() *AlibabawdkelemebillgetAPIRequest {
	return &AlibabawdkelemebillgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkelemebillgetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.eleme.bill.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkelemebillgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkelemebillgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEleBillRequest is EleBillRequest Setter
// 对账单查询参数
func (r *AlibabawdkelemebillgetAPIRequest) SetEleBillRequest(_eleBillRequest *EleBillRequest) error {
	r._eleBillRequest = _eleBillRequest
	r.Set("ele_bill_request", _eleBillRequest)
	return nil
}

// GetEleBillRequest EleBillRequest Getter
func (r AlibabawdkelemebillgetAPIRequest) GetEleBillRequest() *EleBillRequest {
	return r._eleBillRequest
}
