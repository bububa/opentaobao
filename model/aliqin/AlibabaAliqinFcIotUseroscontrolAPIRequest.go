package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinfciotuseroscontrolAPIRequest 物联卡用户管理停开机功能 API请求
// alibaba.aliqin.fc.iot.useroscontrol
//
// 物联网针对用户级管理停开支持
type AlibabaaliqinfciotuseroscontrolAPIRequest struct {
	model.Params
	// 物联卡的iccid
	_iccid string
	// 用户停开的操作类型：MANAGE_RESUME、MANAGE_STOP
	_action string
}

// NewAlibabaaliqinfciotuseroscontrolRequest 初始化AlibabaaliqinfciotuseroscontrolAPIRequest对象
func NewAlibabaaliqinfciotuseroscontrolRequest() *AlibabaaliqinfciotuseroscontrolAPIRequest {
	return &AlibabaaliqinfciotuseroscontrolAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliqinfciotuseroscontrolAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.iot.useroscontrol"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliqinfciotuseroscontrolAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliqinfciotuseroscontrolAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIccid is Iccid Setter
// 物联卡的iccid
func (r *AlibabaaliqinfciotuseroscontrolAPIRequest) SetIccid(_iccid string) error {
	r._iccid = _iccid
	r.Set("iccid", _iccid)
	return nil
}

// GetIccid Iccid Getter
func (r AlibabaaliqinfciotuseroscontrolAPIRequest) GetIccid() string {
	return r._iccid
}

// SetAction is Action Setter
// 用户停开的操作类型：MANAGE_RESUME、MANAGE_STOP
func (r *AlibabaaliqinfciotuseroscontrolAPIRequest) SetAction(_action string) error {
	r._action = _action
	r.Set("action", _action)
	return nil
}

// GetAction Action Getter
func (r AlibabaaliqinfciotuseroscontrolAPIRequest) GetAction() string {
	return r._action
}
