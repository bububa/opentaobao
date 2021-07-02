package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabUserProfileGetAPIRequest 查询用户信息 API请求
// alibaba.ailab.user.profile.get
//
// 提供天猫精灵用户头像、昵称的查询接口，供本田车载天猫精灵使用
type AlibabaAilabUserProfileGetAPIRequest struct {
	model.Params
	// open uid
	_openUid string
	// client id
	_clientId string
}

// NewAlibabaAilabUserProfileGetRequest 初始化AlibabaAilabUserProfileGetAPIRequest对象
func NewAlibabaAilabUserProfileGetRequest() *AlibabaAilabUserProfileGetAPIRequest {
	return &AlibabaAilabUserProfileGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabUserProfileGetAPIRequest) GetApiMethodName() string {
	return "alibaba.ailab.user.profile.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabUserProfileGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOpenUid is OpenUid Setter
// open uid
func (r *AlibabaAilabUserProfileGetAPIRequest) SetOpenUid(_openUid string) error {
	r._openUid = _openUid
	r.Set("open_uid", _openUid)
	return nil
}

// GetOpenUid OpenUid Getter
func (r AlibabaAilabUserProfileGetAPIRequest) GetOpenUid() string {
	return r._openUid
}

// SetClientId is ClientId Setter
// client id
func (r *AlibabaAilabUserProfileGetAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabaAilabUserProfileGetAPIRequest) GetClientId() string {
	return r._clientId
}
