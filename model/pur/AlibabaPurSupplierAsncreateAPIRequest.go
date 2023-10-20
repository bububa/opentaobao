package pur

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabapursupplierasncreateAPIRequest asn创建 API请求
// alibaba.pur.supplier.asncreate
//
// asn创建
type AlibabapursupplierasncreateAPIRequest struct {
	model.Params
	// asn头信息
	_asn *SupplierAsnInfoVo
}

// NewAlibabapursupplierasncreateRequest 初始化AlibabapursupplierasncreateAPIRequest对象
func NewAlibabapursupplierasncreateRequest() *AlibabapursupplierasncreateAPIRequest {
	return &AlibabapursupplierasncreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabapursupplierasncreateAPIRequest) GetApiMethodName() string {
	return "alibaba.pur.supplier.asncreate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabapursupplierasncreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabapursupplierasncreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAsn is Asn Setter
// asn头信息
func (r *AlibabapursupplierasncreateAPIRequest) SetAsn(_asn *SupplierAsnInfoVo) error {
	r._asn = _asn
	r.Set("asn", _asn)
	return nil
}

// GetAsn Asn Getter
func (r AlibabapursupplierasncreateAPIRequest) GetAsn() *SupplierAsnInfoVo {
	return r._asn
}
