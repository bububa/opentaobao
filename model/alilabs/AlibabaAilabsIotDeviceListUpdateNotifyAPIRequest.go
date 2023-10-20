package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabsiotdevicelistupdatenotifyAPIRequest 设备列表更新通知 API请求
// alibaba.ailabs.iot.device.list.update.notify
//
// 用于人工智能实验室IoT合作厂商上报三方接入IoT设备列表更新时的通知
type AlibabaailabsiotdevicelistupdatenotifyAPIRequest struct {
	model.Params
	// 用户OAuth授权后的token
	_token string
	// 厂商在天猫精灵开放平台申请的技能id
	_skillId string
	// 更新类型 1：设备更新 2：场景更新
	_type string
}

// NewAlibabaailabsiotdevicelistupdatenotifyRequest 初始化AlibabaailabsiotdevicelistupdatenotifyAPIRequest对象
func NewAlibabaailabsiotdevicelistupdatenotifyRequest() *AlibabaailabsiotdevicelistupdatenotifyAPIRequest {
	return &AlibabaailabsiotdevicelistupdatenotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabsiotdevicelistupdatenotifyAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.iot.device.list.update.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabsiotdevicelistupdatenotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabsiotdevicelistupdatenotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// 用户OAuth授权后的token
func (r *AlibabaailabsiotdevicelistupdatenotifyAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaailabsiotdevicelistupdatenotifyAPIRequest) GetToken() string {
	return r._token
}

// SetSkillId is SkillId Setter
// 厂商在天猫精灵开放平台申请的技能id
func (r *AlibabaailabsiotdevicelistupdatenotifyAPIRequest) SetSkillId(_skillId string) error {
	r._skillId = _skillId
	r.Set("skill_id", _skillId)
	return nil
}

// GetSkillId SkillId Getter
func (r AlibabaailabsiotdevicelistupdatenotifyAPIRequest) GetSkillId() string {
	return r._skillId
}

// SetType is Type Setter
// 更新类型 1：设备更新 2：场景更新
func (r *AlibabaailabsiotdevicelistupdatenotifyAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaailabsiotdevicelistupdatenotifyAPIRequest) GetType() string {
	return r._type
}
