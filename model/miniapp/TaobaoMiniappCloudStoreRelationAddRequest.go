package miniapp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
云存储添加 APIRequest
taobao.miniapp.cloud.store.relation.add

用于用户上传文件之后回写云存储的关联关系
*/
type TaobaoMiniappCloudStoreRelationAddRequest struct {
    model.Params

    // 环境 test/online
    env   string 

    // 文件类型 image/audio/video/font/other
    fileType   string 

    // 上传平台返回的文件唯一ID
    specialId   string 

    // 上传平台返回的文件url，部分文件类型无固定url，非必填
    url   string 

    // 文件路径
    cloudPath   string 

}

func NewTaobaoMiniappCloudStoreRelationAddRequest() *TaobaoMiniappCloudStoreRelationAddRequest{
    return &TaobaoMiniappCloudStoreRelationAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMiniappCloudStoreRelationAddRequest) GetApiMethodName() string {
    return "taobao.miniapp.cloud.store.relation.add"
}

func (r TaobaoMiniappCloudStoreRelationAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMiniappCloudStoreRelationAddRequest) SetEnv(env string) error {
    r.env = env
    r.Set("env", env)
    return nil
}

func (r TaobaoMiniappCloudStoreRelationAddRequest) GetEnv() string {
    return r.env
}

func (r *TaobaoMiniappCloudStoreRelationAddRequest) SetFileType(fileType string) error {
    r.fileType = fileType
    r.Set("file_type", fileType)
    return nil
}

func (r TaobaoMiniappCloudStoreRelationAddRequest) GetFileType() string {
    return r.fileType
}

func (r *TaobaoMiniappCloudStoreRelationAddRequest) SetSpecialId(specialId string) error {
    r.specialId = specialId
    r.Set("special_id", specialId)
    return nil
}

func (r TaobaoMiniappCloudStoreRelationAddRequest) GetSpecialId() string {
    return r.specialId
}

func (r *TaobaoMiniappCloudStoreRelationAddRequest) SetUrl(url string) error {
    r.url = url
    r.Set("url", url)
    return nil
}

func (r TaobaoMiniappCloudStoreRelationAddRequest) GetUrl() string {
    return r.url
}

func (r *TaobaoMiniappCloudStoreRelationAddRequest) SetCloudPath(cloudPath string) error {
    r.cloudPath = cloudPath
    r.Set("cloud_path", cloudPath)
    return nil
}

func (r TaobaoMiniappCloudStoreRelationAddRequest) GetCloudPath() string {
    return r.cloudPath
}

