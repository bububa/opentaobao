package mc

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallMcDeviceCircleCheckAPIRequest) Reset() {
	r._outerCode = ""
	r._channelId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallMcDeviceCircleCheckAPIRequest) GetApiMethodName() string {
	return "tmall.mc.device.circle.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallMcDeviceCircleCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallMcDeviceCircleCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterCode is OuterCode Setter
// 外部设备编码
func (r *TmallMcDeviceCircleCheckAPIRequest) SetOuterCode(_outerCode string) error {
	r._outerCode = _outerCode
	r.Set("outer_code", _outerCode)
	return nil
}

// GetOuterCode OuterCode Getter
func (r TmallMcDeviceCircleCheckAPIRequest) GetOuterCode() string {
	return r._outerCode
}

// SetChannelId is ChannelId Setter
// 渠道编码
func (r *TmallMcDeviceCircleCheckAPIRequest) SetChannelId(_channelId string) error {
	r._channelId = _channelId
	r.Set("channel_id", _channelId)
	return nil
}

// GetChannelId ChannelId Getter
func (r TmallMcDeviceCircleCheckAPIRequest) GetChannelId() string {
	return r._channelId
}

var poolTmallMcDeviceCircleCheckAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallMcDeviceCircleCheckRequest()
	},
}

// GetTmallMcDeviceCircleCheckRequest 从 sync.Pool 获取 TmallMcDeviceCircleCheckAPIRequest
func GetTmallMcDeviceCircleCheckAPIRequest() *TmallMcDeviceCircleCheckAPIRequest {
	return poolTmallMcDeviceCircleCheckAPIRequest.Get().(*TmallMcDeviceCircleCheckAPIRequest)
}

// ReleaseTmallMcDeviceCircleCheckAPIRequest 将 TmallMcDeviceCircleCheckAPIRequest 放入 sync.Pool
func ReleaseTmallMcDeviceCircleCheckAPIRequest(v *TmallMcDeviceCircleCheckAPIRequest) {
	v.Reset()
	poolTmallMcDeviceCircleCheckAPIRequest.Put(v)
}
