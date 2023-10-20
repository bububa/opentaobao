package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosfinancebankinfoquerybankAPIRequest 供应商银行账号查询 API请求
// alibaba.mos.finance.bankinfo.querybank
//
// 查询供应商对应的银行账号信息
type AlibabamosfinancebankinfoquerybankAPIRequest struct {
	model.Params
	// 供应商id
	_supplierId string
	// 门店号
	_storeNo string
	// 签约主体id
	_companyId string
}

// NewAlibabamosfinancebankinfoquerybankRequest 初始化AlibabamosfinancebankinfoquerybankAPIRequest对象
func NewAlibabamosfinancebankinfoquerybankRequest() *AlibabamosfinancebankinfoquerybankAPIRequest {
	return &AlibabamosfinancebankinfoquerybankAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamosfinancebankinfoquerybankAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.finance.bankinfo.querybank"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamosfinancebankinfoquerybankAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamosfinancebankinfoquerybankAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSupplierId is SupplierId Setter
// 供应商id
func (r *AlibabamosfinancebankinfoquerybankAPIRequest) SetSupplierId(_supplierId string) error {
	r._supplierId = _supplierId
	r.Set("supplier_id", _supplierId)
	return nil
}

// GetSupplierId SupplierId Getter
func (r AlibabamosfinancebankinfoquerybankAPIRequest) GetSupplierId() string {
	return r._supplierId
}

// SetStoreNo is StoreNo Setter
// 门店号
func (r *AlibabamosfinancebankinfoquerybankAPIRequest) SetStoreNo(_storeNo string) error {
	r._storeNo = _storeNo
	r.Set("store_no", _storeNo)
	return nil
}

// GetStoreNo StoreNo Getter
func (r AlibabamosfinancebankinfoquerybankAPIRequest) GetStoreNo() string {
	return r._storeNo
}

// SetCompanyId is CompanyId Setter
// 签约主体id
func (r *AlibabamosfinancebankinfoquerybankAPIRequest) SetCompanyId(_companyId string) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// GetCompanyId CompanyId Getter
func (r AlibabamosfinancebankinfoquerybankAPIRequest) GetCompanyId() string {
	return r._companyId
}
