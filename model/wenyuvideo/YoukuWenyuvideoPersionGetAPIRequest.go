package wenyuvideo

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuWenyuvideoPersionGetAPIRequest 根据优酷人物ID获取人物详情页，包含相关影视和相关人物 API请求
// youku.wenyuvideo.persion.get
//
// 根据优酷人物ID获取人物详情页，包含相关影视和相关人物
type YoukuWenyuvideoPersionGetAPIRequest struct {
	model.Params
	// 设备信息
	_systemInfo string
	// 人物ID
	_personId int64
}

// NewYoukuWenyuvideoPersionGetRequest 初始化YoukuWenyuvideoPersionGetAPIRequest对象
func NewYoukuWenyuvideoPersionGetRequest() *YoukuWenyuvideoPersionGetAPIRequest {
	return &YoukuWenyuvideoPersionGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YoukuWenyuvideoPersionGetAPIRequest) Reset() {
	r._systemInfo = ""
	r._personId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuWenyuvideoPersionGetAPIRequest) GetApiMethodName() string {
	return "youku.wenyuvideo.persion.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuWenyuvideoPersionGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukuWenyuvideoPersionGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSystemInfo is SystemInfo Setter
// 设备信息
func (r *YoukuWenyuvideoPersionGetAPIRequest) SetSystemInfo(_systemInfo string) error {
	r._systemInfo = _systemInfo
	r.Set("system_info", _systemInfo)
	return nil
}

// GetSystemInfo SystemInfo Getter
func (r YoukuWenyuvideoPersionGetAPIRequest) GetSystemInfo() string {
	return r._systemInfo
}

// SetPersonId is PersonId Setter
// 人物ID
func (r *YoukuWenyuvideoPersionGetAPIRequest) SetPersonId(_personId int64) error {
	r._personId = _personId
	r.Set("person_id", _personId)
	return nil
}

// GetPersonId PersonId Getter
func (r YoukuWenyuvideoPersionGetAPIRequest) GetPersonId() int64 {
	return r._personId
}

var poolYoukuWenyuvideoPersionGetAPIRequest = sync.Pool{
	New: func() any {
		return NewYoukuWenyuvideoPersionGetRequest()
	},
}

// GetYoukuWenyuvideoPersionGetRequest 从 sync.Pool 获取 YoukuWenyuvideoPersionGetAPIRequest
func GetYoukuWenyuvideoPersionGetAPIRequest() *YoukuWenyuvideoPersionGetAPIRequest {
	return poolYoukuWenyuvideoPersionGetAPIRequest.Get().(*YoukuWenyuvideoPersionGetAPIRequest)
}

// ReleaseYoukuWenyuvideoPersionGetAPIRequest 将 YoukuWenyuvideoPersionGetAPIRequest 放入 sync.Pool
func ReleaseYoukuWenyuvideoPersionGetAPIRequest(v *YoukuWenyuvideoPersionGetAPIRequest) {
	v.Reset()
	poolYoukuWenyuvideoPersionGetAPIRequest.Put(v)
}
