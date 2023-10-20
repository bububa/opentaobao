package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceApkinfoAPIRequest 获取停开服apk信息 API请求
// yunos.tvpubadmin.device.apkinfo
//
// 获取停开服apk信息
type YunosTvpubadminDeviceApkinfoAPIRequest struct {
	model.Params
	// apkid
	_id int64
}

// NewYunosTvpubadminDeviceApkinfoRequest 初始化YunosTvpubadminDeviceApkinfoAPIRequest对象
func NewYunosTvpubadminDeviceApkinfoRequest() *YunosTvpubadminDeviceApkinfoAPIRequest {
	return &YunosTvpubadminDeviceApkinfoAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminDeviceApkinfoAPIRequest) Reset() {
	r._id = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceApkinfoAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.apkinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceApkinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminDeviceApkinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// apkid
func (r *YunosTvpubadminDeviceApkinfoAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r YunosTvpubadminDeviceApkinfoAPIRequest) GetId() int64 {
	return r._id
}

var poolYunosTvpubadminDeviceApkinfoAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminDeviceApkinfoRequest()
	},
}

// GetYunosTvpubadminDeviceApkinfoRequest 从 sync.Pool 获取 YunosTvpubadminDeviceApkinfoAPIRequest
func GetYunosTvpubadminDeviceApkinfoAPIRequest() *YunosTvpubadminDeviceApkinfoAPIRequest {
	return poolYunosTvpubadminDeviceApkinfoAPIRequest.Get().(*YunosTvpubadminDeviceApkinfoAPIRequest)
}

// ReleaseYunosTvpubadminDeviceApkinfoAPIRequest 将 YunosTvpubadminDeviceApkinfoAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminDeviceApkinfoAPIRequest(v *YunosTvpubadminDeviceApkinfoAPIRequest) {
	v.Reset()
	poolYunosTvpubadminDeviceApkinfoAPIRequest.Put(v)
}
