package miniapp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
云存储添加 API请求
taobao.miniapp.cloud.store.relation.add

用于用户上传文件之后回写云存储的关联关系
*/
type TaobaoMiniappCloudStoreRelationAddRequest struct {
    model.Params
    // 环境 test/online
    _env   string
    // 文件类型 image/audio/video/font/other
    _fileType   string
    // 上传平台返回的文件唯一ID
    _specialId   string
    // 上传平台返回的文件url，部分文件类型无固定url，非必填
    _url   string
    // 文件路径
    _cloudPath   string
}

// 初始化TaobaoMiniappCloudStoreRelationAddRequest对象
func NewTaobaoMiniappCloudStoreRelationAddRequest() *TaobaoMiniappCloudStoreRelationAddRequest{
    return &TaobaoMiniappCloudStoreRelationAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappCloudStoreRelationAddRequest) GetApiMethodName() string {
    return "taobao.miniapp.cloud.store.relation.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappCloudStoreRelationAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Env Setter
// 环境 test/online
func (r *TaobaoMiniappCloudStoreRelationAddRequest) SetEnv(_env string) error {
    r._env = _env
    r.Set("env", _env)
    return nil
}

// Env Getter
func (r TaobaoMiniappCloudStoreRelationAddRequest) GetEnv() string {
    return r._env
}
// FileType Setter
// 文件类型 image/audio/video/font/other
func (r *TaobaoMiniappCloudStoreRelationAddRequest) SetFileType(_fileType string) error {
    r._fileType = _fileType
    r.Set("file_type", _fileType)
    return nil
}

// FileType Getter
func (r TaobaoMiniappCloudStoreRelationAddRequest) GetFileType() string {
    return r._fileType
}
// SpecialId Setter
// 上传平台返回的文件唯一ID
func (r *TaobaoMiniappCloudStoreRelationAddRequest) SetSpecialId(_specialId string) error {
    r._specialId = _specialId
    r.Set("special_id", _specialId)
    return nil
}

// SpecialId Getter
func (r TaobaoMiniappCloudStoreRelationAddRequest) GetSpecialId() string {
    return r._specialId
}
// Url Setter
// 上传平台返回的文件url，部分文件类型无固定url，非必填
func (r *TaobaoMiniappCloudStoreRelationAddRequest) SetUrl(_url string) error {
    r._url = _url
    r.Set("url", _url)
    return nil
}

// Url Getter
func (r TaobaoMiniappCloudStoreRelationAddRequest) GetUrl() string {
    return r._url
}
// CloudPath Setter
// 文件路径
func (r *TaobaoMiniappCloudStoreRelationAddRequest) SetCloudPath(_cloudPath string) error {
    r._cloudPath = _cloudPath
    r.Set("cloud_path", _cloudPath)
    return nil
}

// CloudPath Getter
func (r TaobaoMiniappCloudStoreRelationAddRequest) GetCloudPath() string {
    return r._cloudPath
}
