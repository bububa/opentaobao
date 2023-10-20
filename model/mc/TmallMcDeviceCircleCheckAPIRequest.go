package mc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallmcdevicecirclecheckAPIRequest 云码设备圈选情况查询 API请求
// tmall.mc.device.circle.check
//
// 云码设备圈选情况查询
type TmallmcdevicecirclecheckAPIRequest struct {
	model.Params
	// 外部设备编码
	_outerCode string
	// 渠道编码
	_channelId string
}

// NewTmallmcdevicecirclecheckRequest 初始化TmallmcdevicecirclecheckAPIRequest对象
func NewTmallmcdevicecirclecheckRequest() *TmallmcdevicecirclecheckAPIRequest {
	return &TmallmcdevicecirclecheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallmcdevicecirclecheckAPIRequest) GetApiMethodName() string {
	return "tmall.mc.device.circle.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallmcdevicecirclecheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallmcdevicecirclecheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterCode is OuterCode Setter
// 外部设备编码
func (r *TmallmcdevicecirclecheckAPIRequest) SetOuterCode(_outerCode string) error {
	r._outerCode = _outerCode
	r.Set("outer_code", _outerCode)
	return nil
}

// GetOuterCode OuterCode Getter
func (r TmallmcdevicecirclecheckAPIRequest) GetOuterCode() string {
	return r._outerCode
}

// SetChannelId is ChannelId Setter
// 渠道编码
func (r *TmallmcdevicecirclecheckAPIRequest) SetChannelId(_channelId string) error {
	r._channelId = _channelId
	r.Set("channel_id", _channelId)
	return nil
}

// GetChannelId ChannelId Getter
func (r TmallmcdevicecirclecheckAPIRequest) GetChannelId() string {
	return r._channelId
}
