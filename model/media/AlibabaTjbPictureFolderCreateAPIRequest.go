package media

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTjbPictureFolderCreateAPIRequest) Reset() {
	r._folderName = ""
	r._parentFolderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTjbPictureFolderCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.tjb.picture.folder.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTjbPictureFolderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTjbPictureFolderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaTjbPictureFolderCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTjbPictureFolderCreateRequest()
	},
}

// GetAlibabaTjbPictureFolderCreateRequest 从 sync.Pool 获取 AlibabaTjbPictureFolderCreateAPIRequest
func GetAlibabaTjbPictureFolderCreateAPIRequest() *AlibabaTjbPictureFolderCreateAPIRequest {
	return poolAlibabaTjbPictureFolderCreateAPIRequest.Get().(*AlibabaTjbPictureFolderCreateAPIRequest)
}

// ReleaseAlibabaTjbPictureFolderCreateAPIRequest 将 AlibabaTjbPictureFolderCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaTjbPictureFolderCreateAPIRequest(v *AlibabaTjbPictureFolderCreateAPIRequest) {
	v.Reset()
	poolAlibabaTjbPictureFolderCreateAPIRequest.Put(v)
}
