package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpUopSupplierReverseorderInstorageResultAPIRequest 逆向销退入库单到仓结果回告 API请求
// alibaba.ascp.uop.supplier.reverseorder.instorage.result
//
// ERP回告销退入库单到仓信息回告
type AlibabaAscpUopSupplierReverseorderInstorageResultAPIRequest struct {
	model.Params
	// 消退入库单结果请求
	_instorageResultRequest *Instorageresultrequest
}

// NewAlibabaAscpUopSupplierReverseorderInstorageResultRequest 初始化AlibabaAscpUopSupplierReverseorderInstorageResultAPIRequest对象
func NewAlibabaAscpUopSupplierReverseorderInstorageResultRequest() *AlibabaAscpUopSupplierReverseorderInstorageResultAPIRequest {
	return &AlibabaAscpUopSupplierReverseorderInstorageResultAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpUopSupplierReverseorderInstorageResultAPIRequest) Reset() {
	r._instorageResultRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopSupplierReverseorderInstorageResultAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.supplier.reverseorder.instorage.result"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopSupplierReverseorderInstorageResultAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpUopSupplierReverseorderInstorageResultAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInstorageResultRequest is InstorageResultRequest Setter
// 消退入库单结果请求
func (r *AlibabaAscpUopSupplierReverseorderInstorageResultAPIRequest) SetInstorageResultRequest(_instorageResultRequest *Instorageresultrequest) error {
	r._instorageResultRequest = _instorageResultRequest
	r.Set("instorage_result_request", _instorageResultRequest)
	return nil
}

// GetInstorageResultRequest InstorageResultRequest Getter
func (r AlibabaAscpUopSupplierReverseorderInstorageResultAPIRequest) GetInstorageResultRequest() *Instorageresultrequest {
	return r._instorageResultRequest
}

var poolAlibabaAscpUopSupplierReverseorderInstorageResultAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpUopSupplierReverseorderInstorageResultRequest()
	},
}

// GetAlibabaAscpUopSupplierReverseorderInstorageResultRequest 从 sync.Pool 获取 AlibabaAscpUopSupplierReverseorderInstorageResultAPIRequest
func GetAlibabaAscpUopSupplierReverseorderInstorageResultAPIRequest() *AlibabaAscpUopSupplierReverseorderInstorageResultAPIRequest {
	return poolAlibabaAscpUopSupplierReverseorderInstorageResultAPIRequest.Get().(*AlibabaAscpUopSupplierReverseorderInstorageResultAPIRequest)
}

// ReleaseAlibabaAscpUopSupplierReverseorderInstorageResultAPIRequest 将 AlibabaAscpUopSupplierReverseorderInstorageResultAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpUopSupplierReverseorderInstorageResultAPIRequest(v *AlibabaAscpUopSupplierReverseorderInstorageResultAPIRequest) {
	v.Reset()
	poolAlibabaAscpUopSupplierReverseorderInstorageResultAPIRequest.Put(v)
}
