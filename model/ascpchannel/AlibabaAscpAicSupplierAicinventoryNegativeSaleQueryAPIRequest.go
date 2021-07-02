package ascpchannel

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.aic.supplier.aicinventory.negative.sale.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AicinventoryQueryRequest Setter
// 库存查询参数
func (r *AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest) SetAicinventoryQueryRequest(_aicinventoryQueryRequest *Aicinventoryqueryrequest) error {
	r._aicinventoryQueryRequest = _aicinventoryQueryRequest
	r.Set("aicinventory_query_request", _aicinventoryQueryRequest)
	return nil
}

// Get AicinventoryQueryRequest Getter
func (r AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest) GetAicinventoryQueryRequest() *Aicinventoryqueryrequest {
	return r._aicinventoryQueryRequest
}
