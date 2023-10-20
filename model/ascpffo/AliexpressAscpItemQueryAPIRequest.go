package ascpffo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressascpitemqueryAPIRequest AliExpress货品查询查询API API请求
// aliexpress.ascp.item.query
//
// AE货品查询API
type AliexpressascpitemqueryAPIRequest struct {
	model.Params
	// DTO
	_scItemQuery *ScItemQueryDto
}

// NewAliexpressascpitemqueryRequest 初始化AliexpressascpitemqueryAPIRequest对象
func NewAliexpressascpitemqueryRequest() *AliexpressascpitemqueryAPIRequest {
	return &AliexpressascpitemqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressascpitemqueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.ascp.item.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressascpitemqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressascpitemqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetScItemQuery is ScItemQuery Setter
// DTO
func (r *AliexpressascpitemqueryAPIRequest) SetScItemQuery(_scItemQuery *ScItemQueryDto) error {
	r._scItemQuery = _scItemQuery
	r.Set("sc_item_query", _scItemQuery)
	return nil
}

// GetScItemQuery ScItemQuery Getter
func (r AliexpressascpitemqueryAPIRequest) GetScItemQuery() *ScItemQueryDto {
	return r._scItemQuery
}
