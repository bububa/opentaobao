package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTjbPictureFolderQueryAPIRequest 淘特图片空间用户文件夹查询 API请求
// alibaba.tjb.picture.folder.query
//
// 淘特图片空间用户文件夹查询，返回用户所有的文件夹。
type AlibabaTjbPictureFolderQueryAPIRequest struct {
	model.Params
}

// NewAlibabaTjbPictureFolderQueryRequest 初始化AlibabaTjbPictureFolderQueryAPIRequest对象
func NewAlibabaTjbPictureFolderQueryRequest() *AlibabaTjbPictureFolderQueryAPIRequest {
	return &AlibabaTjbPictureFolderQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTjbPictureFolderQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.tjb.picture.folder.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTjbPictureFolderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTjbPictureFolderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}
