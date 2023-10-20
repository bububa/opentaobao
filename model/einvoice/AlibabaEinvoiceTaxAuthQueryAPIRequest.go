package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoicetaxauthqueryAPIRequest 发票中台授权信息获取 API请求
// alibaba.einvoice.tax.auth.query
//
// 发票中台授权信息获取
type AlibabaeinvoicetaxauthqueryAPIRequest struct {
	model.Params
	// 授权信息查询请求
	_taxAuthTokenQueryDto *TaxAuthTokenQueryDto
}

// NewAlibabaeinvoicetaxauthqueryRequest 初始化AlibabaeinvoicetaxauthqueryAPIRequest对象
func NewAlibabaeinvoicetaxauthqueryRequest() *AlibabaeinvoicetaxauthqueryAPIRequest {
	return &AlibabaeinvoicetaxauthqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoicetaxauthqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.tax.auth.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoicetaxauthqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoicetaxauthqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTaxAuthTokenQueryDto is TaxAuthTokenQueryDto Setter
// 授权信息查询请求
func (r *AlibabaeinvoicetaxauthqueryAPIRequest) SetTaxAuthTokenQueryDto(_taxAuthTokenQueryDto *TaxAuthTokenQueryDto) error {
	r._taxAuthTokenQueryDto = _taxAuthTokenQueryDto
	r.Set("tax_auth_token_query_dto", _taxAuthTokenQueryDto)
	return nil
}

// GetTaxAuthTokenQueryDto TaxAuthTokenQueryDto Getter
func (r AlibabaeinvoicetaxauthqueryAPIRequest) GetTaxAuthTokenQueryDto() *TaxAuthTokenQueryDto {
	return r._taxAuthTokenQueryDto
}
