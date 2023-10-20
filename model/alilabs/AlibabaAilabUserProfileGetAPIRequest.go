package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabuserprofilegetAPIRequest 查询用户信息 API请求
// alibaba.ailab.user.profile.get
//
// 提供天猫精灵用户头像、昵称的查询接口，供本田车载天猫精灵使用
type AlibabaailabuserprofilegetAPIRequest struct {
	model.Params
	// open uid
	_openUid string
	// client id
	_clientId string
}

// NewAlibabaailabuserprofilegetRequest 初始化AlibabaailabuserprofilegetAPIRequest对象
func NewAlibabaailabuserprofilegetRequest() *AlibabaailabuserprofilegetAPIRequest {
	return &AlibabaailabuserprofilegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabuserprofilegetAPIRequest) GetApiMethodName() string {
	return "alibaba.ailab.user.profile.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabuserprofilegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabuserprofilegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenUid is OpenUid Setter
// open uid
func (r *AlibabaailabuserprofilegetAPIRequest) SetOpenUid(_openUid string) error {
	r._openUid = _openUid
	r.Set("open_uid", _openUid)
	return nil
}

// GetOpenUid OpenUid Getter
func (r AlibabaailabuserprofilegetAPIRequest) GetOpenUid() string {
	return r._openUid
}

// SetClientId is ClientId Setter
// client id
func (r *AlibabaailabuserprofilegetAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabaailabuserprofilegetAPIRequest) GetClientId() string {
	return r._clientId
}
