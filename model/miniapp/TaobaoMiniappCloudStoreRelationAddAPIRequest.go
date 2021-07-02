package miniapp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappCloudStoreRelationAddAPIRequest 云存储添加 API请求
// taobao.miniapp.cloud.store.relation.add
//
// 用于用户上传文件之后回写云存储的关联关系
type TaobaoMiniappCloudStoreRelationAddAPIRequest struct {
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

// NewTaobaoMiniappCloudStoreRelationAddRequest 初始化TaobaoMiniappCloudStoreRelationAddAPIRequest对象
func NewTaobaoMiniappCloudStoreRelationAddRequest() *TaobaoMiniappCloudStoreRelationAddAPIRequest {
	return &TaobaoMiniappCloudStoreRelationAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappCloudStoreRelationAddAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.cloud.store.relation.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappCloudStoreRelationAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Env Setter
// 环境 test/online
func (r *TaobaoMiniappCloudStoreRelationAddAPIRequest) SetEnv(_env string) error {
	r._env = _env
	r.Set("env", _env)
	return nil
}

// Get Env Getter
func (r TaobaoMiniappCloudStoreRelationAddAPIRequest) GetEnv() string {
	return r._env
}

// Set is FileType Setter
// 文件类型 image/audio/video/font/other
func (r *TaobaoMiniappCloudStoreRelationAddAPIRequest) SetFileType(_fileType string) error {
	r._fileType = _fileType
	r.Set("file_type", _fileType)
	return nil
}

// Get FileType Getter
func (r TaobaoMiniappCloudStoreRelationAddAPIRequest) GetFileType() string {
	return r._fileType
}

// Set is SpecialId Setter
// 上传平台返回的文件唯一ID
func (r *TaobaoMiniappCloudStoreRelationAddAPIRequest) SetSpecialId(_specialId string) error {
	r._specialId = _specialId
	r.Set("special_id", _specialId)
	return nil
}

// Get SpecialId Getter
func (r TaobaoMiniappCloudStoreRelationAddAPIRequest) GetSpecialId() string {
	return r._specialId
}

// Set is Url Setter
// 上传平台返回的文件url，部分文件类型无固定url，非必填
func (r *TaobaoMiniappCloudStoreRelationAddAPIRequest) SetUrl(_url string) error {
	r._url = _url
	r.Set("url", _url)
	return nil
}

// Get Url Getter
func (r TaobaoMiniappCloudStoreRelationAddAPIRequest) GetUrl() string {
	return r._url
}

// Set is CloudPath Setter
// 文件路径
func (r *TaobaoMiniappCloudStoreRelationAddAPIRequest) SetCloudPath(_cloudPath string) error {
	r._cloudPath = _cloudPath
	r.Set("cloud_path", _cloudPath)
	return nil
}

// Get CloudPath Getter
func (r TaobaoMiniappCloudStoreRelationAddAPIRequest) GetCloudPath() string {
	return r._cloudPath
}
