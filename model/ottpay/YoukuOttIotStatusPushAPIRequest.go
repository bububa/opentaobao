package ottpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YoukuottiotstatuspushAPIRequest iot设备状态变化通知接口 API请求
// youku.ott.iot.status.push
//
// ott iot设备状态通知
type YoukuottiotstatuspushAPIRequest struct {
	model.Params
	// 变更信息
	_changeInfo string
}

// NewYoukuottiotstatuspushRequest 初始化YoukuottiotstatuspushAPIRequest对象
func NewYoukuottiotstatuspushRequest() *YoukuottiotstatuspushAPIRequest {
	return &YoukuottiotstatuspushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuottiotstatuspushAPIRequest) GetApiMethodName() string {
	return "youku.ott.iot.status.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuottiotstatuspushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukuottiotstatuspushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChangeInfo is ChangeInfo Setter
// 变更信息
func (r *YoukuottiotstatuspushAPIRequest) SetChangeInfo(_changeInfo string) error {
	r._changeInfo = _changeInfo
	r.Set("change_info", _changeInfo)
	return nil
}

// GetChangeInfo ChangeInfo Getter
func (r YoukuottiotstatuspushAPIRequest) GetChangeInfo() string {
	return r._changeInfo
}
