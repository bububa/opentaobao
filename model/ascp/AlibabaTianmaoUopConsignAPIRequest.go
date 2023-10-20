package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTianmaoUopConsignAPIRequest 阿里巴巴.天猫. 履约订单. 发货 API请求
// alibaba.tianmao.uop.consign
//
// 阿里巴巴.天猫. 履约订单. 发货
type AlibabaTianmaoUopConsignAPIRequest struct {
	model.Params
	// SCP单号
	_scpCode string
	// 货主id
	_ownerId int64
}

// NewAlibabaTianmaoUopConsignRequest 初始化AlibabaTianmaoUopConsignAPIRequest对象
func NewAlibabaTianmaoUopConsignRequest() *AlibabaTianmaoUopConsignAPIRequest {
	return &AlibabaTianmaoUopConsignAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTianmaoUopConsignAPIRequest) Reset() {
	r._scpCode = ""
	r._ownerId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTianmaoUopConsignAPIRequest) GetApiMethodName() string {
	return "alibaba.tianmao.uop.consign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTianmaoUopConsignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTianmaoUopConsignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetScpCode is ScpCode Setter
// SCP单号
func (r *AlibabaTianmaoUopConsignAPIRequest) SetScpCode(_scpCode string) error {
	r._scpCode = _scpCode
	r.Set("scp_code", _scpCode)
	return nil
}

// GetScpCode ScpCode Getter
func (r AlibabaTianmaoUopConsignAPIRequest) GetScpCode() string {
	return r._scpCode
}

// SetOwnerId is OwnerId Setter
// 货主id
func (r *AlibabaTianmaoUopConsignAPIRequest) SetOwnerId(_ownerId int64) error {
	r._ownerId = _ownerId
	r.Set("owner_id", _ownerId)
	return nil
}

// GetOwnerId OwnerId Getter
func (r AlibabaTianmaoUopConsignAPIRequest) GetOwnerId() int64 {
	return r._ownerId
}

var poolAlibabaTianmaoUopConsignAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTianmaoUopConsignRequest()
	},
}

// GetAlibabaTianmaoUopConsignRequest 从 sync.Pool 获取 AlibabaTianmaoUopConsignAPIRequest
func GetAlibabaTianmaoUopConsignAPIRequest() *AlibabaTianmaoUopConsignAPIRequest {
	return poolAlibabaTianmaoUopConsignAPIRequest.Get().(*AlibabaTianmaoUopConsignAPIRequest)
}

// ReleaseAlibabaTianmaoUopConsignAPIRequest 将 AlibabaTianmaoUopConsignAPIRequest 放入 sync.Pool
func ReleaseAlibabaTianmaoUopConsignAPIRequest(v *AlibabaTianmaoUopConsignAPIRequest) {
	v.Reset()
	poolAlibabaTianmaoUopConsignAPIRequest.Put(v)
}
