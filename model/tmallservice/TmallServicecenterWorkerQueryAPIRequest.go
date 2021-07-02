package tmallservice

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkerQueryAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.worker.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkerQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
