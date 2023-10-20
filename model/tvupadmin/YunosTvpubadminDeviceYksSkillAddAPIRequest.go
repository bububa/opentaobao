package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceYksSkillAddAPIRequest 添加技能 API请求
// yunos.tvpubadmin.device.yks.skill.add
//
// 添加技能
type YunosTvpubadminDeviceYksSkillAddAPIRequest struct {
	model.Params
	// 技能名称
	_name string
	// 图片地址
	_iconImageUrl string
	// 技能id
	_skillId int64
	// 设备id
	_botId int64
}

// NewYunosTvpubadminDeviceYksSkillAddRequest 初始化YunosTvpubadminDeviceYksSkillAddAPIRequest对象
func NewYunosTvpubadminDeviceYksSkillAddRequest() *YunosTvpubadminDeviceYksSkillAddAPIRequest {
	return &YunosTvpubadminDeviceYksSkillAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceYksSkillAddAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.yks.skill.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceYksSkillAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminDeviceYksSkillAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 技能名称
func (r *YunosTvpubadminDeviceYksSkillAddAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r YunosTvpubadminDeviceYksSkillAddAPIRequest) GetName() string {
	return r._name
}

// SetIconImageUrl is IconImageUrl Setter
// 图片地址
func (r *YunosTvpubadminDeviceYksSkillAddAPIRequest) SetIconImageUrl(_iconImageUrl string) error {
	r._iconImageUrl = _iconImageUrl
	r.Set("icon_image_url", _iconImageUrl)
	return nil
}

// GetIconImageUrl IconImageUrl Getter
func (r YunosTvpubadminDeviceYksSkillAddAPIRequest) GetIconImageUrl() string {
	return r._iconImageUrl
}

// SetSkillId is SkillId Setter
// 技能id
func (r *YunosTvpubadminDeviceYksSkillAddAPIRequest) SetSkillId(_skillId int64) error {
	r._skillId = _skillId
	r.Set("skill_id", _skillId)
	return nil
}

// GetSkillId SkillId Getter
func (r YunosTvpubadminDeviceYksSkillAddAPIRequest) GetSkillId() int64 {
	return r._skillId
}

// SetBotId is BotId Setter
// 设备id
func (r *YunosTvpubadminDeviceYksSkillAddAPIRequest) SetBotId(_botId int64) error {
	r._botId = _botId
	r.Set("bot_id", _botId)
	return nil
}

// GetBotId BotId Getter
func (r YunosTvpubadminDeviceYksSkillAddAPIRequest) GetBotId() int64 {
	return r._botId
}
