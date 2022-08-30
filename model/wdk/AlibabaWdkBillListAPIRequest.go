package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkBillListAPIRequest 五道口账单拉取接口 API请求
// alibaba.wdk.bill.list
//
// 五道口账单拉取接口
type AlibabaWdkBillListAPIRequest struct {
	model.Params
	// 入参
	_txdBillListGetRequest *TxdBillListGetRequest
}

// NewAlibabaWdkBillListRequest 初始化AlibabaWdkBillListAPIRequest对象
func NewAlibabaWdkBillListRequest() *AlibabaWdkBillListAPIRequest {
	return &AlibabaWdkBillListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkBillListAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.bill.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkBillListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTxdBillListGetRequest is TxdBillListGetRequest Setter
// 入参
func (r *AlibabaWdkBillListAPIRequest) SetTxdBillListGetRequest(_txdBillListGetRequest *TxdBillListGetRequest) error {
	r._txdBillListGetRequest = _txdBillListGetRequest
	r.Set("txd_bill_list_get_request", _txdBillListGetRequest)
	return nil
}

// GetTxdBillListGetRequest TxdBillListGetRequest Getter
func (r AlibabaWdkBillListAPIRequest) GetTxdBillListGetRequest() *TxdBillListGetRequest {
	return r._txdBillListGetRequest
}
