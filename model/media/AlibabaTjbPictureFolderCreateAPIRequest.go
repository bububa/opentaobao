package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTjbPictureFolderCreateAPIRequest 淘特图片空间创建文件夹 API请求
// alibaba.tjb.picture.folder.create
//
// 淘特图片空间创建文件夹
type AlibabaTjbPictureFolderCreateAPIRequest struct {
	model.Params
	// 文件夹名称，最长50个字符，仅可输入中英文、数字及!@#%^&()-=_+.,~`特殊字符。
	_folderName string
	// 父文件夹id。如果是传到根文件夹，该值为0。
	_parentFolderId int64
}

// NewAlibabaTjbPictureFolderCreateRequest 初始化AlibabaTjbPictureFolderCreateAPIRequest对象
func NewAlibabaTjbPictureFolderCreateRequest() *AlibabaTjbPictureFolderCreateAPIRequest {
	return &AlibabaTjbPictureFolderCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTjbPictureFolderCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.tjb.picture.folder.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTjbPictureFolderCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetFolderName is FolderName Setter
// 文件夹名称，最长50个字符，仅可输入中英文、数字及!@#%^&amp;()-=_+.,~`特殊字符。
func (r *AlibabaTjbPictureFolderCreateAPIRequest) SetFolderName(_folderName string) error {
	r._folderName = _folderName
	r.Set("folder_name", _folderName)
	return nil
}

// GetFolderName FolderName Getter
func (r AlibabaTjbPictureFolderCreateAPIRequest) GetFolderName() string {
	return r._folderName
}

// SetParentFolderId is ParentFolderId Setter
// 父文件夹id。如果是传到根文件夹，该值为0。
func (r *AlibabaTjbPictureFolderCreateAPIRequest) SetParentFolderId(_parentFolderId int64) error {
	r._parentFolderId = _parentFolderId
	r.Set("parent_folder_id", _parentFolderId)
	return nil
}

// GetParentFolderId ParentFolderId Getter
func (r AlibabaTjbPictureFolderCreateAPIRequest) GetParentFolderId() int64 {
	return r._parentFolderId
}
