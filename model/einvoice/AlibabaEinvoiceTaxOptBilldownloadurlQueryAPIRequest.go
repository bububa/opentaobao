package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoicetaxoptbilldownloadurlqueryAPIRequest 税筹业务账单文件下载URL查询 API请求
// alibaba.einvoice.tax.opt.billdownloadurl.query
//
// 税筹业务账单文件下载的URL查询
type AlibabaeinvoicetaxoptbilldownloadurlqueryAPIRequest struct {
	model.Params
	// 指定账单的日期
	_billDate string
	// 平台提供的公司编码
	_companyCode string
	// 用户类型，建议传递，供应商请传递固定值:CONTRACTOR
	_userType string
	// 供应商提供服务的的合作企业的公司编码，当user_type为CONTRACTOR时，建议提供此参数。特别是供应商使用同一主体编码面向多个企业提供服务，务必提供此参数，用于明确区分需要下载哪个合作企业的业务账单。
	_belongingBusinessScenario string
}

// NewAlibabaeinvoicetaxoptbilldownloadurlqueryRequest 初始化AlibabaeinvoicetaxoptbilldownloadurlqueryAPIRequest对象
func NewAlibabaeinvoicetaxoptbilldownloadurlqueryRequest() *AlibabaeinvoicetaxoptbilldownloadurlqueryAPIRequest {
	return &AlibabaeinvoicetaxoptbilldownloadurlqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoicetaxoptbilldownloadurlqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.tax.opt.billdownloadurl.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoicetaxoptbilldownloadurlqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoicetaxoptbilldownloadurlqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBillDate is BillDate Setter
// 指定账单的日期
func (r *AlibabaeinvoicetaxoptbilldownloadurlqueryAPIRequest) SetBillDate(_billDate string) error {
	r._billDate = _billDate
	r.Set("bill_date", _billDate)
	return nil
}

// GetBillDate BillDate Getter
func (r AlibabaeinvoicetaxoptbilldownloadurlqueryAPIRequest) GetBillDate() string {
	return r._billDate
}

// SetCompanyCode is CompanyCode Setter
// 平台提供的公司编码
func (r *AlibabaeinvoicetaxoptbilldownloadurlqueryAPIRequest) SetCompanyCode(_companyCode string) error {
	r._companyCode = _companyCode
	r.Set("company_code", _companyCode)
	return nil
}

// GetCompanyCode CompanyCode Getter
func (r AlibabaeinvoicetaxoptbilldownloadurlqueryAPIRequest) GetCompanyCode() string {
	return r._companyCode
}

// SetUserType is UserType Setter
// 用户类型，建议传递，供应商请传递固定值:CONTRACTOR
func (r *AlibabaeinvoicetaxoptbilldownloadurlqueryAPIRequest) SetUserType(_userType string) error {
	r._userType = _userType
	r.Set("user_type", _userType)
	return nil
}

// GetUserType UserType Getter
func (r AlibabaeinvoicetaxoptbilldownloadurlqueryAPIRequest) GetUserType() string {
	return r._userType
}

// SetBelongingBusinessScenario is BelongingBusinessScenario Setter
// 供应商提供服务的的合作企业的公司编码，当user_type为CONTRACTOR时，建议提供此参数。特别是供应商使用同一主体编码面向多个企业提供服务，务必提供此参数，用于明确区分需要下载哪个合作企业的业务账单。
func (r *AlibabaeinvoicetaxoptbilldownloadurlqueryAPIRequest) SetBelongingBusinessScenario(_belongingBusinessScenario string) error {
	r._belongingBusinessScenario = _belongingBusinessScenario
	r.Set("belonging_business_scenario", _belongingBusinessScenario)
	return nil
}

// GetBelongingBusinessScenario BelongingBusinessScenario Getter
func (r AlibabaeinvoicetaxoptbilldownloadurlqueryAPIRequest) GetBelongingBusinessScenario() string {
	return r._belongingBusinessScenario
}
