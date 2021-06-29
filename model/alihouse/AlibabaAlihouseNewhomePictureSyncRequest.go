package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
图片数据同步 APIRequest
alibaba.alihouse.newhome.picture.sync

图片数据同步
*/
type AlibabaAlihouseNewhomePictureSyncRequest struct {
    model.Params

    // 数据
    projectPictureData   *ProjectPictureDto 

}

func NewAlibabaAlihouseNewhomePictureSyncRequest() *AlibabaAlihouseNewhomePictureSyncRequest{
    return &AlibabaAlihouseNewhomePictureSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihouseNewhomePictureSyncRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.picture.sync"
}

func (r AlibabaAlihouseNewhomePictureSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihouseNewhomePictureSyncRequest) SetProjectPictureData(projectPictureData *ProjectPictureDto) error {
    r.projectPictureData = projectPictureData
    r.Set("project_picture_data", projectPictureData)
    return nil
}

func (r AlibabaAlihouseNewhomePictureSyncRequest) GetProjectPictureData() *ProjectPictureDto {
    return r.projectPictureData
}

