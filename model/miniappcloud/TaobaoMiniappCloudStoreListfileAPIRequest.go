package miniappcloud

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappCloudStoreListfileAPIRequest 云存储根据文件名反查地址 API请求
// taobao.miniapp.cloud.store.listfile
//
// 云存储中，根据文件名反查地址
type TaobaoMiniappCloudStoreListfileAPIRequest struct {
	model.Params
	// 环境(online/test)
	_env string
	// 文件全路径名
	_filePath string
}

// NewTaobaoMiniappCloudStoreListfileRequest 初始化TaobaoMiniappCloudStoreListfileAPIRequest对象
func NewTaobaoMiniappCloudStoreListfileRequest() *TaobaoMiniappCloudStoreListfileAPIRequest {
	return &TaobaoMiniappCloudStoreListfileAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappCloudStoreListfileAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.cloud.store.listfile"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappCloudStoreListfileAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappCloudStoreListfileAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEnv is Env Setter
// 环境(online/test)
func (r *TaobaoMiniappCloudStoreListfileAPIRequest) SetEnv(_env string) error {
	r._env = _env
	r.Set("env", _env)
	return nil
}

// GetEnv Env Getter
func (r TaobaoMiniappCloudStoreListfileAPIRequest) GetEnv() string {
	return r._env
}

// SetFilePath is FilePath Setter
// 文件全路径名
func (r *TaobaoMiniappCloudStoreListfileAPIRequest) SetFilePath(_filePath string) error {
	r._filePath = _filePath
	r.Set("file_path", _filePath)
	return nil
}

// GetFilePath FilePath Getter
func (r TaobaoMiniappCloudStoreListfileAPIRequest) GetFilePath() string {
	return r._filePath
}
