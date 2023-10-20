package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkElemeBillGetAPIRequest 饿了么日维度对账单查询 API请求
// alibaba.wdk.eleme.bill.get
//
// 查询饿了么日维度对账单信息
type AlibabaWdkElemeBillGetAPIRequest struct {
	model.Params
	// 对账单查询参数
	_eleBillRequest *EleBillRequest
}

// NewAlibabaWdkElemeBillGetRequest 初始化AlibabaWdkElemeBillGetAPIRequest对象
func NewAlibabaWdkElemeBillGetRequest() *AlibabaWdkElemeBillGetAPIRequest {
	return &AlibabaWdkElemeBillGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkElemeBillGetAPIRequest) Reset() {
	r._eleBillRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkElemeBillGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.eleme.bill.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkElemeBillGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkElemeBillGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEleBillRequest is EleBillRequest Setter
// 对账单查询参数
func (r *AlibabaWdkElemeBillGetAPIRequest) SetEleBillRequest(_eleBillRequest *EleBillRequest) error {
	r._eleBillRequest = _eleBillRequest
	r.Set("ele_bill_request", _eleBillRequest)
	return nil
}

// GetEleBillRequest EleBillRequest Getter
func (r AlibabaWdkElemeBillGetAPIRequest) GetEleBillRequest() *EleBillRequest {
	return r._eleBillRequest
}

var poolAlibabaWdkElemeBillGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkElemeBillGetRequest()
	},
}

// GetAlibabaWdkElemeBillGetRequest 从 sync.Pool 获取 AlibabaWdkElemeBillGetAPIRequest
func GetAlibabaWdkElemeBillGetAPIRequest() *AlibabaWdkElemeBillGetAPIRequest {
	return poolAlibabaWdkElemeBillGetAPIRequest.Get().(*AlibabaWdkElemeBillGetAPIRequest)
}

// ReleaseAlibabaWdkElemeBillGetAPIRequest 将 AlibabaWdkElemeBillGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkElemeBillGetAPIRequest(v *AlibabaWdkElemeBillGetAPIRequest) {
	v.Reset()
	poolAlibabaWdkElemeBillGetAPIRequest.Put(v)
}
