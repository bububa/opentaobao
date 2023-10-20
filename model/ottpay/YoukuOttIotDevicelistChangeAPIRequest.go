package ottpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YoukuottiotdevicelistchangeAPIRequest iot设备列表变化接口 API请求
// youku.ott.iot.devicelist.change
//
// iot设备列表变化接口
type YoukuottiotdevicelistchangeAPIRequest struct {
	model.Params
	// 变更信息
	_changeInfo string
}

// NewYoukuottiotdevicelistchangeRequest 初始化YoukuottiotdevicelistchangeAPIRequest对象
func NewYoukuottiotdevicelistchangeRequest() *YoukuottiotdevicelistchangeAPIRequest {
	return &YoukuottiotdevicelistchangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuottiotdevicelistchangeAPIRequest) GetApiMethodName() string {
	return "youku.ott.iot.devicelist.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuottiotdevicelistchangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukuottiotdevicelistchangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChangeInfo is ChangeInfo Setter
// 变更信息
func (r *YoukuottiotdevicelistchangeAPIRequest) SetChangeInfo(_changeInfo string) error {
	r._changeInfo = _changeInfo
	r.Set("change_info", _changeInfo)
	return nil
}

// GetChangeInfo ChangeInfo Getter
func (r YoukuottiotdevicelistchangeAPIRequest) GetChangeInfo() string {
	return r._changeInfo
}
