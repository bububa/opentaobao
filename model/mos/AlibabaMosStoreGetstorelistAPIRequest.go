package mos

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosStoreGetstorelistAPIRequest 根据屏编号获取专柜集 API请求
// alibaba.mos.store.getstorelist
//
// 根据屏编号获取专柜集
type AlibabaMosStoreGetstorelistAPIRequest struct {
	model.Params
	// 屏编号
	_screenNo string
}

// NewAlibabaMosStoreGetstorelistRequest 初始化AlibabaMosStoreGetstorelistAPIRequest对象
func NewAlibabaMosStoreGetstorelistRequest() *AlibabaMosStoreGetstorelistAPIRequest {
	return &AlibabaMosStoreGetstorelistAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMosStoreGetstorelistAPIRequest) Reset() {
	r._screenNo = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosStoreGetstorelistAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.store.getstorelist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosStoreGetstorelistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMosStoreGetstorelistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetScreenNo is ScreenNo Setter
// 屏编号
func (r *AlibabaMosStoreGetstorelistAPIRequest) SetScreenNo(_screenNo string) error {
	r._screenNo = _screenNo
	r.Set("screen_no", _screenNo)
	return nil
}

// GetScreenNo ScreenNo Getter
func (r AlibabaMosStoreGetstorelistAPIRequest) GetScreenNo() string {
	return r._screenNo
}

var poolAlibabaMosStoreGetstorelistAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMosStoreGetstorelistRequest()
	},
}

// GetAlibabaMosStoreGetstorelistRequest 从 sync.Pool 获取 AlibabaMosStoreGetstorelistAPIRequest
func GetAlibabaMosStoreGetstorelistAPIRequest() *AlibabaMosStoreGetstorelistAPIRequest {
	return poolAlibabaMosStoreGetstorelistAPIRequest.Get().(*AlibabaMosStoreGetstorelistAPIRequest)
}

// ReleaseAlibabaMosStoreGetstorelistAPIRequest 将 AlibabaMosStoreGetstorelistAPIRequest 放入 sync.Pool
func ReleaseAlibabaMosStoreGetstorelistAPIRequest(v *AlibabaMosStoreGetstorelistAPIRequest) {
	v.Reset()
	poolAlibabaMosStoreGetstorelistAPIRequest.Put(v)
}
