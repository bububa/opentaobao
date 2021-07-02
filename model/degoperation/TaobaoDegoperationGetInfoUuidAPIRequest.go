package degoperation

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDegoperationGetInfoUuidAPIRequest 根据uuid用户抽奖次数限制 API请求
// taobao.degoperation.get.info.uuid
//
// 根据uuid用户抽奖次数限制
type TaobaoDegoperationGetInfoUuidAPIRequest struct {
	model.Params
	// 活动后台配置eventkey
	_degAppKey string
	// 活动后台配置appkey
	_degEventKey string
	// 设备id
	_uuid string
}

// NewTaobaoDegoperationGetInfoUuidRequest 初始化TaobaoDegoperationGetInfoUuidAPIRequest对象
func NewTaobaoDegoperationGetInfoUuidRequest() *TaobaoDegoperationGetInfoUuidAPIRequest {
	return &TaobaoDegoperationGetInfoUuidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDegoperationGetInfoUuidAPIRequest) GetApiMethodName() string {
	return "taobao.degoperation.get.info.uuid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDegoperationGetInfoUuidAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDegAppKey is DegAppKey Setter
// 活动后台配置eventkey
func (r *TaobaoDegoperationGetInfoUuidAPIRequest) SetDegAppKey(_degAppKey string) error {
	r._degAppKey = _degAppKey
	r.Set("deg_app_key", _degAppKey)
	return nil
}

// GetDegAppKey DegAppKey Getter
func (r TaobaoDegoperationGetInfoUuidAPIRequest) GetDegAppKey() string {
	return r._degAppKey
}

// SetDegEventKey is DegEventKey Setter
// 活动后台配置appkey
func (r *TaobaoDegoperationGetInfoUuidAPIRequest) SetDegEventKey(_degEventKey string) error {
	r._degEventKey = _degEventKey
	r.Set("deg_event_key", _degEventKey)
	return nil
}

// GetDegEventKey DegEventKey Getter
func (r TaobaoDegoperationGetInfoUuidAPIRequest) GetDegEventKey() string {
	return r._degEventKey
}

// SetUuid is Uuid Setter
// 设备id
func (r *TaobaoDegoperationGetInfoUuidAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r TaobaoDegoperationGetInfoUuidAPIRequest) GetUuid() string {
	return r._uuid
}
