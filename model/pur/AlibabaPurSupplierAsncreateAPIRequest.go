package pur

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPurSupplierAsncreateAPIRequest asn创建 API请求
// alibaba.pur.supplier.asncreate
//
// asn创建
type AlibabaPurSupplierAsncreateAPIRequest struct {
	model.Params
	// asn头信息
	_asn *SupplierAsnInfoVO
}

// NewAlibabaPurSupplierAsncreateRequest 初始化AlibabaPurSupplierAsncreateAPIRequest对象
func NewAlibabaPurSupplierAsncreateRequest() *AlibabaPurSupplierAsncreateAPIRequest {
	return &AlibabaPurSupplierAsncreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPurSupplierAsncreateAPIRequest) GetApiMethodName() string {
	return "alibaba.pur.supplier.asncreate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPurSupplierAsncreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Asn Setter
// asn头信息
func (r *AlibabaPurSupplierAsncreateAPIRequest) SetAsn(_asn *SupplierAsnInfoVO) error {
	r._asn = _asn
	r.Set("asn", _asn)
	return nil
}

// Get Asn Getter
func (r AlibabaPurSupplierAsncreateAPIRequest) GetAsn() *SupplierAsnInfoVO {
	return r._asn
}
