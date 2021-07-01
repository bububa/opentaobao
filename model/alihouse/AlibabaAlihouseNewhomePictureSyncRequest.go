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
type AlibabaAlihouseNewhomePictureSyncAPIRequest struct {
    model.Params
    // 数据
    _projectPictureData   *ProjectPictureDTO
}

// 初始化AlibabaAlihouseNewhomePictureSyncAPIRequest对象
func NewAlibabaAlihouseNewhomePictureSyncRequest() *AlibabaAlihouseNewhomePictureSyncAPIRequest{
    return &AlibabaAlihouseNewhomePictureSyncAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomePictureSyncAPIRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.picture.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomePictureSyncAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProjectPictureData Setter
// 数据
func (r *AlibabaAlihouseNewhomePictureSyncAPIRequest) SetProjectPictureData(_projectPictureData *ProjectPictureDTO) error {
    r._projectPictureData = _projectPictureData
    r.Set("project_picture_data", _projectPictureData)
    return nil
}

// ProjectPictureData Getter
func (r AlibabaAlihouseNewhomePictureSyncAPIRequest) GetProjectPictureData() *ProjectPictureDTO {
    return r._projectPictureData
}
