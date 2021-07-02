package mc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallMcDeviceCircleCheckAPIRequest 云码设备圈选情况查询 API请求
// tmall.mc.device.circle.check
//
// 云码设备圈选情况查询
type TmallMcDeviceCircleCheckAPIRequest struct {
	model.Params
	// 外部设备编码
	_outerCode string
	// 渠道编码
	_channelId string
}

// NewTmallMcDeviceCircleCheckRequest 初始化TmallMcDeviceCircleCheckAPIRequest对象
func NewTmallMcDeviceCircleCheckRequest() *TmallMcDeviceCircleCheckAPIRequest {
	return &TmallMcDeviceCircleCheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallMcDeviceCircleCheckAPIRequest) GetApiMethodName() string {
	return "tmall.mc.device.circle.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallMcDeviceCircleCheckAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OuterCode Setter
// 外部设备编码
func (r *TmallMcDeviceCircleCheckAPIRequest) SetOuterCode(_outerCode string) error {
	r._outerCode = _outerCode
	r.Set("outer_code", _outerCode)
	return nil
}

// Get OuterCode Getter
func (r TmallMcDeviceCircleCheckAPIRequest) GetOuterCode() string {
	return r._outerCode
}

// Set is ChannelId Setter
// 渠道编码
func (r *TmallMcDeviceCircleCheckAPIRequest) SetChannelId(_channelId string) error {
	r._channelId = _channelId
	r.Set("channel_id", _channelId)
	return nil
}

// Get ChannelId Getter
func (r TmallMcDeviceCircleCheckAPIRequest) GetChannelId() string {
	return r._channelId
}
