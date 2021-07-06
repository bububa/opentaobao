package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMediaFileAddAPIRequest 多媒体平台文件添加 API请求
// taobao.media.file.add
//
// 用户通过top上传文件到多媒体平台
type TaobaoMediaFileAddAPIRequest struct {
	model.Params
	// 上传文件的名称
	_name string
	// 接入多媒体平台的业务code每个应用有一个特有的业务code
	_bizCode string
	// 文件属于的那个目录的目录编号
	_dirId int64
	// 额外信息
	_ext int64
	// 文件
	_fileData *model.File
}

// NewTaobaoMediaFileAddRequest 初始化TaobaoMediaFileAddAPIRequest对象
func NewTaobaoMediaFileAddRequest() *TaobaoMediaFileAddAPIRequest {
	return &TaobaoMediaFileAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMediaFileAddAPIRequest) GetApiMethodName() string {
	return "taobao.media.file.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMediaFileAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetName is Name Setter
// 上传文件的名称
func (r *TaobaoMediaFileAddAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoMediaFileAddAPIRequest) GetName() string {
	return r._name
}

// SetBizCode is BizCode Setter
// 接入多媒体平台的业务code每个应用有一个特有的业务code
func (r *TaobaoMediaFileAddAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r TaobaoMediaFileAddAPIRequest) GetBizCode() string {
	return r._bizCode
}

// SetDirId is DirId Setter
// 文件属于的那个目录的目录编号
func (r *TaobaoMediaFileAddAPIRequest) SetDirId(_dirId int64) error {
	r._dirId = _dirId
	r.Set("dir_id", _dirId)
	return nil
}

// GetDirId DirId Getter
func (r TaobaoMediaFileAddAPIRequest) GetDirId() int64 {
	return r._dirId
}

// SetExt is Ext Setter
// 额外信息
func (r *TaobaoMediaFileAddAPIRequest) SetExt(_ext int64) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoMediaFileAddAPIRequest) GetExt() int64 {
	return r._ext
}

// SetFileData is FileData Setter
// 文件
func (r *TaobaoMediaFileAddAPIRequest) SetFileData(_fileData *model.File) error {
	r._fileData = _fileData
	r.Set("file_data", _fileData)
	return nil
}

// GetFileData FileData Getter
func (r TaobaoMediaFileAddAPIRequest) GetFileData() *model.File {
	return r._fileData
}
