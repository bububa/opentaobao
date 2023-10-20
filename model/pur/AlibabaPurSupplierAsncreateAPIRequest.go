package pur

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPurSupplierAsncreateAPIRequest asn创建 API请求
// alibaba.pur.supplier.asncreate
//
// asn创建
type AlibabaPurSupplierAsncreateAPIRequest struct {
	model.Params
	// asn头信息
	_asn *SupplierAsnInfoVo
}

// NewAlibabaPurSupplierAsncreateRequest 初始化AlibabaPurSupplierAsncreateAPIRequest对象
func NewAlibabaPurSupplierAsncreateRequest() *AlibabaPurSupplierAsncreateAPIRequest {
	return &AlibabaPurSupplierAsncreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaPurSupplierAsncreateAPIRequest) Reset() {
	r._asn = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPurSupplierAsncreateAPIRequest) GetApiMethodName() string {
	return "alibaba.pur.supplier.asncreate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPurSupplierAsncreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaPurSupplierAsncreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAsn is Asn Setter
// asn头信息
func (r *AlibabaPurSupplierAsncreateAPIRequest) SetAsn(_asn *SupplierAsnInfoVo) error {
	r._asn = _asn
	r.Set("asn", _asn)
	return nil
}

// GetAsn Asn Getter
func (r AlibabaPurSupplierAsncreateAPIRequest) GetAsn() *SupplierAsnInfoVo {
	return r._asn
}

var poolAlibabaPurSupplierAsncreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaPurSupplierAsncreateRequest()
	},
}

// GetAlibabaPurSupplierAsncreateRequest 从 sync.Pool 获取 AlibabaPurSupplierAsncreateAPIRequest
func GetAlibabaPurSupplierAsncreateAPIRequest() *AlibabaPurSupplierAsncreateAPIRequest {
	return poolAlibabaPurSupplierAsncreateAPIRequest.Get().(*AlibabaPurSupplierAsncreateAPIRequest)
}

// ReleaseAlibabaPurSupplierAsncreateAPIRequest 将 AlibabaPurSupplierAsncreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaPurSupplierAsncreateAPIRequest(v *AlibabaPurSupplierAsncreateAPIRequest) {
	v.Reset()
	poolAlibabaPurSupplierAsncreateAPIRequest.Put(v)
}
