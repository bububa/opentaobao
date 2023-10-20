package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkerQueryAPIRequest 工人信息查询 API请求
// tmall.servicecenter.worker.query
//
// 查询服务商对应的工人信息
type TmallServicecenterWorkerQueryAPIRequest struct {
	model.Params
	// 身份证号
	_identityId string
}

// NewTmallServicecenterWorkerQueryRequest 初始化TmallServicecenterWorkerQueryAPIRequest对象
func NewTmallServicecenterWorkerQueryRequest() *TmallServicecenterWorkerQueryAPIRequest {
	return &TmallServicecenterWorkerQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterWorkerQueryAPIRequest) Reset() {
	r._identityId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkerQueryAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.worker.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkerQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterWorkerQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIdentityId is IdentityId Setter
// 身份证号
func (r *TmallServicecenterWorkerQueryAPIRequest) SetIdentityId(_identityId string) error {
	r._identityId = _identityId
	r.Set("identity_id", _identityId)
	return nil
}

// GetIdentityId IdentityId Getter
func (r TmallServicecenterWorkerQueryAPIRequest) GetIdentityId() string {
	return r._identityId
}

var poolTmallServicecenterWorkerQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterWorkerQueryRequest()
	},
}

// GetTmallServicecenterWorkerQueryRequest 从 sync.Pool 获取 TmallServicecenterWorkerQueryAPIRequest
func GetTmallServicecenterWorkerQueryAPIRequest() *TmallServicecenterWorkerQueryAPIRequest {
	return poolTmallServicecenterWorkerQueryAPIRequest.Get().(*TmallServicecenterWorkerQueryAPIRequest)
}

// ReleaseTmallServicecenterWorkerQueryAPIRequest 将 TmallServicecenterWorkerQueryAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterWorkerQueryAPIRequest(v *TmallServicecenterWorkerQueryAPIRequest) {
	v.Reset()
	poolTmallServicecenterWorkerQueryAPIRequest.Put(v)
}
