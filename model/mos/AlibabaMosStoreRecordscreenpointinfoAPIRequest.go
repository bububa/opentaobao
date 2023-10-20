package mos

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosStoreRecordscreenpointinfoAPIRequest 云屏埋点数据记录接口 API请求
// alibaba.mos.store.recordscreenpointinfo
//
// 记录云屏埋点数据
type AlibabaMosStoreRecordscreenpointinfoAPIRequest struct {
	model.Params
	// 云屏埋点信息
	_screenPointInfo string
}

// NewAlibabaMosStoreRecordscreenpointinfoRequest 初始化AlibabaMosStoreRecordscreenpointinfoAPIRequest对象
func NewAlibabaMosStoreRecordscreenpointinfoRequest() *AlibabaMosStoreRecordscreenpointinfoAPIRequest {
	return &AlibabaMosStoreRecordscreenpointinfoAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMosStoreRecordscreenpointinfoAPIRequest) Reset() {
	r._screenPointInfo = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosStoreRecordscreenpointinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.store.recordscreenpointinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosStoreRecordscreenpointinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMosStoreRecordscreenpointinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetScreenPointInfo is ScreenPointInfo Setter
// 云屏埋点信息
func (r *AlibabaMosStoreRecordscreenpointinfoAPIRequest) SetScreenPointInfo(_screenPointInfo string) error {
	r._screenPointInfo = _screenPointInfo
	r.Set("screen_point_info", _screenPointInfo)
	return nil
}

// GetScreenPointInfo ScreenPointInfo Getter
func (r AlibabaMosStoreRecordscreenpointinfoAPIRequest) GetScreenPointInfo() string {
	return r._screenPointInfo
}

var poolAlibabaMosStoreRecordscreenpointinfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMosStoreRecordscreenpointinfoRequest()
	},
}

// GetAlibabaMosStoreRecordscreenpointinfoRequest 从 sync.Pool 获取 AlibabaMosStoreRecordscreenpointinfoAPIRequest
func GetAlibabaMosStoreRecordscreenpointinfoAPIRequest() *AlibabaMosStoreRecordscreenpointinfoAPIRequest {
	return poolAlibabaMosStoreRecordscreenpointinfoAPIRequest.Get().(*AlibabaMosStoreRecordscreenpointinfoAPIRequest)
}

// ReleaseAlibabaMosStoreRecordscreenpointinfoAPIRequest 将 AlibabaMosStoreRecordscreenpointinfoAPIRequest 放入 sync.Pool
func ReleaseAlibabaMosStoreRecordscreenpointinfoAPIRequest(v *AlibabaMosStoreRecordscreenpointinfoAPIRequest) {
	v.Reset()
	poolAlibabaMosStoreRecordscreenpointinfoAPIRequest.Put(v)
}
