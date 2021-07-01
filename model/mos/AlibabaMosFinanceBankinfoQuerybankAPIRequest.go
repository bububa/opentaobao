package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosFinanceBankinfoQuerybankAPIRequest
供应商银行账号查询 API请求
alibaba.mos.finance.bankinfo.querybank

查询供应商对应的银行账号信息 */
type AlibabaMosFinanceBankinfoQuerybankAPIRequest struct {
	model.Params
	// 供应商id
	_supplierId string
	// 门店号
	_storeNo string
	// 签约主体id
	_companyId string
}

// NewAlibabaMosFinanceBankinfoQuerybankRequest 初始化AlibabaMosFinanceBankinfoQuerybankAPIRequest对象
func NewAlibabaMosFinanceBankinfoQuerybankRequest() *AlibabaMosFinanceBankinfoQuerybankAPIRequest {
	return &AlibabaMosFinanceBankinfoQuerybankAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosFinanceBankinfoQuerybankAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.finance.bankinfo.querybank"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosFinanceBankinfoQuerybankAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SupplierId Setter
// 供应商id
func (r *AlibabaMosFinanceBankinfoQuerybankAPIRequest) SetSupplierId(_supplierId string) error {
	r._supplierId = _supplierId
	r.Set("supplier_id", _supplierId)
	return nil
}

// Get SupplierId Getter
func (r AlibabaMosFinanceBankinfoQuerybankAPIRequest) GetSupplierId() string {
	return r._supplierId
}

// Set is StoreNo Setter
// 门店号
func (r *AlibabaMosFinanceBankinfoQuerybankAPIRequest) SetStoreNo(_storeNo string) error {
	r._storeNo = _storeNo
	r.Set("store_no", _storeNo)
	return nil
}

// Get StoreNo Getter
func (r AlibabaMosFinanceBankinfoQuerybankAPIRequest) GetStoreNo() string {
	return r._storeNo
}

// Set is CompanyId Setter
// 签约主体id
func (r *AlibabaMosFinanceBankinfoQuerybankAPIRequest) SetCompanyId(_companyId string) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// Get CompanyId Getter
func (r AlibabaMosFinanceBankinfoQuerybankAPIRequest) GetCompanyId() string {
	return r._companyId
}
