package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDiccontroltaskUpdateAPIRequest 停开服任务状态变更 API请求
// yunos.tvpubadmin.diccontroltask.update
//
// 停开服任务状态变更
type YunosTvpubadminDiccontroltaskUpdateAPIRequest struct {
	model.Params
	// 任务ID
	_id int64
	// 任务状态
	_status int64
	// 牌照方
	_license int64
}

// NewYunosTvpubadminDiccontroltaskUpdateRequest 初始化YunosTvpubadminDiccontroltaskUpdateAPIRequest对象
func NewYunosTvpubadminDiccontroltaskUpdateRequest() *YunosTvpubadminDiccontroltaskUpdateAPIRequest {
	return &YunosTvpubadminDiccontroltaskUpdateAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminDiccontroltaskUpdateAPIRequest) Reset() {
	r._id = 0
	r._status = 0
	r._license = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDiccontroltaskUpdateAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.diccontroltask.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDiccontroltaskUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminDiccontroltaskUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 任务ID
func (r *YunosTvpubadminDiccontroltaskUpdateAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r YunosTvpubadminDiccontroltaskUpdateAPIRequest) GetId() int64 {
	return r._id
}

// SetStatus is Status Setter
// 任务状态
func (r *YunosTvpubadminDiccontroltaskUpdateAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r YunosTvpubadminDiccontroltaskUpdateAPIRequest) GetStatus() int64 {
	return r._status
}

// SetLicense is License Setter
// 牌照方
func (r *YunosTvpubadminDiccontroltaskUpdateAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunosTvpubadminDiccontroltaskUpdateAPIRequest) GetLicense() int64 {
	return r._license
}

var poolYunosTvpubadminDiccontroltaskUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminDiccontroltaskUpdateRequest()
	},
}

// GetYunosTvpubadminDiccontroltaskUpdateRequest 从 sync.Pool 获取 YunosTvpubadminDiccontroltaskUpdateAPIRequest
func GetYunosTvpubadminDiccontroltaskUpdateAPIRequest() *YunosTvpubadminDiccontroltaskUpdateAPIRequest {
	return poolYunosTvpubadminDiccontroltaskUpdateAPIRequest.Get().(*YunosTvpubadminDiccontroltaskUpdateAPIRequest)
}

// ReleaseYunosTvpubadminDiccontroltaskUpdateAPIRequest 将 YunosTvpubadminDiccontroltaskUpdateAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminDiccontroltaskUpdateAPIRequest(v *YunosTvpubadminDiccontroltaskUpdateAPIRequest) {
	v.Reset()
	poolYunosTvpubadminDiccontroltaskUpdateAPIRequest.Put(v)
}
