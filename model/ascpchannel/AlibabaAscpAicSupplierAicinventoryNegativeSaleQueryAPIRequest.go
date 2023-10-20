package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest 商家查询负卖库存 API请求
// alibaba.ascp.aic.supplier.aicinventory.negative.sale.query
//
// 商家根据当前接口查询负卖货品的库存
type AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest struct {
	model.Params
	// 库存查询参数
	_aicinventoryQueryRequest *Aicinventoryqueryrequest
}

// NewAlibabaAscpAicSupplierAicinventoryNegativeSaleQueryRequest 初始化AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest对象
func NewAlibabaAscpAicSupplierAicinventoryNegativeSaleQueryRequest() *AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest {
	return &AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest) Reset() {
	r._aicinventoryQueryRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.aic.supplier.aicinventory.negative.sale.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAicinventoryQueryRequest is AicinventoryQueryRequest Setter
// 库存查询参数
func (r *AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest) SetAicinventoryQueryRequest(_aicinventoryQueryRequest *Aicinventoryqueryrequest) error {
	r._aicinventoryQueryRequest = _aicinventoryQueryRequest
	r.Set("aicinventory_query_request", _aicinventoryQueryRequest)
	return nil
}

// GetAicinventoryQueryRequest AicinventoryQueryRequest Getter
func (r AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest) GetAicinventoryQueryRequest() *Aicinventoryqueryrequest {
	return r._aicinventoryQueryRequest
}

var poolAlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpAicSupplierAicinventoryNegativeSaleQueryRequest()
	},
}

// GetAlibabaAscpAicSupplierAicinventoryNegativeSaleQueryRequest 从 sync.Pool 获取 AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest
func GetAlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest() *AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest {
	return poolAlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest.Get().(*AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest)
}

// ReleaseAlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest 将 AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest(v *AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest) {
	v.Reset()
	poolAlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest.Put(v)
}
