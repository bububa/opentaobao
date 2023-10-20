package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceTvidAPIRequest 查询终端信息 API请求
// yunos.tvpubadmin.device.tvid
//
// 通过UUID查询终端信息
type YunosTvpubadminDeviceTvidAPIRequest struct {
	model.Params
	// 设备的UUID
	_uuid string
}

// NewYunosTvpubadminDeviceTvidRequest 初始化YunosTvpubadminDeviceTvidAPIRequest对象
func NewYunosTvpubadminDeviceTvidRequest() *YunosTvpubadminDeviceTvidAPIRequest {
	return &YunosTvpubadminDeviceTvidAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminDeviceTvidAPIRequest) Reset() {
	r._uuid = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceTvidAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.tvid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceTvidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminDeviceTvidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// 设备的UUID
func (r *YunosTvpubadminDeviceTvidAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r YunosTvpubadminDeviceTvidAPIRequest) GetUuid() string {
	return r._uuid
}

var poolYunosTvpubadminDeviceTvidAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminDeviceTvidRequest()
	},
}

// GetYunosTvpubadminDeviceTvidRequest 从 sync.Pool 获取 YunosTvpubadminDeviceTvidAPIRequest
func GetYunosTvpubadminDeviceTvidAPIRequest() *YunosTvpubadminDeviceTvidAPIRequest {
	return poolYunosTvpubadminDeviceTvidAPIRequest.Get().(*YunosTvpubadminDeviceTvidAPIRequest)
}

// ReleaseYunosTvpubadminDeviceTvidAPIRequest 将 YunosTvpubadminDeviceTvidAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminDeviceTvidAPIRequest(v *YunosTvpubadminDeviceTvidAPIRequest) {
	v.Reset()
	poolYunosTvpubadminDeviceTvidAPIRequest.Put(v)
}
