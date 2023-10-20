package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabuseropenuidgetAPIRequest access token 获取精灵用户 id API请求
// alibaba.ailab.user.open.uid.get
//
// access token 获取精灵用户 id
type AlibabaailabuseropenuidgetAPIRequest struct {
	model.Params
	// access token
	_skillAccessToken string
	// skill id
	_skillId int64
}

// NewAlibabaailabuseropenuidgetRequest 初始化AlibabaailabuseropenuidgetAPIRequest对象
func NewAlibabaailabuseropenuidgetRequest() *AlibabaailabuseropenuidgetAPIRequest {
	return &AlibabaailabuseropenuidgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabuseropenuidgetAPIRequest) GetApiMethodName() string {
	return "alibaba.ailab.user.open.uid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabuseropenuidgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabuseropenuidgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkillAccessToken is SkillAccessToken Setter
// access token
func (r *AlibabaailabuseropenuidgetAPIRequest) SetSkillAccessToken(_skillAccessToken string) error {
	r._skillAccessToken = _skillAccessToken
	r.Set("skill_access_token", _skillAccessToken)
	return nil
}

// GetSkillAccessToken SkillAccessToken Getter
func (r AlibabaailabuseropenuidgetAPIRequest) GetSkillAccessToken() string {
	return r._skillAccessToken
}

// SetSkillId is SkillId Setter
// skill id
func (r *AlibabaailabuseropenuidgetAPIRequest) SetSkillId(_skillId int64) error {
	r._skillId = _skillId
	r.Set("skill_id", _skillId)
	return nil
}

// GetSkillId SkillId Getter
func (r AlibabaailabuseropenuidgetAPIRequest) GetSkillId() int64 {
	return r._skillId
}
