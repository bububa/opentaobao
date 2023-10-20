package alilabs

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest 设备列表更新通知 API请求
// alibaba.ailabs.iot.device.list.update.notify
//
// 用于人工智能实验室IoT合作厂商上报三方接入IoT设备列表更新时的通知
type AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest struct {
	model.Params
	// 用户OAuth授权后的token
	_token string
	// 厂商在天猫精灵开放平台申请的技能id
	_skillId string
	// 更新类型 1：设备更新 2：场景更新
	_type string
}

// NewAlibabaAilabsIotDeviceListUpdateNotifyRequest 初始化AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest对象
func NewAlibabaAilabsIotDeviceListUpdateNotifyRequest() *AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest {
	return &AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest) Reset() {
	r._token = ""
	r._skillId = ""
	r._type = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.iot.device.list.update.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// 用户OAuth授权后的token
func (r *AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest) GetToken() string {
	return r._token
}

// SetSkillId is SkillId Setter
// 厂商在天猫精灵开放平台申请的技能id
func (r *AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest) SetSkillId(_skillId string) error {
	r._skillId = _skillId
	r.Set("skill_id", _skillId)
	return nil
}

// GetSkillId SkillId Getter
func (r AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest) GetSkillId() string {
	return r._skillId
}

// SetType is Type Setter
// 更新类型 1：设备更新 2：场景更新
func (r *AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest) GetType() string {
	return r._type
}

var poolAlibabaAilabsIotDeviceListUpdateNotifyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAilabsIotDeviceListUpdateNotifyRequest()
	},
}

// GetAlibabaAilabsIotDeviceListUpdateNotifyRequest 从 sync.Pool 获取 AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest
func GetAlibabaAilabsIotDeviceListUpdateNotifyAPIRequest() *AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest {
	return poolAlibabaAilabsIotDeviceListUpdateNotifyAPIRequest.Get().(*AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest)
}

// ReleaseAlibabaAilabsIotDeviceListUpdateNotifyAPIRequest 将 AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest 放入 sync.Pool
func ReleaseAlibabaAilabsIotDeviceListUpdateNotifyAPIRequest(v *AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest) {
	v.Reset()
	poolAlibabaAilabsIotDeviceListUpdateNotifyAPIRequest.Put(v)
}
