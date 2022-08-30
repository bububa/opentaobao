package ascpchannel

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopSelfSupplierWaybillQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.self.supplier.waybill.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopSelfSupplierWaybillQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
