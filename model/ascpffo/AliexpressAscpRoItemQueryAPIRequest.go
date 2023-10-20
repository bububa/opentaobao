package ascpffo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressascproitemqueryAPIRequest AliExpress退供单明细查询API API请求
// aliexpress.ascp.ro.item.query
//
// AE仓发 单个退供单明细查询
type AliexpressascproitemqueryAPIRequest struct {
	model.Params
	// dto
	_returnOrderItemQuery *ReturnOrderItemQueryDto
}

// NewAliexpressascproitemqueryRequest 初始化AliexpressascproitemqueryAPIRequest对象
func NewAliexpressascproitemqueryRequest() *AliexpressascproitemqueryAPIRequest {
	return &AliexpressascproitemqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressascproitemqueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.ascp.ro.item.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressascproitemqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressascproitemqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReturnOrderItemQuery is ReturnOrderItemQuery Setter
// dto
func (r *AliexpressascproitemqueryAPIRequest) SetReturnOrderItemQuery(_returnOrderItemQuery *ReturnOrderItemQueryDto) error {
	r._returnOrderItemQuery = _returnOrderItemQuery
	r.Set("return_order_item_query", _returnOrderItemQuery)
	return nil
}

// GetReturnOrderItemQuery ReturnOrderItemQuery Getter
func (r AliexpressascproitemqueryAPIRequest) GetReturnOrderItemQuery() *ReturnOrderItemQueryDto {
	return r._returnOrderItemQuery
}
