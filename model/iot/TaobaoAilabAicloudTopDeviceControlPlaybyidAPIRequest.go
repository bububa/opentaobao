package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopdevicecontrolplaybyidAPIRequest 通过id播放歌曲 API请求
// taobao.ailab.aicloud.top.device.control.playbyid
//
// 通过id播放歌曲
type TaobaoailabaicloudtopdevicecontrolplaybyidAPIRequest struct {
	model.Params
	// 设备id
	_param1 string
	// 音频id
	_param2 string
	// 音频来源
	_param3 string
	// 音频类型，如果没有音频类型默认填children_song
	_param4 string
	// 用户信息
	_param0 *OpenBaseInfo
}

// NewTaobaoailabaicloudtopdevicecontrolplaybyidRequest 初始化TaobaoailabaicloudtopdevicecontrolplaybyidAPIRequest对象
func NewTaobaoailabaicloudtopdevicecontrolplaybyidRequest() *TaobaoailabaicloudtopdevicecontrolplaybyidAPIRequest {
	return &TaobaoailabaicloudtopdevicecontrolplaybyidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtopdevicecontrolplaybyidAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.control.playbyid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtopdevicecontrolplaybyidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtopdevicecontrolplaybyidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam1 is Param1 Setter
// 设备id
func (r *TaobaoailabaicloudtopdevicecontrolplaybyidAPIRequest) SetParam1(_param1 string) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r TaobaoailabaicloudtopdevicecontrolplaybyidAPIRequest) GetParam1() string {
	return r._param1
}

// SetParam2 is Param2 Setter
// 音频id
func (r *TaobaoailabaicloudtopdevicecontrolplaybyidAPIRequest) SetParam2(_param2 string) error {
	r._param2 = _param2
	r.Set("param2", _param2)
	return nil
}

// GetParam2 Param2 Getter
func (r TaobaoailabaicloudtopdevicecontrolplaybyidAPIRequest) GetParam2() string {
	return r._param2
}

// SetParam3 is Param3 Setter
// 音频来源
func (r *TaobaoailabaicloudtopdevicecontrolplaybyidAPIRequest) SetParam3(_param3 string) error {
	r._param3 = _param3
	r.Set("param3", _param3)
	return nil
}

// GetParam3 Param3 Getter
func (r TaobaoailabaicloudtopdevicecontrolplaybyidAPIRequest) GetParam3() string {
	return r._param3
}

// SetParam4 is Param4 Setter
// 音频类型，如果没有音频类型默认填children_song
func (r *TaobaoailabaicloudtopdevicecontrolplaybyidAPIRequest) SetParam4(_param4 string) error {
	r._param4 = _param4
	r.Set("param4", _param4)
	return nil
}

// GetParam4 Param4 Getter
func (r TaobaoailabaicloudtopdevicecontrolplaybyidAPIRequest) GetParam4() string {
	return r._param4
}

// SetParam0 is Param0 Setter
// 用户信息
func (r *TaobaoailabaicloudtopdevicecontrolplaybyidAPIRequest) SetParam0(_param0 *OpenBaseInfo) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoailabaicloudtopdevicecontrolplaybyidAPIRequest) GetParam0() *OpenBaseInfo {
	return r._param0
}
