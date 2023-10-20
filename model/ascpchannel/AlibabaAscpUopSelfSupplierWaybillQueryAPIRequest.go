package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpUopSelfSupplierWaybillQueryAPIRequest 商家仓自营配电子面单取号 API请求
// alibaba.ascp.uop.self.supplier.waybill.query
//
// ERP调用打印面单取号接口
type AlibabaAscpUopSelfSupplierWaybillQueryAPIRequest struct {
	model.Params
	// 查询面单请求参数
	_waybillQueryRequest *Waybillqueryrequest
}

// NewAlibabaAscpUopSelfSupplierWaybillQueryRequest 初始化AlibabaAscpUopSelfSupplierWaybillQueryAPIRequest对象
func NewAlibabaAscpUopSelfSupplierWaybillQueryRequest() *AlibabaAscpUopSelfSupplierWaybillQueryAPIRequest {
	return &AlibabaAscpUopSelfSupplierWaybillQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpUopSelfSupplierWaybillQueryAPIRequest) Reset() {
	r._waybillQueryRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopSelfSupplierWaybillQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.self.supplier.waybill.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopSelfSupplierWaybillQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpUopSelfSupplierWaybillQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWaybillQueryRequest is WaybillQueryRequest Setter
// 查询面单请求参数
func (r *AlibabaAscpUopSelfSupplierWaybillQueryAPIRequest) SetWaybillQueryRequest(_waybillQueryRequest *Waybillqueryrequest) error {
	r._waybillQueryRequest = _waybillQueryRequest
	r.Set("waybill_query_request", _waybillQueryRequest)
	return nil
}

// GetWaybillQueryRequest WaybillQueryRequest Getter
func (r AlibabaAscpUopSelfSupplierWaybillQueryAPIRequest) GetWaybillQueryRequest() *Waybillqueryrequest {
	return r._waybillQueryRequest
}

var poolAlibabaAscpUopSelfSupplierWaybillQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpUopSelfSupplierWaybillQueryRequest()
	},
}

// GetAlibabaAscpUopSelfSupplierWaybillQueryRequest 从 sync.Pool 获取 AlibabaAscpUopSelfSupplierWaybillQueryAPIRequest
func GetAlibabaAscpUopSelfSupplierWaybillQueryAPIRequest() *AlibabaAscpUopSelfSupplierWaybillQueryAPIRequest {
	return poolAlibabaAscpUopSelfSupplierWaybillQueryAPIRequest.Get().(*AlibabaAscpUopSelfSupplierWaybillQueryAPIRequest)
}

// ReleaseAlibabaAscpUopSelfSupplierWaybillQueryAPIRequest 将 AlibabaAscpUopSelfSupplierWaybillQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpUopSelfSupplierWaybillQueryAPIRequest(v *AlibabaAscpUopSelfSupplierWaybillQueryAPIRequest) {
	v.Reset()
	poolAlibabaAscpUopSelfSupplierWaybillQueryAPIRequest.Put(v)
}
