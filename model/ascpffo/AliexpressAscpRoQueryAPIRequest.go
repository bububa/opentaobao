package ascpffo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAscpRoQueryAPIRequest AliExpress退供单查询API API请求
// aliexpress.ascp.ro.query
//
// AE仓发商家单个退供单查询接口
type AliexpressAscpRoQueryAPIRequest struct {
	model.Params
	// dto
	_returnOrderQuery *ReturnOrderQueryDto
}

// NewAliexpressAscpRoQueryRequest 初始化AliexpressAscpRoQueryAPIRequest对象
func NewAliexpressAscpRoQueryRequest() *AliexpressAscpRoQueryAPIRequest {
	return &AliexpressAscpRoQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressAscpRoQueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.ascp.ro.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressAscpRoQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressAscpRoQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReturnOrderQuery is ReturnOrderQuery Setter
// dto
func (r *AliexpressAscpRoQueryAPIRequest) SetReturnOrderQuery(_returnOrderQuery *ReturnOrderQueryDto) error {
	r._returnOrderQuery = _returnOrderQuery
	r.Set("return_order_query", _returnOrderQuery)
	return nil
}

// GetReturnOrderQuery ReturnOrderQuery Getter
func (r AliexpressAscpRoQueryAPIRequest) GetReturnOrderQuery() *ReturnOrderQueryDto {
	return r._returnOrderQuery
}
