package media

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTjbPictureUserstorageQueryAPIRequest 淘特图片空间用户容量查询 API请求
// alibaba.tjb.picture.userstorage.query
//
// 淘特图片空间用户容量查询
type AlibabaTjbPictureUserstorageQueryAPIRequest struct {
	model.Params
}

// NewAlibabaTjbPictureUserstorageQueryRequest 初始化AlibabaTjbPictureUserstorageQueryAPIRequest对象
func NewAlibabaTjbPictureUserstorageQueryRequest() *AlibabaTjbPictureUserstorageQueryAPIRequest {
	return &AlibabaTjbPictureUserstorageQueryAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTjbPictureUserstorageQueryAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTjbPictureUserstorageQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.tjb.picture.userstorage.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTjbPictureUserstorageQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTjbPictureUserstorageQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaTjbPictureUserstorageQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTjbPictureUserstorageQueryRequest()
	},
}

// GetAlibabaTjbPictureUserstorageQueryRequest 从 sync.Pool 获取 AlibabaTjbPictureUserstorageQueryAPIRequest
func GetAlibabaTjbPictureUserstorageQueryAPIRequest() *AlibabaTjbPictureUserstorageQueryAPIRequest {
	return poolAlibabaTjbPictureUserstorageQueryAPIRequest.Get().(*AlibabaTjbPictureUserstorageQueryAPIRequest)
}

// ReleaseAlibabaTjbPictureUserstorageQueryAPIRequest 将 AlibabaTjbPictureUserstorageQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaTjbPictureUserstorageQueryAPIRequest(v *AlibabaTjbPictureUserstorageQueryAPIRequest) {
	v.Reset()
	poolAlibabaTjbPictureUserstorageQueryAPIRequest.Put(v)
}
