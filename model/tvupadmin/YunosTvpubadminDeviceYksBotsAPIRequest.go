package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceYksBotsAPIRequest 获取设备列表 API请求
// yunos.tvpubadmin.device.yks.bots
//
// 获取设备列表
type YunosTvpubadminDeviceYksBotsAPIRequest struct {
	model.Params
}

// NewYunosTvpubadminDeviceYksBotsRequest 初始化YunosTvpubadminDeviceYksBotsAPIRequest对象
func NewYunosTvpubadminDeviceYksBotsRequest() *YunosTvpubadminDeviceYksBotsAPIRequest {
	return &YunosTvpubadminDeviceYksBotsAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminDeviceYksBotsAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceYksBotsAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.yks.bots"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceYksBotsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminDeviceYksBotsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolYunosTvpubadminDeviceYksBotsAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminDeviceYksBotsRequest()
	},
}

// GetYunosTvpubadminDeviceYksBotsRequest 从 sync.Pool 获取 YunosTvpubadminDeviceYksBotsAPIRequest
func GetYunosTvpubadminDeviceYksBotsAPIRequest() *YunosTvpubadminDeviceYksBotsAPIRequest {
	return poolYunosTvpubadminDeviceYksBotsAPIRequest.Get().(*YunosTvpubadminDeviceYksBotsAPIRequest)
}

// ReleaseYunosTvpubadminDeviceYksBotsAPIRequest 将 YunosTvpubadminDeviceYksBotsAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminDeviceYksBotsAPIRequest(v *YunosTvpubadminDeviceYksBotsAPIRequest) {
	v.Reset()
	poolYunosTvpubadminDeviceYksBotsAPIRequest.Put(v)
}
