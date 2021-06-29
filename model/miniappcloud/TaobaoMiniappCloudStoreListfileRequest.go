package miniappcloud

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
云存储根据文件名反查地址 APIRequest
taobao.miniapp.cloud.store.listfile

云存储中，根据文件名反查地址
*/
type TaobaoMiniappCloudStoreListfileRequest struct {
    model.Params

    // 环境(online/test)
    env   string 

    // 文件全路径名
    filePath   string 

}

func NewTaobaoMiniappCloudStoreListfileRequest() *TaobaoMiniappCloudStoreListfileRequest{
    return &TaobaoMiniappCloudStoreListfileRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMiniappCloudStoreListfileRequest) GetApiMethodName() string {
    return "taobao.miniapp.cloud.store.listfile"
}

func (r TaobaoMiniappCloudStoreListfileRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMiniappCloudStoreListfileRequest) SetEnv(env string) error {
    r.env = env
    r.Set("env", env)
    return nil
}

func (r TaobaoMiniappCloudStoreListfileRequest) GetEnv() string {
    return r.env
}

func (r *TaobaoMiniappCloudStoreListfileRequest) SetFilePath(filePath string) error {
    r.filePath = filePath
    r.Set("file_path", filePath)
    return nil
}

func (r TaobaoMiniappCloudStoreListfileRequest) GetFilePath() string {
    return r.filePath
}

