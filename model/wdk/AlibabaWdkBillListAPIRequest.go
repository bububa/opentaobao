package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkBillListAPIRequest) Reset() {
	r._txdBillListGetRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkBillListAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.bill.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkBillListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkBillListAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaWdkBillListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkBillListRequest()
	},
}

// GetAlibabaWdkBillListRequest 从 sync.Pool 获取 AlibabaWdkBillListAPIRequest
func GetAlibabaWdkBillListAPIRequest() *AlibabaWdkBillListAPIRequest {
	return poolAlibabaWdkBillListAPIRequest.Get().(*AlibabaWdkBillListAPIRequest)
}

// ReleaseAlibabaWdkBillListAPIRequest 将 AlibabaWdkBillListAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkBillListAPIRequest(v *AlibabaWdkBillListAPIRequest) {
	v.Reset()
	poolAlibabaWdkBillListAPIRequest.Put(v)
}
