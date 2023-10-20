package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmindeviceyksskilladdAPIRequest 添加技能 API请求
// yunos.tvpubadmin.device.yks.skill.add
//
// 添加技能
type YunostvpubadmindeviceyksskilladdAPIRequest struct {
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

// NewYunostvpubadmindeviceyksskilladdRequest 初始化YunostvpubadmindeviceyksskilladdAPIRequest对象
func NewYunostvpubadmindeviceyksskilladdRequest() *YunostvpubadmindeviceyksskilladdAPIRequest {
	return &YunostvpubadmindeviceyksskilladdAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmindeviceyksskilladdAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.yks.skill.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmindeviceyksskilladdAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmindeviceyksskilladdAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 技能名称
func (r *YunostvpubadmindeviceyksskilladdAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r YunostvpubadmindeviceyksskilladdAPIRequest) GetName() string {
	return r._name
}

// SetIconImageUrl is IconImageUrl Setter
// 图片地址
func (r *YunostvpubadmindeviceyksskilladdAPIRequest) SetIconImageUrl(_iconImageUrl string) error {
	r._iconImageUrl = _iconImageUrl
	r.Set("icon_image_url", _iconImageUrl)
	return nil
}

// GetIconImageUrl IconImageUrl Getter
func (r YunostvpubadmindeviceyksskilladdAPIRequest) GetIconImageUrl() string {
	return r._iconImageUrl
}

// SetSkillId is SkillId Setter
// 技能id
func (r *YunostvpubadmindeviceyksskilladdAPIRequest) SetSkillId(_skillId int64) error {
	r._skillId = _skillId
	r.Set("skill_id", _skillId)
	return nil
}

// GetSkillId SkillId Getter
func (r YunostvpubadmindeviceyksskilladdAPIRequest) GetSkillId() int64 {
	return r._skillId
}

// SetBotId is BotId Setter
// 设备id
func (r *YunostvpubadmindeviceyksskilladdAPIRequest) SetBotId(_botId int64) error {
	r._botId = _botId
	r.Set("bot_id", _botId)
	return nil
}

// GetBotId BotId Getter
func (r YunostvpubadmindeviceyksskilladdAPIRequest) GetBotId() int64 {
	return r._botId
}
