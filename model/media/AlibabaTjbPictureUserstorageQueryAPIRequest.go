package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatjbpictureuserstoragequeryAPIRequest 淘特图片空间用户容量查询 API请求
// alibaba.tjb.picture.userstorage.query
//
// 淘特图片空间用户容量查询
type AlibabatjbpictureuserstoragequeryAPIRequest struct {
	model.Params
}

// NewAlibabatjbpictureuserstoragequeryRequest 初始化AlibabatjbpictureuserstoragequeryAPIRequest对象
func NewAlibabatjbpictureuserstoragequeryRequest() *AlibabatjbpictureuserstoragequeryAPIRequest {
	return &AlibabatjbpictureuserstoragequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatjbpictureuserstoragequeryAPIRequest) GetApiMethodName() string {
	return "alibaba.tjb.picture.userstorage.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatjbpictureuserstoragequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatjbpictureuserstoragequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}
