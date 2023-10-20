package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkbilllistAPIRequest 五道口账单拉取接口 API请求
// alibaba.wdk.bill.list
//
// 五道口账单拉取接口
type AlibabawdkbilllistAPIRequest struct {
	model.Params
	// 入参
	_txdBillListGetRequest *TxdBillListGetRequest
}

// NewAlibabawdkbilllistRequest 初始化AlibabawdkbilllistAPIRequest对象
func NewAlibabawdkbilllistRequest() *AlibabawdkbilllistAPIRequest {
	return &AlibabawdkbilllistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkbilllistAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.bill.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkbilllistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkbilllistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTxdBillListGetRequest is TxdBillListGetRequest Setter
// 入参
func (r *AlibabawdkbilllistAPIRequest) SetTxdBillListGetRequest(_txdBillListGetRequest *TxdBillListGetRequest) error {
	r._txdBillListGetRequest = _txdBillListGetRequest
	r.Set("txd_bill_list_get_request", _txdBillListGetRequest)
	return nil
}

// GetTxdBillListGetRequest TxdBillListGetRequest Getter
func (r AlibabawdkbilllistAPIRequest) GetTxdBillListGetRequest() *TxdBillListGetRequest {
	return r._txdBillListGetRequest
}
