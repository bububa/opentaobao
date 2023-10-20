package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentDeviceGetterminaltypemapAPIRequest 获取终端类型表 API请求
// yunos.tvpubadmin.content.device.getterminaltypemap
//
// 获取终端类型表
type YunosTvpubadminContentDeviceGetterminaltypemapAPIRequest struct {
	model.Params
}

// NewYunosTvpubadminContentDeviceGetterminaltypemapRequest 初始化YunosTvpubadminContentDeviceGetterminaltypemapAPIRequest对象
func NewYunosTvpubadminContentDeviceGetterminaltypemapRequest() *YunosTvpubadminContentDeviceGetterminaltypemapAPIRequest {
	return &YunosTvpubadminContentDeviceGetterminaltypemapAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminContentDeviceGetterminaltypemapAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentDeviceGetterminaltypemapAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.device.getterminaltypemap"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentDeviceGetterminaltypemapAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminContentDeviceGetterminaltypemapAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolYunosTvpubadminContentDeviceGetterminaltypemapAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminContentDeviceGetterminaltypemapRequest()
	},
}

// GetYunosTvpubadminContentDeviceGetterminaltypemapRequest 从 sync.Pool 获取 YunosTvpubadminContentDeviceGetterminaltypemapAPIRequest
func GetYunosTvpubadminContentDeviceGetterminaltypemapAPIRequest() *YunosTvpubadminContentDeviceGetterminaltypemapAPIRequest {
	return poolYunosTvpubadminContentDeviceGetterminaltypemapAPIRequest.Get().(*YunosTvpubadminContentDeviceGetterminaltypemapAPIRequest)
}

// ReleaseYunosTvpubadminContentDeviceGetterminaltypemapAPIRequest 将 YunosTvpubadminContentDeviceGetterminaltypemapAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminContentDeviceGetterminaltypemapAPIRequest(v *YunosTvpubadminContentDeviceGetterminaltypemapAPIRequest) {
	v.Reset()
	poolYunosTvpubadminContentDeviceGetterminaltypemapAPIRequest.Put(v)
}
