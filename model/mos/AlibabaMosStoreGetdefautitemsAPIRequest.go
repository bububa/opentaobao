package mos

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosStoreGetdefautitemsAPIRequest 获取默认状态下商品列表 API请求
// alibaba.mos.store.getdefautitems
//
// 获取默认状态下商品列表
type AlibabaMosStoreGetdefautitemsAPIRequest struct {
	model.Params
	// 屏编号
	_screenNo string
	// 分页查询开始index
	_start int64
	// 分页查询每页记录数
	_limitCount int64
}

// NewAlibabaMosStoreGetdefautitemsRequest 初始化AlibabaMosStoreGetdefautitemsAPIRequest对象
func NewAlibabaMosStoreGetdefautitemsRequest() *AlibabaMosStoreGetdefautitemsAPIRequest {
	return &AlibabaMosStoreGetdefautitemsAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMosStoreGetdefautitemsAPIRequest) Reset() {
	r._screenNo = ""
	r._start = 0
	r._limitCount = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosStoreGetdefautitemsAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.store.getdefautitems"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosStoreGetdefautitemsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMosStoreGetdefautitemsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetScreenNo is ScreenNo Setter
// 屏编号
func (r *AlibabaMosStoreGetdefautitemsAPIRequest) SetScreenNo(_screenNo string) error {
	r._screenNo = _screenNo
	r.Set("screen_no", _screenNo)
	return nil
}

// GetScreenNo ScreenNo Getter
func (r AlibabaMosStoreGetdefautitemsAPIRequest) GetScreenNo() string {
	return r._screenNo
}

// SetStart is Start Setter
// 分页查询开始index
func (r *AlibabaMosStoreGetdefautitemsAPIRequest) SetStart(_start int64) error {
	r._start = _start
	r.Set("start", _start)
	return nil
}

// GetStart Start Getter
func (r AlibabaMosStoreGetdefautitemsAPIRequest) GetStart() int64 {
	return r._start
}

// SetLimitCount is LimitCount Setter
// 分页查询每页记录数
func (r *AlibabaMosStoreGetdefautitemsAPIRequest) SetLimitCount(_limitCount int64) error {
	r._limitCount = _limitCount
	r.Set("limit_count", _limitCount)
	return nil
}

// GetLimitCount LimitCount Getter
func (r AlibabaMosStoreGetdefautitemsAPIRequest) GetLimitCount() int64 {
	return r._limitCount
}

var poolAlibabaMosStoreGetdefautitemsAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMosStoreGetdefautitemsRequest()
	},
}

// GetAlibabaMosStoreGetdefautitemsRequest 从 sync.Pool 获取 AlibabaMosStoreGetdefautitemsAPIRequest
func GetAlibabaMosStoreGetdefautitemsAPIRequest() *AlibabaMosStoreGetdefautitemsAPIRequest {
	return poolAlibabaMosStoreGetdefautitemsAPIRequest.Get().(*AlibabaMosStoreGetdefautitemsAPIRequest)
}

// ReleaseAlibabaMosStoreGetdefautitemsAPIRequest 将 AlibabaMosStoreGetdefautitemsAPIRequest 放入 sync.Pool
func ReleaseAlibabaMosStoreGetdefautitemsAPIRequest(v *AlibabaMosStoreGetdefautitemsAPIRequest) {
	v.Reset()
	poolAlibabaMosStoreGetdefautitemsAPIRequest.Put(v)
}
