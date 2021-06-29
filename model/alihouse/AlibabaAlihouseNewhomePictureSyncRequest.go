package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
图片数据同步 API请求
alibaba.alihouse.newhome.picture.sync

图片数据同步
*/
type AlibabaAlihouseNewhomePictureSyncRequest struct {
    model.Params
    // 数据
    _projectPictureData   *ProjectPictureDto
}

// 初始化AlibabaAlihouseNewhomePictureSyncRequest对象
func NewAlibabaAlihouseNewhomePictureSyncRequest() *AlibabaAlihouseNewhomePictureSyncRequest{
    return &AlibabaAlihouseNewhomePictureSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomePictureSyncRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.picture.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomePictureSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProjectPictureData Setter
// 数据
func (r *AlibabaAlihouseNewhomePictureSyncRequest) SetProjectPictureData(_projectPictureData *ProjectPictureDto) error {
    r._projectPictureData = _projectPictureData
    r.Set("project_picture_data", _projectPictureData)
    return nil
}

// ProjectPictureData Getter
func (r AlibabaAlihouseNewhomePictureSyncRequest) GetProjectPictureData() *ProjectPictureDto {
    return r._projectPictureData
}
