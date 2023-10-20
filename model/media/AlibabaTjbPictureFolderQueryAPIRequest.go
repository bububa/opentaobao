package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatjbpicturefolderqueryAPIRequest 淘特图片空间用户文件夹查询 API请求
// alibaba.tjb.picture.folder.query
//
// 淘特图片空间用户文件夹查询，返回用户所有的文件夹。
type AlibabatjbpicturefolderqueryAPIRequest struct {
	model.Params
}

// NewAlibabatjbpicturefolderqueryRequest 初始化AlibabatjbpicturefolderqueryAPIRequest对象
func NewAlibabatjbpicturefolderqueryRequest() *AlibabatjbpicturefolderqueryAPIRequest {
	return &AlibabatjbpicturefolderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatjbpicturefolderqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.tjb.picture.folder.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatjbpicturefolderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatjbpicturefolderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}
