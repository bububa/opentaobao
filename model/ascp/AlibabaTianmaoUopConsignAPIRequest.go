package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatianmaouopconsignAPIRequest 阿里巴巴.天猫. 履约订单. 发货 API请求
// alibaba.tianmao.uop.consign
//
// 阿里巴巴.天猫. 履约订单. 发货
type AlibabatianmaouopconsignAPIRequest struct {
	model.Params
	// SCP单号
	_scpCode string
	// 货主id
	_ownerId int64
}

// NewAlibabatianmaouopconsignRequest 初始化AlibabatianmaouopconsignAPIRequest对象
func NewAlibabatianmaouopconsignRequest() *AlibabatianmaouopconsignAPIRequest {
	return &AlibabatianmaouopconsignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatianmaouopconsignAPIRequest) GetApiMethodName() string {
	return "alibaba.tianmao.uop.consign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatianmaouopconsignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatianmaouopconsignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetScpCode is ScpCode Setter
// SCP单号
func (r *AlibabatianmaouopconsignAPIRequest) SetScpCode(_scpCode string) error {
	r._scpCode = _scpCode
	r.Set("scp_code", _scpCode)
	return nil
}

// GetScpCode ScpCode Getter
func (r AlibabatianmaouopconsignAPIRequest) GetScpCode() string {
	return r._scpCode
}

// SetOwnerId is OwnerId Setter
// 货主id
func (r *AlibabatianmaouopconsignAPIRequest) SetOwnerId(_ownerId int64) error {
	r._ownerId = _ownerId
	r.Set("owner_id", _ownerId)
	return nil
}

// GetOwnerId OwnerId Getter
func (r AlibabatianmaouopconsignAPIRequest) GetOwnerId() int64 {
	return r._ownerId
}
