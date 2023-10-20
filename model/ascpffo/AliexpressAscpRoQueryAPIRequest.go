package ascpffo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressascproqueryAPIRequest AliExpress退供单查询API API请求
// aliexpress.ascp.ro.query
//
// AE仓发商家单个退供单查询接口
type AliexpressascproqueryAPIRequest struct {
	model.Params
	// dto
	_returnOrderQuery *ReturnOrderQueryDto
}

// NewAliexpressascproqueryRequest 初始化AliexpressascproqueryAPIRequest对象
func NewAliexpressascproqueryRequest() *AliexpressascproqueryAPIRequest {
	return &AliexpressascproqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressascproqueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.ascp.ro.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressascproqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressascproqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReturnOrderQuery is ReturnOrderQuery Setter
// dto
func (r *AliexpressascproqueryAPIRequest) SetReturnOrderQuery(_returnOrderQuery *ReturnOrderQueryDto) error {
	r._returnOrderQuery = _returnOrderQuery
	r.Set("return_order_query", _returnOrderQuery)
	return nil
}

// GetReturnOrderQuery ReturnOrderQuery Getter
func (r AliexpressascproqueryAPIRequest) GetReturnOrderQuery() *ReturnOrderQueryDto {
	return r._returnOrderQuery
}
