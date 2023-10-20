package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpuopselfsupplierwaybillqueryAPIRequest 商家仓自营配电子面单取号 API请求
// alibaba.ascp.uop.self.supplier.waybill.query
//
// ERP调用打印面单取号接口
type AlibabaascpuopselfsupplierwaybillqueryAPIRequest struct {
	model.Params
	// 查询面单请求参数
	_waybillQueryRequest *Waybillqueryrequest
}

// NewAlibabaascpuopselfsupplierwaybillqueryRequest 初始化AlibabaascpuopselfsupplierwaybillqueryAPIRequest对象
func NewAlibabaascpuopselfsupplierwaybillqueryRequest() *AlibabaascpuopselfsupplierwaybillqueryAPIRequest {
	return &AlibabaascpuopselfsupplierwaybillqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpuopselfsupplierwaybillqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.self.supplier.waybill.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpuopselfsupplierwaybillqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpuopselfsupplierwaybillqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWaybillQueryRequest is WaybillQueryRequest Setter
// 查询面单请求参数
func (r *AlibabaascpuopselfsupplierwaybillqueryAPIRequest) SetWaybillQueryRequest(_waybillQueryRequest *Waybillqueryrequest) error {
	r._waybillQueryRequest = _waybillQueryRequest
	r.Set("waybill_query_request", _waybillQueryRequest)
	return nil
}

// GetWaybillQueryRequest WaybillQueryRequest Getter
func (r AlibabaascpuopselfsupplierwaybillqueryAPIRequest) GetWaybillQueryRequest() *Waybillqueryrequest {
	return r._waybillQueryRequest
}
