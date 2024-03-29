package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpUopSupplierReverseorderCreateAPIRequest 商家ERP发起创建销退单服务 API请求
// alibaba.ascp.uop.supplier.reverseorder.create
//
// 商家在收到消费者实物退货后，在ERP发起创建销退单服务
type AlibabaAscpUopSupplierReverseorderCreateAPIRequest struct {
	model.Params
	// 逆向销退单创建请求
	_reverseCreateRequest *ReverseCreateRequest
}

// NewAlibabaAscpUopSupplierReverseorderCreateRequest 初始化AlibabaAscpUopSupplierReverseorderCreateAPIRequest对象
func NewAlibabaAscpUopSupplierReverseorderCreateRequest() *AlibabaAscpUopSupplierReverseorderCreateAPIRequest {
	return &AlibabaAscpUopSupplierReverseorderCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpUopSupplierReverseorderCreateAPIRequest) Reset() {
	r._reverseCreateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopSupplierReverseorderCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.supplier.reverseorder.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopSupplierReverseorderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpUopSupplierReverseorderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReverseCreateRequest is ReverseCreateRequest Setter
// 逆向销退单创建请求
func (r *AlibabaAscpUopSupplierReverseorderCreateAPIRequest) SetReverseCreateRequest(_reverseCreateRequest *ReverseCreateRequest) error {
	r._reverseCreateRequest = _reverseCreateRequest
	r.Set("reverse_create_request", _reverseCreateRequest)
	return nil
}

// GetReverseCreateRequest ReverseCreateRequest Getter
func (r AlibabaAscpUopSupplierReverseorderCreateAPIRequest) GetReverseCreateRequest() *ReverseCreateRequest {
	return r._reverseCreateRequest
}

var poolAlibabaAscpUopSupplierReverseorderCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpUopSupplierReverseorderCreateRequest()
	},
}

// GetAlibabaAscpUopSupplierReverseorderCreateRequest 从 sync.Pool 获取 AlibabaAscpUopSupplierReverseorderCreateAPIRequest
func GetAlibabaAscpUopSupplierReverseorderCreateAPIRequest() *AlibabaAscpUopSupplierReverseorderCreateAPIRequest {
	return poolAlibabaAscpUopSupplierReverseorderCreateAPIRequest.Get().(*AlibabaAscpUopSupplierReverseorderCreateAPIRequest)
}

// ReleaseAlibabaAscpUopSupplierReverseorderCreateAPIRequest 将 AlibabaAscpUopSupplierReverseorderCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpUopSupplierReverseorderCreateAPIRequest(v *AlibabaAscpUopSupplierReverseorderCreateAPIRequest) {
	v.Reset()
	poolAlibabaAscpUopSupplierReverseorderCreateAPIRequest.Put(v)
}
