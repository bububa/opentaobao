package ascpffo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressAscpRoQueryAPIRequest
AliExpress退供单查询API API请求
aliexpress.ascp.ro.query

AE仓发商家单个退供单查询接口 */
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
func (r AliexpressAscpRoQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ReturnOrderQuery Setter
// dto
func (r *AliexpressAscpRoQueryAPIRequest) SetReturnOrderQuery(_returnOrderQuery *ReturnOrderQueryDto) error {
	r._returnOrderQuery = _returnOrderQuery
	r.Set("return_order_query", _returnOrderQuery)
	return nil
}

// Get ReturnOrderQuery Getter
func (r AliexpressAscpRoQueryAPIRequest) GetReturnOrderQuery() *ReturnOrderQueryDto {
	return r._returnOrderQuery
}
