package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkerqueryAPIRequest 工人信息查询 API请求
// tmall.servicecenter.worker.query
//
// 查询服务商对应的工人信息
type TmallservicecenterworkerqueryAPIRequest struct {
	model.Params
	// 身份证号
	_identityId string
}

// NewTmallservicecenterworkerqueryRequest 初始化TmallservicecenterworkerqueryAPIRequest对象
func NewTmallservicecenterworkerqueryRequest() *TmallservicecenterworkerqueryAPIRequest {
	return &TmallservicecenterworkerqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterworkerqueryAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.worker.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterworkerqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterworkerqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIdentityId is IdentityId Setter
// 身份证号
func (r *TmallservicecenterworkerqueryAPIRequest) SetIdentityId(_identityId string) error {
	r._identityId = _identityId
	r.Set("identity_id", _identityId)
	return nil
}

// GetIdentityId IdentityId Getter
func (r TmallservicecenterworkerqueryAPIRequest) GetIdentityId() string {
	return r._identityId
}
