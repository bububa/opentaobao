package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentAdvertGettypesAPIRequest 获取广告位类型 API请求
// yunos.tvpubadmin.content.advert.gettypes
//
// 获取广告位类型
type YunosTvpubadminContentAdvertGettypesAPIRequest struct {
	model.Params
}

// NewYunosTvpubadminContentAdvertGettypesRequest 初始化YunosTvpubadminContentAdvertGettypesAPIRequest对象
func NewYunosTvpubadminContentAdvertGettypesRequest() *YunosTvpubadminContentAdvertGettypesAPIRequest {
	return &YunosTvpubadminContentAdvertGettypesAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminContentAdvertGettypesAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentAdvertGettypesAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.advert.gettypes"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentAdvertGettypesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminContentAdvertGettypesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolYunosTvpubadminContentAdvertGettypesAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminContentAdvertGettypesRequest()
	},
}

// GetYunosTvpubadminContentAdvertGettypesRequest 从 sync.Pool 获取 YunosTvpubadminContentAdvertGettypesAPIRequest
func GetYunosTvpubadminContentAdvertGettypesAPIRequest() *YunosTvpubadminContentAdvertGettypesAPIRequest {
	return poolYunosTvpubadminContentAdvertGettypesAPIRequest.Get().(*YunosTvpubadminContentAdvertGettypesAPIRequest)
}

// ReleaseYunosTvpubadminContentAdvertGettypesAPIRequest 将 YunosTvpubadminContentAdvertGettypesAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminContentAdvertGettypesAPIRequest(v *YunosTvpubadminContentAdvertGettypesAPIRequest) {
	v.Reset()
	poolYunosTvpubadminContentAdvertGettypesAPIRequest.Put(v)
}
