package mos

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosFinanceBankinfoQuerybankAPIRequest 供应商银行账号查询 API请求
// alibaba.mos.finance.bankinfo.querybank
//
// 查询供应商对应的银行账号信息
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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMosFinanceBankinfoQuerybankAPIRequest) Reset() {
	r._supplierId = ""
	r._storeNo = ""
	r._companyId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosFinanceBankinfoQuerybankAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.finance.bankinfo.querybank"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosFinanceBankinfoQuerybankAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMosFinanceBankinfoQuerybankAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSupplierId is SupplierId Setter
// 供应商id
func (r *AlibabaMosFinanceBankinfoQuerybankAPIRequest) SetSupplierId(_supplierId string) error {
	r._supplierId = _supplierId
	r.Set("supplier_id", _supplierId)
	return nil
}

// GetSupplierId SupplierId Getter
func (r AlibabaMosFinanceBankinfoQuerybankAPIRequest) GetSupplierId() string {
	return r._supplierId
}

// SetStoreNo is StoreNo Setter
// 门店号
func (r *AlibabaMosFinanceBankinfoQuerybankAPIRequest) SetStoreNo(_storeNo string) error {
	r._storeNo = _storeNo
	r.Set("store_no", _storeNo)
	return nil
}

// GetStoreNo StoreNo Getter
func (r AlibabaMosFinanceBankinfoQuerybankAPIRequest) GetStoreNo() string {
	return r._storeNo
}

// SetCompanyId is CompanyId Setter
// 签约主体id
func (r *AlibabaMosFinanceBankinfoQuerybankAPIRequest) SetCompanyId(_companyId string) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// GetCompanyId CompanyId Getter
func (r AlibabaMosFinanceBankinfoQuerybankAPIRequest) GetCompanyId() string {
	return r._companyId
}

var poolAlibabaMosFinanceBankinfoQuerybankAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMosFinanceBankinfoQuerybankRequest()
	},
}

// GetAlibabaMosFinanceBankinfoQuerybankRequest 从 sync.Pool 获取 AlibabaMosFinanceBankinfoQuerybankAPIRequest
func GetAlibabaMosFinanceBankinfoQuerybankAPIRequest() *AlibabaMosFinanceBankinfoQuerybankAPIRequest {
	return poolAlibabaMosFinanceBankinfoQuerybankAPIRequest.Get().(*AlibabaMosFinanceBankinfoQuerybankAPIRequest)
}

// ReleaseAlibabaMosFinanceBankinfoQuerybankAPIRequest 将 AlibabaMosFinanceBankinfoQuerybankAPIRequest 放入 sync.Pool
func ReleaseAlibabaMosFinanceBankinfoQuerybankAPIRequest(v *AlibabaMosFinanceBankinfoQuerybankAPIRequest) {
	v.Reset()
	poolAlibabaMosFinanceBankinfoQuerybankAPIRequest.Put(v)
}
