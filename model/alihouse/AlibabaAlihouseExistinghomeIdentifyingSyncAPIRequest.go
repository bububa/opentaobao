package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomeidentifyingsyncAPIRequest 登陆标识信息同步 API请求
// alibaba.alihouse.existinghome.identifying.sync
//
// 登陆标识信息同步
type AlibabaalihouseexistinghomeidentifyingsyncAPIRequest struct {
	model.Params
	// 身份信息
	_identityInfo *IdentityInfoDto
}

// NewAlibabaalihouseexistinghomeidentifyingsyncRequest 初始化AlibabaalihouseexistinghomeidentifyingsyncAPIRequest对象
func NewAlibabaalihouseexistinghomeidentifyingsyncRequest() *AlibabaalihouseexistinghomeidentifyingsyncAPIRequest {
	return &AlibabaalihouseexistinghomeidentifyingsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomeidentifyingsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.identifying.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomeidentifyingsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomeidentifyingsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIdentityInfo is IdentityInfo Setter
// 身份信息
func (r *AlibabaalihouseexistinghomeidentifyingsyncAPIRequest) SetIdentityInfo(_identityInfo *IdentityInfoDto) error {
	r._identityInfo = _identityInfo
	r.Set("identity_info", _identityInfo)
	return nil
}

// GetIdentityInfo IdentityInfo Getter
func (r AlibabaalihouseexistinghomeidentifyingsyncAPIRequest) GetIdentityInfo() *IdentityInfoDto {
	return r._identityInfo
}
