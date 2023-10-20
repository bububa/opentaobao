package miniapp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappcloudstorerelationaddAPIRequest 云存储添加 API请求
// taobao.miniapp.cloud.store.relation.add
//
// 用于用户上传文件之后回写云存储的关联关系
type TaobaominiappcloudstorerelationaddAPIRequest struct {
	model.Params
	// 环境 test/online
	_env string
	// 文件类型 image/audio/video/font/other
	_fileType string
	// 上传平台返回的文件唯一ID
	_specialId string
	// 上传平台返回的文件url，部分文件类型无固定url，非必填
	_url string
	// 文件路径
	_cloudPath string
}

// NewTaobaominiappcloudstorerelationaddRequest 初始化TaobaominiappcloudstorerelationaddAPIRequest对象
func NewTaobaominiappcloudstorerelationaddRequest() *TaobaominiappcloudstorerelationaddAPIRequest {
	return &TaobaominiappcloudstorerelationaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiappcloudstorerelationaddAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.cloud.store.relation.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiappcloudstorerelationaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiappcloudstorerelationaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEnv is Env Setter
// 环境 test/online
func (r *TaobaominiappcloudstorerelationaddAPIRequest) SetEnv(_env string) error {
	r._env = _env
	r.Set("env", _env)
	return nil
}

// GetEnv Env Getter
func (r TaobaominiappcloudstorerelationaddAPIRequest) GetEnv() string {
	return r._env
}

// SetFileType is FileType Setter
// 文件类型 image/audio/video/font/other
func (r *TaobaominiappcloudstorerelationaddAPIRequest) SetFileType(_fileType string) error {
	r._fileType = _fileType
	r.Set("file_type", _fileType)
	return nil
}

// GetFileType FileType Getter
func (r TaobaominiappcloudstorerelationaddAPIRequest) GetFileType() string {
	return r._fileType
}

// SetSpecialId is SpecialId Setter
// 上传平台返回的文件唯一ID
func (r *TaobaominiappcloudstorerelationaddAPIRequest) SetSpecialId(_specialId string) error {
	r._specialId = _specialId
	r.Set("special_id", _specialId)
	return nil
}

// GetSpecialId SpecialId Getter
func (r TaobaominiappcloudstorerelationaddAPIRequest) GetSpecialId() string {
	return r._specialId
}

// SetUrl is Url Setter
// 上传平台返回的文件url，部分文件类型无固定url，非必填
func (r *TaobaominiappcloudstorerelationaddAPIRequest) SetUrl(_url string) error {
	r._url = _url
	r.Set("url", _url)
	return nil
}

// GetUrl Url Getter
func (r TaobaominiappcloudstorerelationaddAPIRequest) GetUrl() string {
	return r._url
}

// SetCloudPath is CloudPath Setter
// 文件路径
func (r *TaobaominiappcloudstorerelationaddAPIRequest) SetCloudPath(_cloudPath string) error {
	r._cloudPath = _cloudPath
	r.Set("cloud_path", _cloudPath)
	return nil
}

// GetCloudPath CloudPath Getter
func (r TaobaominiappcloudstorerelationaddAPIRequest) GetCloudPath() string {
	return r._cloudPath
}
