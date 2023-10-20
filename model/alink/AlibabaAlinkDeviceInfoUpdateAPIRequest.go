package alink

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalinkdeviceinfoupdateAPIRequest 更新设备昵称等信息 API请求
// alibaba.alink.device.info.update
//
// 更新设备昵称等信息
type AlibabaalinkdeviceinfoupdateAPIRequest struct {
	model.Params
	// 设备id
	_uuid string
	// 设备昵称
	_nickName string
}

// NewAlibabaalinkdeviceinfoupdateRequest 初始化AlibabaalinkdeviceinfoupdateAPIRequest对象
func NewAlibabaalinkdeviceinfoupdateRequest() *AlibabaalinkdeviceinfoupdateAPIRequest {
	return &AlibabaalinkdeviceinfoupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalinkdeviceinfoupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alink.device.info.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalinkdeviceinfoupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalinkdeviceinfoupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// 设备id
func (r *AlibabaalinkdeviceinfoupdateAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaalinkdeviceinfoupdateAPIRequest) GetUuid() string {
	return r._uuid
}

// SetNickName is NickName Setter
// 设备昵称
func (r *AlibabaalinkdeviceinfoupdateAPIRequest) SetNickName(_nickName string) error {
	r._nickName = _nickName
	r.Set("nick_name", _nickName)
	return nil
}

// GetNickName NickName Getter
func (r AlibabaalinkdeviceinfoupdateAPIRequest) GetNickName() string {
	return r._nickName
}
