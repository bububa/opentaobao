package wenyuvideo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YoukuwenyuvideopersiongetAPIRequest 根据优酷人物ID获取人物详情页，包含相关影视和相关人物 API请求
// youku.wenyuvideo.persion.get
//
// 根据优酷人物ID获取人物详情页，包含相关影视和相关人物
type YoukuwenyuvideopersiongetAPIRequest struct {
	model.Params
	// 设备信息
	_systemInfo string
	// 人物ID
	_personId int64
}

// NewYoukuwenyuvideopersiongetRequest 初始化YoukuwenyuvideopersiongetAPIRequest对象
func NewYoukuwenyuvideopersiongetRequest() *YoukuwenyuvideopersiongetAPIRequest {
	return &YoukuwenyuvideopersiongetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuwenyuvideopersiongetAPIRequest) GetApiMethodName() string {
	return "youku.wenyuvideo.persion.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuwenyuvideopersiongetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukuwenyuvideopersiongetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSystemInfo is SystemInfo Setter
// 设备信息
func (r *YoukuwenyuvideopersiongetAPIRequest) SetSystemInfo(_systemInfo string) error {
	r._systemInfo = _systemInfo
	r.Set("system_info", _systemInfo)
	return nil
}

// GetSystemInfo SystemInfo Getter
func (r YoukuwenyuvideopersiongetAPIRequest) GetSystemInfo() string {
	return r._systemInfo
}

// SetPersonId is PersonId Setter
// 人物ID
func (r *YoukuwenyuvideopersiongetAPIRequest) SetPersonId(_personId int64) error {
	r._personId = _personId
	r.Set("person_id", _personId)
	return nil
}

// GetPersonId PersonId Getter
func (r YoukuwenyuvideopersiongetAPIRequest) GetPersonId() int64 {
	return r._personId
}
