package media

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTjbPictureUserstorageQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.tjb.picture.userstorage.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTjbPictureUserstorageQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
