package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpaicsupplieraicinventorynegativesalequeryAPIRequest 商家查询负卖库存 API请求
// alibaba.ascp.aic.supplier.aicinventory.negative.sale.query
//
// 商家根据当前接口查询负卖货品的库存
type AlibabaascpaicsupplieraicinventorynegativesalequeryAPIRequest struct {
	model.Params
	// 库存查询参数
	_aicinventoryQueryRequest *Aicinventoryqueryrequest
}

// NewAlibabaascpaicsupplieraicinventorynegativesalequeryRequest 初始化AlibabaascpaicsupplieraicinventorynegativesalequeryAPIRequest对象
func NewAlibabaascpaicsupplieraicinventorynegativesalequeryRequest() *AlibabaascpaicsupplieraicinventorynegativesalequeryAPIRequest {
	return &AlibabaascpaicsupplieraicinventorynegativesalequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpaicsupplieraicinventorynegativesalequeryAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.aic.supplier.aicinventory.negative.sale.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpaicsupplieraicinventorynegativesalequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpaicsupplieraicinventorynegativesalequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAicinventoryQueryRequest is AicinventoryQueryRequest Setter
// 库存查询参数
func (r *AlibabaascpaicsupplieraicinventorynegativesalequeryAPIRequest) SetAicinventoryQueryRequest(_aicinventoryQueryRequest *Aicinventoryqueryrequest) error {
	r._aicinventoryQueryRequest = _aicinventoryQueryRequest
	r.Set("aicinventory_query_request", _aicinventoryQueryRequest)
	return nil
}

// GetAicinventoryQueryRequest AicinventoryQueryRequest Getter
func (r AlibabaascpaicsupplieraicinventorynegativesalequeryAPIRequest) GetAicinventoryQueryRequest() *Aicinventoryqueryrequest {
	return r._aicinventoryQueryRequest
}
