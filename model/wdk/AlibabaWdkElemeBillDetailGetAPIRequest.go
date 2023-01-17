package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkElemeBillDetailGetAPIRequest 饿了么对账单查询，带订单明细 API请求
// alibaba.wdk.eleme.bill.detail.get
//
// 查询饿了么对账单信息，带订单明细
type AlibabaWdkElemeBillDetailGetAPIRequest struct {
	model.Params
	// 对账单查询参数
	_eleBillRequest *EleBillRequest
}

// NewAlibabaWdkElemeBillDetailGetRequest 初始化AlibabaWdkElemeBillDetailGetAPIRequest对象
func NewAlibabaWdkElemeBillDetailGetRequest() *AlibabaWdkElemeBillDetailGetAPIRequest {
	return &AlibabaWdkElemeBillDetailGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkElemeBillDetailGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.eleme.bill.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkElemeBillDetailGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkElemeBillDetailGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEleBillRequest is EleBillRequest Setter
// 对账单查询参数
func (r *AlibabaWdkElemeBillDetailGetAPIRequest) SetEleBillRequest(_eleBillRequest *EleBillRequest) error {
	r._eleBillRequest = _eleBillRequest
	r.Set("ele_bill_request", _eleBillRequest)
	return nil
}

// GetEleBillRequest EleBillRequest Getter
func (r AlibabaWdkElemeBillDetailGetAPIRequest) GetEleBillRequest() *EleBillRequest {
	return r._eleBillRequest
}
