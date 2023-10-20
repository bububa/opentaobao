package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosOsupdateVersionstatusUpdateAPIRequest 更新应用升级状态 API请求
// yunos.osupdate.versionstatus.update
//
// 更新应用升级状态
type YunosOsupdateVersionstatusUpdateAPIRequest struct {
	model.Params
	// 状态值
	_status string
	// 升级任务ID
	_id int64
}

// NewYunosOsupdateVersionstatusUpdateRequest 初始化YunosOsupdateVersionstatusUpdateAPIRequest对象
func NewYunosOsupdateVersionstatusUpdateRequest() *YunosOsupdateVersionstatusUpdateAPIRequest {
	return &YunosOsupdateVersionstatusUpdateAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosOsupdateVersionstatusUpdateAPIRequest) Reset() {
	r._status = ""
	r._id = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosOsupdateVersionstatusUpdateAPIRequest) GetApiMethodName() string {
	return "yunos.osupdate.versionstatus.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosOsupdateVersionstatusUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosOsupdateVersionstatusUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStatus is Status Setter
// 状态值
func (r *YunosOsupdateVersionstatusUpdateAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r YunosOsupdateVersionstatusUpdateAPIRequest) GetStatus() string {
	return r._status
}

// SetId is Id Setter
// 升级任务ID
func (r *YunosOsupdateVersionstatusUpdateAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r YunosOsupdateVersionstatusUpdateAPIRequest) GetId() int64 {
	return r._id
}

var poolYunosOsupdateVersionstatusUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosOsupdateVersionstatusUpdateRequest()
	},
}

// GetYunosOsupdateVersionstatusUpdateRequest 从 sync.Pool 获取 YunosOsupdateVersionstatusUpdateAPIRequest
func GetYunosOsupdateVersionstatusUpdateAPIRequest() *YunosOsupdateVersionstatusUpdateAPIRequest {
	return poolYunosOsupdateVersionstatusUpdateAPIRequest.Get().(*YunosOsupdateVersionstatusUpdateAPIRequest)
}

// ReleaseYunosOsupdateVersionstatusUpdateAPIRequest 将 YunosOsupdateVersionstatusUpdateAPIRequest 放入 sync.Pool
func ReleaseYunosOsupdateVersionstatusUpdateAPIRequest(v *YunosOsupdateVersionstatusUpdateAPIRequest) {
	v.Reset()
	poolYunosOsupdateVersionstatusUpdateAPIRequest.Put(v)
}
