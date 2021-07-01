package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabUserOpenUidGetAPIRequest
access token 获取精灵用户 id API请求
alibaba.ailab.user.open.uid.get

access token 获取精灵用户 id */
type AlibabaAilabUserOpenUidGetAPIRequest struct {
	model.Params
	// access token
	_skillAccessToken string
	// skill id
	_skillId int64
}

// NewAlibabaAilabUserOpenUidGetRequest 初始化AlibabaAilabUserOpenUidGetAPIRequest对象
func NewAlibabaAilabUserOpenUidGetRequest() *AlibabaAilabUserOpenUidGetAPIRequest {
	return &AlibabaAilabUserOpenUidGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabUserOpenUidGetAPIRequest) GetApiMethodName() string {
	return "alibaba.ailab.user.open.uid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabUserOpenUidGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SkillAccessToken Setter
// access token
func (r *AlibabaAilabUserOpenUidGetAPIRequest) SetSkillAccessToken(_skillAccessToken string) error {
	r._skillAccessToken = _skillAccessToken
	r.Set("skill_access_token", _skillAccessToken)
	return nil
}

// Get SkillAccessToken Getter
func (r AlibabaAilabUserOpenUidGetAPIRequest) GetSkillAccessToken() string {
	return r._skillAccessToken
}

// Set is SkillId Setter
// skill id
func (r *AlibabaAilabUserOpenUidGetAPIRequest) SetSkillId(_skillId int64) error {
	r._skillId = _skillId
	r.Set("skill_id", _skillId)
	return nil
}

// Get SkillId Getter
func (r AlibabaAilabUserOpenUidGetAPIRequest) GetSkillId() int64 {
	return r._skillId
}
