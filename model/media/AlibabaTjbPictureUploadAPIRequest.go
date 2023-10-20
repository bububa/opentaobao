package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatjbpictureuploadAPIRequest 淘特图片空间上传单张图片 API请求
// alibaba.tjb.picture.upload
//
// 淘特图片空间上传单张图片
type AlibabatjbpictureuploadAPIRequest struct {
	model.Params
	// 图片名，可包含扩展名，与上传后返回的图片名一致。最长50个字符，不允许/,\,:,*,?,",<,>,|,;这些特殊字符。
	_fileName string
	// 图片的二进制流。允许png、jpg、gif图片格式,3M以内。
	_fileData *model.File
	// 指定传入的文件夹id，根文件夹为0
	_folderId int64
}

// NewAlibabatjbpictureuploadRequest 初始化AlibabatjbpictureuploadAPIRequest对象
func NewAlibabatjbpictureuploadRequest() *AlibabatjbpictureuploadAPIRequest {
	return &AlibabatjbpictureuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatjbpictureuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tjb.picture.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatjbpictureuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatjbpictureuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFileName is FileName Setter
// 图片名，可包含扩展名，与上传后返回的图片名一致。最长50个字符，不允许/,\,:,*,?,&#34;,&lt;,&gt;,|,;这些特殊字符。
func (r *AlibabatjbpictureuploadAPIRequest) SetFileName(_fileName string) error {
	r._fileName = _fileName
	r.Set("file_name", _fileName)
	return nil
}

// GetFileName FileName Getter
func (r AlibabatjbpictureuploadAPIRequest) GetFileName() string {
	return r._fileName
}

// SetFileData is FileData Setter
// 图片的二进制流。允许png、jpg、gif图片格式,3M以内。
func (r *AlibabatjbpictureuploadAPIRequest) SetFileData(_fileData *model.File) error {
	r._fileData = _fileData
	r.Set("file_data", _fileData)
	return nil
}

// GetFileData FileData Getter
func (r AlibabatjbpictureuploadAPIRequest) GetFileData() *model.File {
	return r._fileData
}

// SetFolderId is FolderId Setter
// 指定传入的文件夹id，根文件夹为0
func (r *AlibabatjbpictureuploadAPIRequest) SetFolderId(_folderId int64) error {
	r._folderId = _folderId
	r.Set("folder_id", _folderId)
	return nil
}

// GetFolderId FolderId Getter
func (r AlibabatjbpictureuploadAPIRequest) GetFolderId() int64 {
	return r._folderId
}
