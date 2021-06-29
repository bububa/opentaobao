package miniappcloud

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
云存储根据文件名反查地址 API请求
taobao.miniapp.cloud.store.listfile

云存储中，根据文件名反查地址
*/
type TaobaoMiniappCloudStoreListfileRequest struct {
    model.Params
    // 环境(online/test)
    _env   string
    // 文件全路径名
    _filePath   string
}

// 初始化TaobaoMiniappCloudStoreListfileRequest对象
func NewTaobaoMiniappCloudStoreListfileRequest() *TaobaoMiniappCloudStoreListfileRequest{
    return &TaobaoMiniappCloudStoreListfileRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappCloudStoreListfileRequest) GetApiMethodName() string {
    return "taobao.miniapp.cloud.store.listfile"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappCloudStoreListfileRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Env Setter
// 环境(online/test)
func (r *TaobaoMiniappCloudStoreListfileRequest) SetEnv(_env string) error {
    r._env = _env
    r.Set("env", _env)
    return nil
}

// Env Getter
func (r TaobaoMiniappCloudStoreListfileRequest) GetEnv() string {
    return r._env
}
// FilePath Setter
// 文件全路径名
func (r *TaobaoMiniappCloudStoreListfileRequest) SetFilePath(_filePath string) error {
    r._filePath = _filePath
    r.Set("file_path", _filePath)
    return nil
}

// FilePath Getter
func (r TaobaoMiniappCloudStoreListfileRequest) GetFilePath() string {
    return r._filePath
}
