package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceYksSkillModifyAPIRequest 修改技能 API请求
// yunos.tvpubadmin.device.yks.skill.modify
//
// 修改技能
type YunosTvpubadminDeviceYksSkillModifyAPIRequest struct {
	model.Params
	// 图片地址
	_iconImageUrl string
	// 技能名称
	_name string
	// 技能id
	_skillId int64
}

// NewYunosTvpubadminDeviceYksSkillModifyRequest 初始化YunosTvpubadminDeviceYksSkillModifyAPIRequest对象
func NewYunosTvpubadminDeviceYksSkillModifyRequest() *YunosTvpubadminDeviceYksSkillModifyAPIRequest {
	return &YunosTvpubadminDeviceYksSkillModifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceYksSkillModifyAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.yks.skill.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceYksSkillModifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminDeviceYksSkillModifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIconImageUrl is IconImageUrl Setter
// 图片地址
func (r *YunosTvpubadminDeviceYksSkillModifyAPIRequest) SetIconImageUrl(_iconImageUrl string) error {
	r._iconImageUrl = _iconImageUrl
	r.Set("icon_image_url", _iconImageUrl)
	return nil
}

// GetIconImageUrl IconImageUrl Getter
func (r YunosTvpubadminDeviceYksSkillModifyAPIRequest) GetIconImageUrl() string {
	return r._iconImageUrl
}

// SetName is Name Setter
// 技能名称
func (r *YunosTvpubadminDeviceYksSkillModifyAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r YunosTvpubadminDeviceYksSkillModifyAPIRequest) GetName() string {
	return r._name
}

// SetSkillId is SkillId Setter
// 技能id
func (r *YunosTvpubadminDeviceYksSkillModifyAPIRequest) SetSkillId(_skillId int64) error {
	r._skillId = _skillId
	r.Set("skill_id", _skillId)
	return nil
}

// GetSkillId SkillId Getter
func (r YunosTvpubadminDeviceYksSkillModifyAPIRequest) GetSkillId() int64 {
	return r._skillId
}
